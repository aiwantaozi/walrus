package ensurer

import (
	"context"
	"fmt"

	"github.com/seal-io/walrus/pkg/templates/ensurers/vcsrepo"

	walruscore "github.com/seal-io/walrus/pkg/apis/walruscore/v1"
)

type Ensurer interface {
	Ensure(ctx context.Context, template *walruscore.Template) error
}

// Ensure sync the template and it's versions.
func Ensure(ctx context.Context, tmpl *walruscore.Template) error {
	var e Ensurer
	switch {
	default:
		return fmt.Errorf("unsupport template format %s", tmpl.Spec.TemplateFormat)
	case tmpl.Spec.VCSRepository != nil:
		e = vcsrepo.New()
	}

	return e.Ensure(ctx, tmpl)
}
