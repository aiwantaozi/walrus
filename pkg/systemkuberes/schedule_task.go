package systemkuberes

import (
	"context"
	"fmt"

	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/utils/pointer"

	walrusutil "github.com/seal-io/walrus/pkg/apis/walrusutil/v1"
	"github.com/seal-io/walrus/pkg/clients/clientset"
	"github.com/seal-io/walrus/pkg/kubeclientset"
	"github.com/seal-io/walrus/pkg/scheduletasks/processors/catalog"
	"github.com/seal-io/walrus/pkg/scheduletasks/processors/connector"
	"github.com/seal-io/walrus/pkg/scheduletasks/processors/resource"
	"github.com/seal-io/walrus/pkg/scheduletasks/processors/resourcecomponents"
	"github.com/seal-io/walrus/pkg/scheduletasks/processors/telemetry"
)

func InstallBuiltinScheduleTasks(ctx context.Context, cli clientset.Interface) error {
	scheduleTasks := []*walrusutil.ScheduleTask{
		catalogSyncScheduleTask(),
		connectorSyncScheduleTask(),
		resourceSyncScheduleTask(),
		resourceComponentsSyncScheduleTask(),
		telemetryReportScheduleTask(),
	}

	scheduleTaskCli := cli.WalrusutilV1().ScheduleTasks(SystemNamespaceName)
	for i := range scheduleTasks {
		_, err := kubeclientset.Create(ctx, scheduleTaskCli, scheduleTasks[i])
		if err != nil {
			return fmt.Errorf("install builtin scheduleTask: %w", err)
		}
	}

	return nil
}

// catalogSyncScheduleTask indicates the schedule task of catalog synchronization,
// default cron expression means sync at 1 o'clock evey day, and new cron must be at least 30 minutes,
// default active task deadline second is 5 minutes.
func catalogSyncScheduleTask() *walrusutil.ScheduleTask {
	return &walrusutil.ScheduleTask{
		ObjectMeta: meta.ObjectMeta{
			Name:      catalog.Name,
			Namespace: SystemNamespaceName,
		},
		Spec: walrusutil.ScheduleTaskSpec{
			Await:                     true,
			Schedule:                  "0 0 1 * * *",
			ActiveTaskDeadlineSeconds: pointer.Int64(60 * 5),
		},
	}
}

// connectorSyncScheduleTask indicates the schedule task of sync connector status,
// default cron expression means executing check every 5 minutes,
// default active task deadline second is 5 minutes.
func connectorSyncScheduleTask() *walrusutil.ScheduleTask {
	return &walrusutil.ScheduleTask{
		ObjectMeta: meta.ObjectMeta{
			Name:      connector.Name,
			Namespace: SystemNamespaceName,
		},
		Spec: walrusutil.ScheduleTaskSpec{
			Schedule:                  "0 */5 * ? * *",
			ActiveTaskDeadlineSeconds: pointer.Int64(60 * 5),
		},
	}
}

// resourceSyncScheduleTask indicates the schedule task of resource relation check,
// default cron expression means deploying every 30 seconds.
// default active task deadline second is 90 seconds.
func resourceSyncScheduleTask() *walrusutil.ScheduleTask {
	return &walrusutil.ScheduleTask{
		ObjectMeta: meta.ObjectMeta{
			Name:      resource.Name,
			Namespace: SystemNamespaceName,
		},
		Spec: walrusutil.ScheduleTaskSpec{
			Schedule:                  "*/30 * * ? * *",
			ActiveTaskDeadlineSeconds: pointer.Int64(30 * 3),
		},
	}
}

// resourceComponentsSyncScheduleTask indicates the schedule task of resource components status sync and discovery,
// default cron expression means stating every 1 minute,
// default active task deadline second is 90 seconds.
func resourceComponentsSyncScheduleTask() *walrusutil.ScheduleTask {
	return &walrusutil.ScheduleTask{
		ObjectMeta: meta.ObjectMeta{
			Name:      resourcecomponents.Name,
			Namespace: SystemNamespaceName,
		},
		Spec: walrusutil.ScheduleTaskSpec{
			Schedule:                  "*/30 * * ? * *",
			ActiveTaskDeadlineSeconds: pointer.Int64(30 * 3),
		},
	}
}

// telemetryReportScheduleTask indicates the schedule task of telemetry report,
// default cron expression means sync at 2 o'clock evey day,
// default active task deadline second is 5 minutes.
func telemetryReportScheduleTask() *walrusutil.ScheduleTask {
	return &walrusutil.ScheduleTask{
		ObjectMeta: meta.ObjectMeta{
			Name:      telemetry.Name,
			Namespace: SystemNamespaceName,
		},
		Spec: walrusutil.ScheduleTaskSpec{
			Await:                     true,
			Schedule:                  "0 0 2 * * *",
			ActiveTaskDeadlineSeconds: pointer.Int64(60 * 5),
		},
	}
}
