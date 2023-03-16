package bus

import (
	"context"

	"github.com/seal-io/seal/pkg/bus/applicationrevision"
	"github.com/seal-io/seal/pkg/bus/connector"
	"github.com/seal-io/seal/pkg/bus/module"
	"github.com/seal-io/seal/pkg/bus/setting"
	"github.com/seal-io/seal/pkg/costs/deployer"
	"github.com/seal-io/seal/pkg/cron"
	"github.com/seal-io/seal/pkg/dao/model"
	"github.com/seal-io/seal/pkg/modules"
	"github.com/seal-io/seal/pkg/platformtf"
)

type SetupOptions struct {
	ModelClient model.ClientSet
}

func Setup(ctx context.Context, opts SetupOptions) error {
	if err := module.AddSubscriber("module-sync-schema", modules.SyncSchema); err != nil {
		return err
	}

	if err := connector.AddSubscriber("deployer-ensure-cost-tools", deployer.EnsureCostTools); err != nil {
		return err
	}
	if err := connector.AddSubscriber("deployer-sync-cost-custom-pricing", deployer.SyncCostCustomPricing); err != nil {
		return err
	}

	if err := applicationrevision.AddSubscriber("terraform-sync-application-revision-status", platformtf.SyncApplicationRevisionStatus); err != nil {
		return err
	}

	if err := setting.AddSubscriber("cron-sync", cron.Sync); err != nil {
		return err
	}

	return nil
}