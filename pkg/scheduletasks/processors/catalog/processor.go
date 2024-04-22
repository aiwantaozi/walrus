package catalog

import (
	"context"
	"fmt"
	"reflect"

	"go.uber.org/multierr"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/klog/v2"

	walruscore "github.com/seal-io/walrus/pkg/apis/walruscore/v1"
	"github.com/seal-io/walrus/pkg/apistatus"
	"github.com/seal-io/walrus/pkg/kubeclientset"
	"github.com/seal-io/walrus/pkg/system"
	"github.com/seal-io/walrus/pkg/systemsetting"
)

const Name = "catalog-sync"

type Processor struct{}

func (in *Processor) Name() string {
	return Name
}

func (in *Processor) Process(ctx context.Context, id string) error {
	logger := klog.Background().WithName("processor").WithName(in.Name()).WithValues("task", id)
	cli := system.LoopbackKubeClient.Get()

	// Disable sync catalog.
	enable, err := systemsetting.EnableSyncCatalog.ValueBool(ctx)
	if err != nil {
		return err
	}
	if !enable {
		logger.Infof("catalog sync is disabled")
		return nil
	}

	catalogs, err := cli.WalruscoreV1().Catalogs(meta.NamespaceAll).List(ctx, meta.ListOptions{})
	if err != nil {
		return err
	}

	if len(catalogs.Items) == 0 {
		return nil
	}

	// Merge the errors to return them all at once,
	// instead of returning the first error.
	var berr error

	for i := range catalogs.Items {
		c := catalogs.Items[i]

		if apistatus.CatalogConditionRefresh.IsUnknown(&c) {
			continue
		}

		logger.V(4).Infof("syncing catalog %s/%s", c.Namespace, c.Name)

		apistatus.CatalogConditionRefresh.Reset(&c, "Refreshing catalog templates")
		alignFn := func(as *walruscore.Catalog) (*walruscore.Catalog, bool, error) {
			if reflect.DeepEqual(as.Status, c.Status) {
				return as, true, nil
			}

			as.Status = c.Status
			return as, false, nil
		}

		_, err = kubeclientset.UpdateStatus(ctx, cli.WalruscoreV1().Catalogs(c.Namespace), &c,
			kubeclientset.WithUpdateStatusAlign(alignFn))
		if err != nil {
			berr = multierr.Append(berr,
				fmt.Errorf("error syncing catalog %s/%s: %w",
					c.Namespace, c.Name, err))
		}
	}

	return berr
}
