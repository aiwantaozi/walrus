package walrus

import (
	"context"
	"fmt"

	authentication "k8s.io/api/authentication/v1"
	core "k8s.io/api/core/v1"
	kerrors "k8s.io/apimachinery/pkg/api/errors"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/apiserver/pkg/registry/rest"
	"k8s.io/utils/ptr"
	ctrlcli "sigs.k8s.io/controller-runtime/pkg/client"

	walrus "github.com/seal-io/walrus/pkg/apis/walrus/v1"
	"github.com/seal-io/walrus/pkg/extensionapi"
	"github.com/seal-io/walrus/pkg/systemauthz"
	"github.com/seal-io/walrus/pkg/systemsetting"
)

// SubjectTokenHandler is the handler for v1.SubjectToken objects,
// which is a subresource of v1.Subject objects.
type SubjectTokenHandler struct {
	extensionapi.ObjectInfo
	extensionapi.CreateOperation

	Client ctrlcli.Client
}

var (
	_ rest.Storage = (*SubjectTokenHandler)(nil)
	_ rest.Creater = (*SubjectTokenHandler)(nil)
)

func newSubjectTokenHandler(opts extensionapi.SetupOptions) *SubjectTokenHandler {
	h := &SubjectTokenHandler{}

	// As storage.
	h.ObjectInfo = &walrus.SubjectToken{}
	h.CreateOperation = extensionapi.WithCreate(h)

	// Set client.
	h.Client = opts.Manager.GetClient()

	return h
}

func (h *SubjectTokenHandler) New() runtime.Object {
	return &walrus.SubjectToken{}
}

func (h *SubjectTokenHandler) Destroy() {}

func (h *SubjectTokenHandler) OnCreate(ctx context.Context, obj runtime.Object, _ ctrlcli.CreateOptions) (runtime.Object, error) {
	subjt := obj.(*walrus.SubjectToken)

	// Validate.
	{
		var errs field.ErrorList
		if es := ptr.Deref(subjt.Spec.ExpirationSeconds, 0); es <= 0 {
			errs = append(errs, field.Invalid(
				field.NewPath("spec.expirationSeconds"), subjt.Spec.ExpirationSeconds, "must be greater than 0"),
			)
		} else {
			limit, err := systemsetting.SubjectTokenMaximumExpirationSeconds.ValueInt64(ctx)
			if err != nil {
				return nil, kerrors.NewInternalError(err)
			}
			if es > limit {
				errs = append(errs, field.Invalid(
					field.NewPath("spec.expirationSeconds"), subjt.Spec.ExpirationSeconds,
					fmt.Sprintf("must be less than %v", limit)),
				)
			}
		}
		if len(errs) > 0 {
			return nil, kerrors.NewInvalid(walrus.SchemeKind("subjecttokens"), subjt.Name, errs)
		}

		subj := &walrus.Subject{
			ObjectMeta: meta.ObjectMeta{
				Namespace: subjt.Namespace,
				Name:      subjt.Name,
			},
		}
		err := h.Client.Get(ctx, ctrlcli.ObjectKeyFromObject(subj), subj)
		if err != nil {
			return nil, kerrors.NewBadRequest("subject is not found")
		}
	}

	// Generate a token.
	sa := &core.ServiceAccount{
		ObjectMeta: meta.ObjectMeta{
			Namespace: subjt.Namespace,
			Name:      systemauthz.ConvertServiceAccountNameFromSubjectName(subjt.Name),
		},
	}
	tr := &authentication.TokenRequest{
		Spec: authentication.TokenRequestSpec{
			ExpirationSeconds: ptr.To(ptr.Deref(subjt.Spec.ExpirationSeconds, 3600)),
		},
	}
	err := h.Client.SubResource("token").Create(ctx, sa, tr)
	if err != nil {
		return nil, kerrors.NewInternalError(err)
	}

	subjt.ResourceVersion = tr.ResourceVersion
	subjt.Status.Token = tr.Status.Token
	subjt.Status.ExpirationTimestamp = tr.Status.ExpirationTimestamp
	return subjt, nil
}
