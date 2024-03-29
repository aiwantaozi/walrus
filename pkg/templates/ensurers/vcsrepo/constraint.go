package vcsrepo

import (
	"path/filepath"

	"github.com/go-git/go-git/v5"
	"github.com/hashicorp/go-version"

	walruscore "github.com/seal-io/walrus/pkg/apis/walruscore/v1"
	"github.com/seal-io/walrus/pkg/templates/loader"
	"github.com/seal-io/walrus/pkg/templates/sourceurl"
	"github.com/seal-io/walrus/pkg/types"
	"github.com/seal-io/walrus/pkg/vcs"
	"k8s.io/klog/v2"
)

// TODO(michelia): handler subpath
// getValidVersions get valid terraform module versions.
func getValidVersions(
	t *walruscore.Template,
	r *git.Repository,
	versions []*version.Version,
) ([]*version.Version, map[*version.Version]*types.SchemaGroup, error) {
	su, err := sourceurl.ParseURLToSourceURL(t.Spec.VCSRepository.URL)
	if err != nil {
		return nil, nil, err
	}

	w, err := r.Worktree()
	if err != nil {
		return nil, nil, err
	}

	validVersions := make([]*version.Version, 0, len(versions))
	versionSchema := make(map[*version.Version]*types.SchemaGroup)

	for i := range versions {
		v := versions[i]
		tag := v.Original()

		resetRef, err := vcs.GetRepoRef(r, tag)
		if err != nil {
			// TODO(michelia): log
			//logger.Warnf("failed to get \"%s:%s\" of catalog %q git reference: %v",
			//	entity.Name, tag, entity.CatalogID, err)
			continue
		}

		hash := resetRef.Hash()

		// If tag is not a commit hash, get commit hash from tag object target.
		object, err := r.TagObject(hash)
		if err == nil {
			hash = object.Target
		}

		err = w.Reset(&git.ResetOptions{
			Commit: hash,
			Mode:   git.HardReset,
		})
		if err != nil {
			// TODO(michelia): log
			//logger.Warnf("failed set \"%s:%s\" of catalog %q: %v", entity.Name, tag, entity.CatalogID, err)
			continue
		}

		klog.V(5).Infof("get schema of template \"%s:%s\", version \"%s\"", t.Namespace, t.Name, tag)

		dir := w.Filesystem.Root()
		if su.SubPath != "" {
			dir = filepath.Join(dir, su.SubPath)
		}

		sg, err := loader.LoadSchema(dir, t.Name, t.Spec.TemplateFormat)
		if err != nil {
			// TODO(michelia): log
			//logger.Warnf("error get \"%s:%s\"'s schema of catalog %q: %v", entity.Name, tag, entity.CatalogID, err)
			continue
		}

		validVersions = append(validVersions, v)
		versionSchema[v] = sg
	}

	return validVersions, versionSchema, nil
}
