package vcssource

import (
	"context"
	"fmt"
	"regexp"
	"sync/atomic"

	"go.uber.org/multierr"

	"github.com/drone/go-scm/scm"
	"github.com/seal-io/utils/pools/gopool"
	"github.com/seal-io/utils/version"
	walruscore "github.com/seal-io/walrus/pkg/apis/walruscore/v1"
	"github.com/seal-io/walrus/pkg/systemsetting"
	"github.com/seal-io/walrus/pkg/templates/common"
	"github.com/seal-io/walrus/pkg/templates/ensurer"
	"github.com/seal-io/walrus/pkg/templates/utils"
	"github.com/seal-io/walrus/pkg/types"
	"github.com/seal-io/walrus/pkg/vcs"
	"github.com/seal-io/walrus/pkg/vcs/driver"
	"github.com/seal-io/walrus/pkg/vcs/options"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/klog/v2"
)

type Lister struct {
}

func New() *Lister {
	return &Lister{}
}

// TODO(michelia)
func (l *Lister) ListOld(ctx context.Context, c walruscore.Catalog) ([]walruscore.Template, error) {
	logger := klog.Background().WithName("catalog")

	tmpls, err := l.List(ctx, c)
	if err != nil {
		return nil, err
	}

	var (
		total     = len(tmpls)
		processed = int32(0)
		failed    = int32(0)

		wg = gopool.Group()
	)

	batchSize := 5
	for i := 0; i < batchSize; i++ {
		s := i

		wg.Go(func() error {
			// Merge the errors to return them all at once,
			// instead of returning the first error.
			var berr error

			for j := s; j < len(tmpls); j += batchSize {
				tmpl := tmpls[j]

				logger.V(4).Info("syncing \"%s:%s\" of catalog \"%s:%s\"", tmpl.Namespace, tmpl.Name, c.Namespace, c.Name)

				serr := ensurer.Ensure(ctx, &tmpl)
				if serr != nil {
					logger.V(4).Errorf(serr, "failed sync \"%s:%s\" of catalog \"%s:%s\"", tmpl.Namespace, tmpl.Name, c.Namespace, c.Name)
					berr = multierr.Append(berr,
						fmt.Errorf("error syncing \"%s:%s\" of catalog %s: %w",
							tmpl.Namespace, tmpl.Name, c.Name, serr))

					atomic.AddInt32(&failed, 1)
				} else {
					atomic.AddInt32(&processed, 1)
				}

				logger.V(4).Infof("synced catalog \"%s:%s\", total: %d, processed: %d, failed: %d",
					tmpl.Namespace, tmpl.Name, total, processed, failed)
			}

			return berr
		})
	}

	return tmpls, wg.Wait()
}

// List returns a list of templates from the given catalog.
func (l *Lister) List(ctx context.Context, c walruscore.Catalog) ([]walruscore.Template, error) {
	var (
		client *scm.Client
		err    error
		source = c.Spec.VCSSource
	)

	orgName, err := vcs.GetOrgFromGitURL(source.URL)
	if err != nil {
		return nil, err
	}

	sid, err := systemsetting.ServeIdentify.Value(ctx)
	if err != nil {
		return nil, err
	}
	ua := version.GetUserAgent() + "; uuid=" + sid

	tlsVerify, err := systemsetting.EnableRemoteTlsVerify.ValueBool(ctx)
	if err != nil {
		return nil, err
	}

	switch source.Platform {
	case types.GitDriverGithub, types.GitDriverGitlab, types.GitDriverGitee:
		opts := []options.ClientOption{options.WithUserAgent(ua)}

		if !tlsVerify {
			opts = append(opts, options.WithInsecureSkipVerify())
		}
		if secret := source.TokenAuth.SecretRef; secret != nil {
			token, err := common.GetToken(ctx, c.Namespace, secret)
			if err != nil {
				return nil, err
			}

			opts = append(opts, options.WithToken(token.Token))
		}

		client, err = driver.NewClientFromURL(source.Platform, source.URL, opts...)
		if err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("unsupported platform %s for catalog %s:%s", source.Platform, c.Namespace, c.Name)
	}

	repos, err := vcs.GetOrgRepos(ctx, client, orgName)
	if err != nil {
		return nil, err
	}

	var (
		includeReg *regexp.Regexp
		excludeReg *regexp.Regexp
	)

	if filtering := c.Spec.Filtering; filtering != nil {
		if filtering.IncludeFilter != "" {
			includeReg, err = regexp.Compile(filtering.IncludeFilter)
			if err != nil {
				return nil, err
			}
		}

		if filtering.ExcludeFilter != "" {
			excludeReg, err = regexp.Compile(filtering.ExcludeFilter)
			if err != nil {
				return nil, err
			}
		}
	}

	//logger.Infof("found %d repositories in %s", len(repos), c.Source)

	list := make([]walruscore.Template, 0, len(repos))
	for i := range repos {
		repo := repos[i]

		if includeReg != nil && !includeReg.MatchString(repo.Name) {
			continue
		}

		if excludeReg != nil && excludeReg.MatchString(repo.Name) {
			continue
		}

		t := walruscore.Template{
			ObjectMeta: metav1.ObjectMeta{
				Namespace: c.Namespace,
				Name:      utils.NormalizeTemplateName(c.Name, repo.Name),
				Labels:    utils.GenWalrusBuiltinLabels(repo.Topics),
				OwnerReferences: []metav1.OwnerReference{
					{
						APIVersion: c.APIVersion,
						Kind:       c.Kind,
						Name:       c.Name,
						UID:        c.UID,
					},
				},
			},
			Spec: walruscore.TemplateSpec{
				TemplateFormat: c.Spec.TemplateFormat,
				Description:    repo.Description,
				VCSRepository: &walruscore.VCSRepository{
					URL:       repo.Link,
					TokenAuth: c.Spec.VCSSource.TokenAuth,
				},
			},
			Status: walruscore.TemplateStatus{
				OriginalName: repo.Name,
			},
		}

		list = append(list, t)

		// TODO(michelia): project id

	}

	//logger.Infof("found %d valid templates in %s", len(repos), c.Source)

	return list, nil
}
