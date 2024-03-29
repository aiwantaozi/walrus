package walruscore

import (
	"context"
	"fmt"

	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	ctrllog "sigs.k8s.io/controller-runtime/pkg/log"
	ctrlreconcile "sigs.k8s.io/controller-runtime/pkg/reconcile"

	"github.com/seal-io/utils/pools/gopool"
	walruscore "github.com/seal-io/walrus/pkg/apis/walruscore/v1"
	"github.com/seal-io/walrus/pkg/controller"
	tmpllister "github.com/seal-io/walrus/pkg/templates/lister"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
)

// CatalogReconciler reconciles a v1.Catalog object.
type CatalogReconciler struct {
	client client.Client
}

var _ ctrlreconcile.Reconciler = (*CatalogReconciler)(nil)

func (r *CatalogReconciler) Reconcile(ctx context.Context, req ctrl.Request) (res ctrl.Result, err error) {
	logger := ctrllog.FromContext(ctx)

	// TODO(michelia)
	logger.V(6).Info("reconciling")

	var obj walruscore.Catalog
	if err := r.client.Get(ctx, req.NamespacedName, &obj); err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	if obj.DeletionTimestamp != nil {
		return ctrl.Result{}, nil
	}

	// Skip catalog already reconciling.
	if walruscore.CatalogStatusReady.IsUnknown(obj) {
		return ctrl.Result{}, nil
	}

	// Skip catalog finished reconciling.
	if (walruscore.CatalogStatusReady.IsTrue(obj) || walruscore.CatalogStatusReady.IsFalse(obj)) && !walruscore.CatalogStatusRefresh.IsUnknown(obj) {
		return ctrl.Result{}, nil
	}

	// Initialize catalog status.
	walruscore.CatalogStatusReady.Unknown(obj, "")
	if walruscore.CatalogStatusRefresh.IsUnknown(obj) {
		walruscore.CatalogStatusRefresh.False(obj, "")
	}
	err = r.client.Update(ctx, &obj)
	if err != nil {
		return ctrl.Result{}, err
	}

	// Reconcile catalog.
	defer func() {
		// Get the latest catalog object.
		getErr := r.client.Get(ctx, req.NamespacedName, &obj)
		if getErr != nil && err == nil {
			err = getErr
		}

		if err != nil {
			walruscore.CatalogStatusReady.False(obj, err.Error())
		} else {
			walruscore.CatalogStatusReady.True(obj, "")
		}

		_ = r.client.Update(ctx, &obj)
	}()

	tmpls, err := tmpllister.List(ctx, obj)
	if err != nil {
		walruscore.CatalogStatusReady.False(obj, err.Error())
		return ctrl.Result{}, err
	}

	// Ensure templates.
	wg := gopool.GroupWithContextIn(ctx)
	for i := range tmpls {
		wg.Go(func(ctx context.Context) error {
			err = r.ensureTemplate(ctx, tmpls[i])
			if err != nil {
				return fmt.Errorf("failed to ensure template %s/%s: %w", tmpls[i].Namespace, tmpls[i].Name, err)
			}

			return nil
		})
	}

	err = wg.Wait()
	if err != nil {
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

func (r *CatalogReconciler) ensureTemplate(ctx context.Context, tmpl walruscore.Template) error {
	var (
		namespaceName = types.NamespacedName{
			Namespace: tmpl.Namespace,
			Name:      tmpl.Name,
		}
		existed walruscore.Template
	)

	err := r.client.Get(ctx, namespaceName, &existed)
	if err != nil {
		if !apierrors.IsNotFound(err) {
			return err
		}

		// Create template.
		err = r.client.Create(ctx, &tmpl, &client.CreateOptions{})
		if err != nil && !apierrors.IsAlreadyExists(err) {
			return err
		}

		if err == nil {
			return nil
		}
		return nil
	}

	// Update template.
	copied := existed.DeepCopy()
	copied.Spec.Description = tmpl.Spec.Description
	copied.Status = tmpl.Status
	return r.client.Update(ctx, copied)
}

func (r *CatalogReconciler) SetupController(_ context.Context, opts controller.SetupOptions) error {
	r.client = opts.Manager.GetClient()

	return ctrl.NewControllerManagedBy(opts.Manager).
		For(&walruscore.Catalog{}).
		Complete(r)
}
