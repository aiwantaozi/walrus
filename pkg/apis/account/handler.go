package account

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"k8s.io/utils/pointer"

	"github.com/seal-io/seal/pkg/apis/account/view"
	"github.com/seal-io/seal/pkg/apis/auth/cache"
	"github.com/seal-io/seal/pkg/apis/auth/session"
	"github.com/seal-io/seal/pkg/apis/runtime"
	"github.com/seal-io/seal/pkg/casdoor"
	"github.com/seal-io/seal/pkg/dao"
	"github.com/seal-io/seal/pkg/dao/model"
	"github.com/seal-io/seal/pkg/dao/model/subject"
	"github.com/seal-io/seal/pkg/settings"
	"github.com/seal-io/seal/utils/req"
)

func Login() runtime.ErrorHandle {
	return func(ctx *gin.Context) error {
		var input struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}

		if err := ctx.Bind(&input); err != nil {
			return err
		}

		// Login.
		internalSessions, err := casdoor.SignInUser(ctx, casdoor.BuiltinApp, casdoor.BuiltinOrg,
			input.Username, input.Password)
		if err != nil {
			return runtime.Error(http.StatusUnauthorized, err)
		}

		// Hold session.
		err = casdoor.HoldSession(ctx.Writer, internalSessions)
		if err != nil {
			return runtime.Errorw(err, "failed to login")
		}

		return nil
	}
}

func Logout() runtime.Handle {
	return func(ctx *gin.Context) {
		internalSession := casdoor.GetSession(ctx.Request.Cookies())
		if internalSession == nil {
			return
		}

		// Logout.
		_ = casdoor.SignOutUser(ctx, []*req.HttpCookie{internalSession})

		// Interrupt session.
		_ = casdoor.InterruptSession(ctx.Writer, []*req.HttpCookie{internalSession})
		cache.CleanSessionSubject(string(internalSession.Value()))
	}
}

func Info(mc model.ClientSet) runtime.ErrorHandle {
	return func(ctx *gin.Context) error {
		switch ctx.Request.Method {
		default:
			ctx.AbortWithStatus(http.StatusMethodNotAllowed)
		case http.MethodPost:
			err := updateInfo(ctx, mc)
			if err != nil {
				return runtime.Errorw(err, "failed to update info")
			}
		case http.MethodGet:
			err := getInfo(ctx, mc)
			if err != nil {
				return runtime.Errorw(err, "failed to get info")
			}
		}

		return nil
	}
}

func updateInfo(ctx *gin.Context, modelClient model.ClientSet) error {
	s := session.LoadSubject(ctx)

	var r view.UpdateInfoRequest
	if err := ctx.ShouldBindJSON(&r); err != nil {
		return runtime.Error(http.StatusBadRequest, err)
	}

	if err := r.Validate(); err != nil {
		return runtime.Error(http.StatusBadRequest, err)
	}

	if r.LoginGroup != nil {
		selves, err := modelClient.Subjects().Query().
			Where(
				subject.Kind("user"),
				subject.Name(s.Name),
			).
			Select(subject.FieldID, subject.FieldGroup, subject.FieldLoginTo).
			All(ctx)
		if err != nil {
			return err
		}
		// Switch login group.
		for i := range selves {
			if *selves[i].LoginTo {
				if selves[i].Group == *r.LoginGroup {
					return runtime.Error(http.StatusBadRequest, "invalid group: the same")
				}
				selves[i].LoginTo = pointer.Bool(false)

				continue
			}

			if selves[i].Group == *r.LoginGroup {
				selves[i].LoginTo = pointer.Bool(true)
			}
		}

		err = modelClient.WithTx(ctx, func(tx *model.Tx) error {
			updates, err := dao.SubjectUpdates(tx, selves...)
			if err != nil {
				return err
			}

			for i := range updates {
				err = updates[i].Exec(ctx)
				if err != nil {
					return err
				}
			}

			return nil
		})
		if err != nil {
			return err
		}
	}

	if r.Password != nil {
		var cred casdoor.ApplicationCredential

		err := settings.CasdoorCred.ValueJSONUnmarshal(ctx, modelClient, &cred)
		if err != nil {
			return err
		}

		err = casdoor.UpdateUserPassword(ctx, cred.ClientID, cred.ClientSecret,
			casdoor.BuiltinOrg, s.Name, *r.OldPassword, *r.Password)
		if err != nil {
			if strings.HasSuffix(err.Error(), "not found") {
				return runtime.Error(http.StatusNotFound,
					"not found user")
			}

			return runtime.Error(http.StatusBadRequest, err)
		}
		// Update setting to indicate the initialized password has been changed.
		if settings.FirstLogin.ShouldValueBool(ctx, modelClient) {
			_, err = settings.FirstLogin.Set(ctx, modelClient, "false")
			return err
		}
	}

	return nil
}

func getInfo(ctx *gin.Context, modelClient model.ClientSet) error {
	var (
		s    = session.LoadSubject(ctx)
		info = &view.GetInfoResponse{
			Name:       s.Name,
			Roles:      s.Roles,
			Policies:   s.Policies,
			LoginGroup: s.Group,
		}
	)

	// Get belong groups.
	selves, err := modelClient.Subjects().Query().
		Where(
			subject.Kind("user"),
			subject.Name(s.Name),
		).
		Select(subject.FieldGroup, subject.FieldPaths).
		All(ctx)
	if err != nil {
		return err
	}
	info.Groups = make([]view.GetInfoGroup, 0, len(selves))

	for i := 0; i < len(selves); i++ {
		u := selves[i]

		info.Groups = append(info.Groups,
			view.GetInfoGroup{
				Name:  u.Group,
				Paths: u.Paths[:len(u.Paths)-1],
			})
	}
	ctx.JSON(http.StatusOK, info)

	return nil
}
