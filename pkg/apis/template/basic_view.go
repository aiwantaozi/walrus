package template

import (
	"errors"
	"fmt"

	"github.com/hashicorp/go-getter"

	"github.com/seal-io/walrus/pkg/apis/runtime"
	"github.com/seal-io/walrus/pkg/dao/model"
	"github.com/seal-io/walrus/pkg/dao/model/predicate"
	"github.com/seal-io/walrus/pkg/dao/model/template"
	"github.com/seal-io/walrus/pkg/dao/types/object"
	"github.com/seal-io/walrus/pkg/vcs"
	"github.com/seal-io/walrus/utils/validation"
)

type (
	CreateRequest struct {
		model.TemplateCreateInput `path:",inline" json:",inline"`
	}

	CreateResponse = *model.TemplateOutput
)

func (r *CreateRequest) Validate() error {
	if err := r.TemplateCreateInput.Validate(); err != nil {
		return err
	}

	if err := validation.IsDNSLabel(r.Name); err != nil {
		return err
	}

	if _, err := getter.Detect(r.Source, "", getter.Detectors); err != nil {
		return fmt.Errorf("invalid source: %w", err)
	}

	if _, err := vcs.ParseURLToRepo(r.Source); err != nil {
		return fmt.Errorf("invalid source: %w", err)
	}

	return nil
}

type (
	GetRequest = model.TemplateQueryInput

	GetResponse = *model.TemplateOutput
)

type UpdateRequest struct {
	model.TemplateUpdateInput `path:",inline" json:",inline"`
}

func (r *UpdateRequest) Validate() error {
	if err := r.TemplateUpdateInput.Validate(); err != nil {
		return err
	}

	return nil
}

type DeleteRequest = model.TemplateDeleteInput

type (
	CollectionGetRequest struct {
		model.TemplateQueryInputs `path:",inline" query:",inline"`

		runtime.RequestCollection[
			predicate.Template, template.OrderOption,
		] `query:",inline"`

		CatalogIDs []object.ID `query:"catalogID,omitempty"`
		NonCatalog bool        `query:"nonCatalog,omitempty"`
		WithGlobal bool        `query:"withGlobal,omitempty"`
		IsService  *bool       `query:"isService,omitempty"`

		Stream *runtime.RequestUnidiStream
	}

	CollectionGetResponse = []*model.TemplateOutput
)

func (r *CollectionGetRequest) Validate() error {
	if err := r.TemplateQueryInputs.Validate(); err != nil {
		return err
	}

	if r.CatalogIDs != nil {
		for i := range r.CatalogIDs {
			if !r.CatalogIDs[i].Valid() {
				return errors.New("invalid catalog id")
			}
		}
	}

	return nil
}

func (r *CollectionGetRequest) SetStream(stream runtime.RequestUnidiStream) {
	r.Stream = &stream
}

type CollectionDeleteRequest = model.TemplateDeleteInputs
