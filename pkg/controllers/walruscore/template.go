package walruscore

import (
	"context"

	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	ctrllog "sigs.k8s.io/controller-runtime/pkg/log"
	ctrlreconcile "sigs.k8s.io/controller-runtime/pkg/reconcile"

	walruscore "github.com/seal-io/walrus/pkg/apis/walruscore/v1"
	"github.com/seal-io/walrus/pkg/controller"
	tmplensurer "github.com/seal-io/walrus/pkg/templates/ensurer"
)

// TemplateReconciler reconciles a v1.Template object.
type TemplateReconciler struct {
	client client.Client
}

var _ ctrlreconcile.Reconciler = (*TemplateReconciler)(nil)

func (r *TemplateReconciler) Reconcile(ctx context.Context, req ctrl.Request) (res ctrl.Result, err error) {
	logger := ctrllog.FromContext(ctx)

	// TODO(michelia)
	logger.V(6).Info("reconciling")

	var obj walruscore.Template
	if err := r.client.Get(ctx, req.NamespacedName, &obj); err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	if obj.DeletionTimestamp != nil {
		return ctrl.Result{}, nil
	}

	// Skip template already reconciling.
	if walruscore.TemplateStatusReady.IsUnknown(obj) {
		return ctrl.Result{}, nil
	}

	// Skip template finished reconciling.
	if (walruscore.TemplateStatusReady.IsTrue(obj) || walruscore.TemplateStatusReady.IsFalse(obj)) &&
		!walruscore.TemplateStatusRefresh.IsUnknown(obj) {
		return ctrl.Result{}, nil
	}

	// Initialize template status.
	walruscore.TemplateStatusReady.Unknown(obj, "")
	if walruscore.TemplateStatusRefresh.IsUnknown(obj) {
		walruscore.TemplateStatusRefresh.False(obj, "")
	}
	err = r.client.Update(ctx, &obj)
	if err != nil {
		return ctrl.Result{}, err
	}

	// Reconcile template.
	defer func() {
		// Get the latest template object.
		getErr := r.client.Get(ctx, req.NamespacedName, &obj)
		if getErr != nil && err == nil {
			err = getErr
		}

		if err != nil {
			walruscore.TemplateStatusReady.False(obj, err.Error())
		} else {
			walruscore.TemplateStatusReady.True(obj, "")
		}

		_ = r.client.Update(ctx, &obj)
	}()

	err = tmplensurer.Ensure(ctx, &obj)
	if err != nil {
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

func (r *TemplateReconciler) SetupController(_ context.Context, opts controller.SetupOptions) error {
	r.client = opts.Manager.GetClient()

	return ctrl.NewControllerManagedBy(opts.Manager).
		For(&walruscore.Template{}).
		Complete(r)
}
