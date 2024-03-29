package utils

import (
	"strings"

	"github.com/seal-io/utils/stringx"
)

// TODO(michelia): This function need to update template name could not include slash.
func NormalizeTemplateName(catalogName, repoName string) string {
	return stringx.Join("-", catalogName, strings.TrimPrefix(repoName, "terraform-"))
}

func NormalizeTemplateVersionConfigmapName(templateName, version, suffix string) string {
	return stringx.Join("-", templateName, version, suffix)
}

// TODO(michelia): since we don't add git:: prefix, need to add git:: prefix when deploy terraform.
func NormalizeTemplateURL(repoLink, ref string) string {
	repoLink += "?ref=" + ref
	return repoLink
}
