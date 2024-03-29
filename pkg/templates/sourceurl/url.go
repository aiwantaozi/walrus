package sourceurl

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-git/go-git/v5/plumbing/transport"

	"github.com/seal-io/walrus/pkg/vcs/driver/gitee"
	"github.com/seal-io/walrus/pkg/vcs/driver/github"
	"github.com/seal-io/walrus/pkg/vcs/driver/gitlab"
)

type SourceURL struct {
	// Link is the link of the repository.
	Link string `json:"link"`
	// StorageType is the storage type of the content. E.G: github, gitlab, http.
	StorageType string `json:"storageType"`
	// Reference is the reference of the source, valid when storage type is version control system.
	// E.G: main, dev, v0.0.1.
	Reference string `json:"reference"`
	// SubPath is the sub path of the source, valid while source include //.
	SubPath string `json:"subPath"`

	// Namespace is the namespace of repository, valid while storage type is vcs.
	Namespace string `json:"namespace"`
	// Name is the name of repository, valid while storage type is vcs.
	Name string `json:"name"`
}

// FileRawURL returns raw URL of a file in a git repository.
func (s *SourceURL) FileRawURL(file string) (string, error) {
	if file == "" {
		return "", nil
	}

	endpoint, err := transport.NewEndpoint(s.Link)
	if err != nil {
		return "", err
	}

	var (
		githubRawHost = "raw.githubusercontent.com"
		gitlabRawHost = "gitlab.com"
		giteeRawHost  = "gitee.com"
		ref           = "HEAD"
	)

	if s.Reference != "" {
		ref = s.Reference
	}

	switch endpoint.Host {
	case "github.com":
		return fmt.Sprintf("https://%s/%s/%s/%s/%s", githubRawHost, s.Namespace, s.Name, ref, file), nil
	case "gitlab.com":
		return fmt.Sprintf("https://%s/%s/%s/-/raw/%s/%s", gitlabRawHost, s.Namespace, s.Name, ref, file), nil
	case "gitee.com":
		return fmt.Sprintf("https://%s/%s/%s/raw/%s/%s", giteeRawHost, s.Namespace, s.Name, ref, file), nil
	}

	if s.StorageType == gitlab.Driver {
		return fmt.Sprintf("%s/-/raw/%s/%s", endpoint.String(), ref, file), nil
	}

	return "", nil
}

// TODO(michelia): better
// ParseURLToSourceURL parses a raw URL to a source url.
func ParseURLToSourceURL(rawURL string) (*SourceURL, error) {
	// Trim git:: prefix.
	rawURL = strings.TrimPrefix(rawURL, "git::")
	ref := ""
	subPath := ""

	endpoint, err := transport.NewEndpoint(rawURL)
	if err != nil {
		return nil, err
	}

	path := endpoint.Path

	// Get ref from path.
	if strings.Contains(path, "?ref=") {
		parts := strings.Split(endpoint.Path, "?ref=")
		ref = parts[1]
		path = strings.TrimSuffix(path, "?ref="+ref)
		rawURL = strings.TrimSuffix(rawURL, "?ref="+ref)
	}

	// Get sub path from path.
	if strings.Contains(path, "//") {
		paths := strings.Split(path, "//")
		if len(paths) > 2 {
			return nil, errors.New("git url contains more than one //")
		}
		subPath = paths[1]

		path = strings.TrimSuffix(path, "//"+subPath)
		rawURL = strings.TrimSuffix(rawURL, "//"+subPath)
	}

	// Trim .git suffix.
	var (
		namespace, name string
	)
	path = strings.TrimSuffix(path, ".git")
	switch endpoint.Protocol {
	case "https", "http":
		parts := strings.Split(path, "/")
		if len(parts) < 3 {
			return nil, errors.New("invalid repository path")
		}
		namespace = parts[1]
		name = parts[2]
	}

	var storageType string
	switch endpoint.Host {
	case "github.com":
		storageType = github.Driver
	case "gitlab.com":
		storageType = gitlab.Driver
	case "gitee.com":
		storageType = gitee.Driver
	}

	return &SourceURL{
		Link:        rawURL,
		Reference:   ref,
		StorageType: storageType,
		SubPath:     subPath,
		Namespace:   namespace,
		Name:        name,
	}, nil
}
