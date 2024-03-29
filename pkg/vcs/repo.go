package vcs

import (
	"context"
	"fmt"
	"sort"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/hashicorp/go-version"

	"github.com/drone/go-scm/scm"
)

// GetOrgRepos returns full repositories list from the given org.
func GetOrgRepos(ctx context.Context, client *scm.Client, orgName string) ([]*scm.Repository, error) {
	opts := scm.ListOptions{Size: 100}

	var list []*scm.Repository

	for {
		repos, meta, err := client.Organizations.ListRepositories(ctx, orgName, opts)
		if err != nil {
			return nil, err
		}

		for _, src := range repos {
			if src != nil && !src.Archived {
				list = append(list, src)
			}
		}

		opts.Page = meta.Page.Next
		opts.URL = meta.Page.NextURL

		if opts.Page == 0 && opts.URL == "" {
			break
		}
	}

	return list, nil
}

// GetRepoRef returns a reference from a git repository.
func GetRepoRef(r *git.Repository, name string) (*plumbing.Reference, error) {
	if ref, err := r.Reference(plumbing.NewTagReferenceName(name), true); err == nil {
		return ref, nil
	}

	if ref, err := r.Reference(plumbing.NewBranchReferenceName(name), true); err == nil {
		return ref, nil
	}

	if ref, err := r.Reference(plumbing.NewRemoteReferenceName("origin", name), true); err == nil {
		return ref, nil
	}

	if ref, err := r.Reference(plumbing.NewNoteReferenceName(name), true); err == nil {
		return ref, nil
	}

	if revision, err := r.ResolveRevision(plumbing.Revision(name)); err == nil {
		return plumbing.NewHashReference(plumbing.ReferenceName(name), *revision), nil
	}

	return nil, fmt.Errorf("failed to get reference: %s", name)
}

// GetRepoVersions returns all versions of a git repository in descending order.
func GetRepoVersions(r *git.Repository) ([]*version.Version, error) {
	// TODO(michelia): logger
	//logger := log.WithName("vcs")

	tagRefs, err := r.Tags()
	if err != nil {
		return nil, err
	}

	var versions []*version.Version

	err = tagRefs.ForEach(func(ref *plumbing.Reference) error {
		v, verr := version.NewVersion(ref.Name().Short())
		if verr != nil {
			// TODO(michelia): logs
			//logger.Warnf("failed to parse tag %s: %v", ref.Name().Short(), verr)

		}

		if v != nil {
			versions = append(versions, v)
		}

		return nil
	})

	sort.Slice(versions, func(i, j int) bool {
		return versions[i].LessThan(versions[j])
	})

	return versions, err
}

// ListRepoTags list tags from the given repository, repository format is orgName/repoName or userName/repoName.
func ListRepoTags(ctx context.Context, client *scm.Client, repoName string) ([]*version.Version, error) {

	opts := scm.ListOptions{Size: 100}

	var versions []*version.Version

	for {
		tagRefs, meta, err := client.Git.ListTags(ctx, repoName, opts)
		if err != nil {
			return nil, err
		}

		for _, tag := range tagRefs {
			v, verr := version.NewVersion(tag.Name)
			if verr != nil {
				// TODO(michelia): log
				//logger.Warnf("failed to parse tag %s: %v", ref.Name().Short(), verr)
			}

			if v != nil {
				versions = append(versions, v)
			}
		}

		opts.Page = meta.Page.Next
		opts.URL = meta.Page.NextURL

		if opts.Page == 0 && opts.URL == "" {
			break
		}
	}

	return versions, nil
}
