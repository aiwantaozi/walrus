package walrus

import (
	"context"

	"k8s.io/apiserver/pkg/registry/rest"
	ctrlcli "sigs.k8s.io/controller-runtime/pkg/client"

	walrus "github.com/seal-io/walrus/pkg/apis/walrus/v1"
	"github.com/seal-io/walrus/pkg/extensionapi"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
)

// SchemaHandler handles v1.Schema objects.
type SchemaHandler struct {
	extensionapi.ObjectInfo
	extensionapi.GetOperation
	extensionapi.UpdateOperation

	Client ctrlcli.Client
}

func (h *SchemaHandler) SetupHandler(
	ctx context.Context,
	opts extensionapi.SetupOptions,
) (gvr schema.GroupVersionResource, srs map[string]rest.Storage, err error) {
	// Declare GVR.
	gvr = walrus.SchemeGroupVersionResource("schemas")

	// As storage.
	h.ObjectInfo = &walrus.Schema{}
	h.GetOperation = extensionapi.WithGet(h)
	h.UpdateOperation = extensionapi.WithUpdate(h)

	// Set client.
	h.Client = opts.Manager.GetClient()

	return
}

var (
	_ rest.Storage = (*SchemaHandler)(nil)
	_ rest.Getter  = (*SchemaHandler)(nil)
	_ rest.Updater = (*SchemaHandler)(nil)
	_ rest.Patcher = (*SchemaHandler)(nil)
)

func (h *SchemaHandler) Destroy() {
}

func (h *SchemaHandler) New() runtime.Object {
	return &walrus.Schema{}
}

func (h *SchemaHandler) NewList() runtime.Object {
	return &walrus.SchemaList{}
}

func (h *SchemaHandler) OnGet(ctx context.Context, key types.NamespacedName, opts ctrlcli.GetOptions) (runtime.Object, error) {
	// Get.
	s := &walrus.Schema{
		ObjectMeta: meta.ObjectMeta{
			Namespace: key.Namespace,
			Name:      key.Name,
		},
	}
	err := h.Client.Get(ctx, key, s)
	if err != nil {
		return nil, err
	}

	return s, nil
}

func (h *SchemaHandler) OnUpdate(ctx context.Context, obj, _ runtime.Object, opts ctrlcli.UpdateOptions) (runtime.Object, error) {
	s := obj.(*walrus.Schema)
	s.Status.Value = s.Spec.Value

	err := h.Client.Update(ctx, s)
	if err != nil {
		return nil, err
	}

	return s, nil
}
