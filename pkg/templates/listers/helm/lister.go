package terraform

import (
	"context"
	"fmt"
	"os"
	"regexp"

	"github.com/seal-io/getter"
	"github.com/seal-io/walrus/pkg/system"

	helmrepo "helm.sh/helm/v3/pkg/repo"

	walrusv1 "github.com/seal-io/walrus/pkg/apis/walrus/v1"
	walruscore "github.com/seal-io/walrus/pkg/apis/walruscore/v1"
	"github.com/seal-io/walrus/pkg/templates/utils"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Lister struct {
}

func New() *Lister {
	return &Lister{}
}

func (l *Lister) List(ctx context.Context, c walrusv1.Catalog) ([]walrusv1.Template, error) {
	//logger := log.WithName("catalog")
	return l.listTemplate(ctx, c)
}

// listTemplate returns a list of templates from the given catalog.
func (l *Lister) listTemplate(ctx context.Context, c walrusv1.Catalog) ([]walrusv1.Template, error) {
	indexFile, err := l.loadIndex(ctx, c)
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

	list := make([]walrusv1.Template, 0, len(indexFile.Entries))
	for name, tmpls := range indexFile.Entries {
		if includeReg != nil && !includeReg.MatchString(name) {
			continue
		}

		if excludeReg != nil && excludeReg.MatchString(name) {
			continue
		}

		latest := tmpls[0]
		url := latest.Home
		if len(latest.Sources) != 0 {
			url = latest.Sources[0]
		}

		// Gen versions.
		var versions = make([]mytypes.TemplateVersion, len(tmpls))
		for i, cv := range tmpls {
			versions[i] = mytypes.TemplateVersion{
				Version: cv.Version,
				URL:     cv.URLs[0],
			}
		}

		t := walrusv1.Template{
			ObjectMeta: metav1.ObjectMeta{
				Namespace: "TODO(michelia)",
				Name:      utils.NormalizeTemplateName(c.Name, name),
				Labels:    utils.GenWalrusBuiltinLabels([]string{"c-kubernetes"}), // TODO(michelia)
				OwnerReferences: []metav1.OwnerReference{
					{
						APIVersion: c.APIVersion,
						Kind:       c.Kind(),
						Name:       c.Name,
						UID:        c.UID,
					},
				},
			},
			Spec: walruscore.TemplateSpec{
				TemplateFormat: mytypes.TemplateFormatHelm,
				Description:    latest.Description,
				URL:            url,
				Auth:           c.Spec.Auth,
				Advance:        c.Spec.Advance,
			},
			Status: walruscore.TemplateStatus{
				OriginalName: name,
				Icon:         latest.Icon,
				Versions:     versions,
			},
		}

		list = append(list, t)
	}

	//logger.Infof("found %d valid templates in %s", len(repos), c.Source)

	return list, nil
}

func (l *Lister) loadIndex(ctx context.Context, c walrusv1.Catalog) (*helmrepo.IndexFile, error) {
	loopbackKubeCli := system.LoopbackKubeClient.Get()

	opts := make([]getter.Option, 0)

	if auth := c.Spec.Auth; auth != nil && auth.AuthSecretRef != nil {
		secret, err := loopbackKubeCli.CoreV1().Secrets(c.Namespace).Get(
			ctx,
			auth.AuthSecretRef.Name,
			metav1.GetOptions{},
		)
		if err != nil {
			return nil, err
		}

		switch auth.AuthType {
		case mytypes.AuthTypeToken:
			token, ok := secret.StringData["token"]
			if !ok {
				return nil, fmt.Errorf("token not found in secret %s", auth.AuthSecretRef.Name)
			}

			opts = append(opts, getter.WithBeaterToken(token))
		case mytypes.AuthTypeBasic:
			username, ok := secret.StringData["username"]
			if !ok {
				return nil, fmt.Errorf("username not found in secret %s", auth.AuthSecretRef.Name)
			}

			password, ok := secret.StringData["password"]
			if !ok {
				return nil, fmt.Errorf("password not found in secret %s", auth.AuthSecretRef.Name)
			}
			opts = append(opts, getter.WithBasicAuth(username, password))
		}
	}

	if advance := c.Spec.Advance; advance != nil {
		opts = append(opts, getter.WithInsecureSkipVerifyTLS(advance.SkipTLSVerification))

		if advance.CertSecretRef != nil {
			secret, err := loopbackKubeCli.CoreV1().Secrets(c.Namespace).Get(
				ctx,
				advance.CertSecretRef.Name,
				metav1.GetOptions{},
			)
			if err != nil {
				return nil, err
			}

			opts = append(opts, getter.WithTLSClientConfig(
				secret.Data["tls.crt"],
				secret.Data["tls.key"],
				secret.Data["ca.crt"],
			))
		}
	}

	g, err := getter.NewHTTPGetter(opts...)
	if err != nil {
		return nil, err
	}

	b, err := g.Get(c.Spec.URL)
	if err != nil {
		return nil, err
	}

	f, err := os.CreateTemp("", "seal-helm-index")
	if err != nil {
		panic(fmt.Errorf("error creating temp file: %w", err))
	}
	defer os.Remove(f.Name())

	_, err = f.Write(b.Bytes())
	if err != nil {
		return nil, err
	}

	return helmrepo.LoadIndexFile(f.Name())
}
