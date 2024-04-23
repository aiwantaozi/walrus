package walruscore

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/validation/field"
	ctrlcli "sigs.k8s.io/controller-runtime/pkg/client"
	ctrlwebhook "sigs.k8s.io/controller-runtime/pkg/webhook"
	ctrladmission "sigs.k8s.io/controller-runtime/pkg/webhook/admission"

	walruscore "github.com/seal-io/walrus/pkg/apis/walruscore/v1"
	"github.com/seal-io/walrus/pkg/apistatus"
	"github.com/seal-io/walrus/pkg/systemmeta"
	"github.com/seal-io/walrus/pkg/templates/api"
	"github.com/seal-io/walrus/pkg/templates/kubehelper"
	"github.com/seal-io/walrus/pkg/webhook"
)

// SchemaWebhook hooks a v1.Schema object.
//
// nolint: lll
// nolint: lll
// +k8s:webhook-gen:mutating:group="walruscore.seal.io",version="v1",resource="schemas",scope="Namespaced"
// +k8s:webhook-gen:mutating:operations=["UPDATE"],failurePolicy="Fail",sideEffects="None",matchPolicy="Equivalent",timeoutSeconds=10
// +k8s:webhook-gen:validating:group="walruscore.seal.io",version="v1",resource="schemas",scope="Namespaced"
// +k8s:webhook-gen:validating:operations=["UPDATE"],failurePolicy="Fail",sideEffects="None",matchPolicy="Equivalent",timeoutSeconds=10
type SchemaWebhook struct {
	webhook.DefaultCustomValidator

	client ctrlcli.Client
}

func (r *SchemaWebhook) SetupWebhook(_ context.Context, opts webhook.SetupOptions) (runtime.Object, error) {
	r.client = opts.Manager.GetClient()

	return &walruscore.Schema{}, nil
}

var (
	_ ctrlwebhook.CustomValidator = (*SchemaWebhook)(nil)
	_ ctrlwebhook.CustomDefaulter = (*SchemaWebhook)(nil)
)

func (r *SchemaWebhook) ValidateUpdate(ctx context.Context, oldObj, newObj runtime.Object) (ctrladmission.Warnings, error) {
	s := newObj.(*walruscore.Schema)

	if s.Status.Value.Size() == 0 {
		return nil, field.Invalid(
			field.NewPath("status.value"), "", "empty schema")
	}

	var ws api.WrapSchema
	err := json.Unmarshal(s.Status.Value.Raw, &ws)
	if err != nil {
		return nil, field.Invalid(
			field.NewPath("status.value"), "", err.Error())
	}

	err = ws.Validate()
	if err != nil {
		return nil, field.Invalid(
			field.NewPath("status.value"), "", err.Error())
	}

	return nil, nil
}

func (r *SchemaWebhook) Default(ctx context.Context, obj runtime.Object) error {
	s := obj.(*walruscore.Schema)

	// Skip template schema and original ui schema.
	if strings.HasSuffix(s.Name, walruscore.NameSuffixTemplateSchema) || strings.HasSuffix(s.Name, walruscore.NameSuffixOriginalUISchema) {
		return nil
	}

	switch {
	case systemmeta.DescribeResourceNote(s, kubehelper.SchemaUserEditedNote) == "false":
		// System edit.
	case apistatus.SchemaStatusReset.IsTrue(obj):
		// User reset.
		var (
			ouis    walruscore.Schema
			ouisKey = ctrlcli.ObjectKey{
				Namespace: s.Namespace,
				Name:      strings.TrimSuffix(s.Name, walruscore.NameSuffixUISchema) + walruscore.NameSuffixOriginalUISchema,
			}
		)
		err := r.client.Get(ctx, ouisKey, &ouis)
		if err != nil {
			return fmt.Errorf("failed to find %s: %w", ouisKey.String(), err)
		}

		apistatus.SchemaStatusReset.False(s, "", "")
		systemmeta.UnnoteResource(s)
		s.Status.Value = ouis.Status.Value

	default:
		// User edit.
		systemmeta.NoteResource(s, "", map[string]string{kubehelper.SchemaUserEditedNote: "true"})
	}

	return nil
}
