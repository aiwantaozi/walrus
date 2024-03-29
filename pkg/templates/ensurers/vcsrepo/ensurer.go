package vcsrepo

import (
	"context"
	"os"
	"path/filepath"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/hashicorp/go-version"

	transphttp "github.com/go-git/go-git/v5/plumbing/transport/http"

	"github.com/seal-io/utils/stringx"
	walruscore "github.com/seal-io/walrus/pkg/apis/walruscore/v1"
	"github.com/seal-io/walrus/pkg/systemsetting"
	"github.com/seal-io/walrus/pkg/templates/common"
	ensurercommon "github.com/seal-io/walrus/pkg/templates/ensurers/common"
	"github.com/seal-io/walrus/pkg/templates/sourceurl"
	"github.com/seal-io/walrus/pkg/types"
	"github.com/seal-io/walrus/pkg/vcs"
	"k8s.io/klog/v2"
)

type Ensurer struct {
}

func New() *Ensurer {
	return &Ensurer{}
}

// Ensure complete the template info and save it.
func (l *Ensurer) Ensure(ctx context.Context, t *walruscore.Template) error {
	tempDir := filepath.Join(os.TempDir(), "seal-template-"+stringx.RandomHex(10))
	defer os.RemoveAll(tempDir)

	tlsVerify, err := systemsetting.EnableRemoteTlsVerify.ValueBool(ctx)
	if err != nil {
		return err
	}

	source := t.Spec.VCSRepository
	opts := vcs.GitCloneOptions{
		URL:             source.URL,
		InsecureSkipTLS: !tlsVerify,
	}

	if auth := source.SecretRef; auth != nil {
		token, err := common.GetToken(ctx, t.Namespace, auth)
		if err != nil {
			return err
		}

		opts.Auth = &transphttp.TokenAuth{
			Token: token.Token,
		}
	}

	cloneCtx, cancel := context.WithTimeout(ctx, 10*time.Minute)
	defer cancel()

	_, err = vcs.GitClone(cloneCtx, tempDir, opts)
	if err != nil {
		return err
	}

	r, err := git.PlainOpen(tempDir)
	if err != nil {
		return err
	}

	// Ensure icon.
	iconURL, err := gitRepoIconURL(r, source.URL)
	if err != nil {
		return err
	}

	// Versions.
	versions, err := vcs.GetRepoVersions(r)
	if err != nil {
		return err
	}

	versions, versionSchema, err := getValidVersions(t, r, versions)
	if err != nil {
		return err
	}

	templateVersions, err := genTemplateVersions(ctx, t, versions, versionSchema)
	if err != nil {
		return err
	}

	// Set template status.
	t.Status.Icon = iconURL
	t.Status.Versions = templateVersions

	return ensurercommon.UpsertTemplate(ctx, t)
}

// genTemplateVersionsFromGitRepo retrieves template versions from a git repository.
func genTemplateVersions(
	ctx context.Context,
	t *walruscore.Template,
	versions []*version.Version,
	versionSchema map[*version.Version]*types.SchemaGroup,
) ([]walruscore.TemplateVersion, error) {

	var (
		logger = klog.NewStandardLogger("WARNING")
		tvs    = make([]walruscore.TemplateVersion, 0, len(versionSchema))
	)

	su, err := sourceurl.ParseURLToSourceURL(t.Spec.VCSRepository.URL)
	if err != nil {
		return nil, err
	}

	for i := range versions {
		var (
			v       = versions[i]
			version = v.Original()
		)

		s := su.Link + "?ref=" + version

		schema, ok := versionSchema[v]
		if !ok {
			logger.Printf("%s/%s version: %s version schema not found", t.Namespace, t.Name, version)
			continue
		}

		// Generate template version.
		tv, err := ensurercommon.GenTemplateVersion(ctx, s, version, t, schema)
		if err != nil {
			return nil, err
		}

		tvs = append(tvs, *tv)
	}

	return tvs, nil
}
