package vcs

import (
	"fmt"
	"net/url"
	"strings"
)

// GetOrgFromGitURL parses the organization from the given URL.
func GetOrgFromGitURL(source string) (string, error) {
	u, err := url.Parse(source)
	if err != nil {
		return "", err
	}

	parts := strings.Split(u.Path, "/")
	if len(parts) < 2 {
		return "", fmt.Errorf("invalid git url")
	}

	return strings.TrimPrefix(u.Path, "/"), nil
}
