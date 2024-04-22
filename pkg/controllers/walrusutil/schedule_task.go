package walrusutil

import (
	"context"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/go-multierror"
	"github.com/seal-io/utils/pools/gopool"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/klog/v2"
	ctrl "sigs.k8s.io/controller-runtime"
	ctrlbuilder "sigs.k8s.io/controller-runtime/pkg/builder"
	ctrlcli "sigs.k8s.io/controller-runtime/pkg/client"
	ctrlevent "sigs.k8s.io/controller-runtime/pkg/event"
	ctrllog "sigs.k8s.io/controller-runtime/pkg/log"
	ctrlpredicate "sigs.k8s.io/controller-runtime/pkg/predicate"
	ctrlreconcile "sigs.k8s.io/controller-runtime/pkg/reconcile"

	walrusutil "github.com/seal-io/walrus/pkg/apis/walrusutil/v1"
	"github.com/seal-io/walrus/pkg/controller"
	"github.com/seal-io/walrus/pkg/kubeclientset"
	"github.com/seal-io/walrus/pkg/scheduletasks/cron"
	"github.com/seal-io/walrus/pkg/scheduletasks/processor"
	"github.com/seal-io/walrus/pkg/systemkuberes"
)

// ScheduleTaskReconciler reconciles a v1.ScheduleTask object.
type ScheduleTaskReconciler struct {
	client ctrlcli.Client
}

var _ ctrlreconcile.Reconciler = (*ScheduleTaskReconciler)(nil)

// Reconcile reconciles the cronjob.
func (r *ScheduleTaskReconciler) Reconcile(ctx context.Context, req ctrl.Request) (res ctrl.Result, err error) {
	logger := ctrllog.FromContext(ctx)

	// Fetch.
	obj := new(walrusutil.ScheduleTask)
	err = r.client.Get(ctx, req.NamespacedName, obj)
	if err != nil {
		logger.Error(err, "fetch schedule task")
		return ctrl.Result{}, ctrlcli.IgnoreNotFound(err)
	}

	// Skip deletion.
	if obj.DeletionTimestamp != nil {
		return ctrl.Result{}, nil
	}

	var (
		now    = time.Now()
		objKey = ctrlcli.ObjectKeyFromObject(obj).String()
	)

	logger.Infof("reconciling schedule task %s, time: %s", objKey, now.String())

	requeueAfter, nextScheduleTime, err := r.reconcile(ctx, obj, now)
	uerr := r.updateStatus(ctx, obj)

	if err != nil || uerr != nil {
		logger.Error(multierror.Append(err, uerr), "failed to reconcile schedule task %s", req.String())
	}

	if requeueAfter != nil {
		logger.Infof("reconciled schedule task %s, time: %s, next time: %s", objKey, now.String(), nextScheduleTime.String())
		return ctrl.Result{RequeueAfter: *requeueAfter}, nil
	}

	return ctrl.Result{}, nil
}

func (r *ScheduleTaskReconciler) reconcile(ctx context.Context, obj *walrusutil.ScheduleTask, now time.Time) (*time.Duration, *time.Time, error) {
	logger := ctrllog.FromContext(ctx)

	objKey := ctrlcli.ObjectKeyFromObject(obj).String()

	// Calculate next schedule time.
	var (
		requeueAfter     *time.Duration
		mostRecentTime   *time.Time
		nextScheduleTime *time.Time
		scheduleInterval *time.Duration
	)
	{
		sched, err := cron.NewCron(obj)
		if err != nil {
			return nil, nil, err
		}

		requeueAfter, scheduleInterval, mostRecentTime, nextScheduleTime = cron.CalculateScheduleTiming(obj, now, sched)
	}

	// Clean
	taskTimeoutInterval := genTaskTimeoutInterval(obj, *scheduleInterval)
	{
		// clean timeout active task
		if obj.Status.ActiveTaskName != nil {
			unixTime, err := strconv.ParseInt(strings.TrimPrefix(*obj.Status.ActiveTaskName, obj.Name+"-"), 10, 64)
			if err != nil {
				// Clean invalid active task name.
				removeFromActive(obj, "")

				return requeueAfter, nextScheduleTime, fmt.Errorf("failed to parse unix time from active task name %s", *obj.Status.ActiveTaskName)
			}

			taskTimeoutTime := time.Unix(unixTime, 0).Add(taskTimeoutInterval)
			if obj.Status.ActiveTaskTimeoutTime != nil {
				taskTimeoutTime = obj.Status.ActiveTaskTimeoutTime.Time
			}

			if taskTimeoutTime.Before(now) {
				logger.V(4).Infof("clean timeout active task %s/%s, timeout time: %s", obj.Namespace, *obj.Status.ActiveTaskName, taskTimeoutTime.String())
				removeFromActive(obj, "")
			}
		}
	}

	// Precheck.
	{
		if obj.Spec.Suspend {
			logger.Info("skip task processing due to suspend configuration")
			return nil, nil, nil
		}

		if obj.Status.ActiveTaskName != nil {
			logger.Infof("skip task processing due to ongoing prior task %s isn't finished, "+
				"time %s, next time %s, active task timeout time %s",
				*obj.Status.ActiveTaskName, mostRecentTime, nextScheduleTime, obj.Status.ActiveTaskTimeoutTime)
			return requeueAfter, nextScheduleTime, nil
		}

		if obj.Status.LastScheduleTime == nil && obj.Spec.Await {
			logger.Info("skip task processing due to awaiting first schedule time")
			obj.Status.LastScheduleTime = &meta.Time{Time: now}
			return requeueAfter, nextScheduleTime, nil
		}

		err := processor.Existed(obj.Name)
		if err != nil {
			logger.Errorf(err, "skip task processing due to %s not found", obj.Name)
			return nil, nil, err
		}
	}

	// Process task.
	var (
		taskName = genTaskName(obj, *mostRecentTime)
		taskKey  = fmt.Sprintf("%s/%s", obj.Namespace, taskName)
	)
	{
		gopool.Go(func() {
			logger = klog.Background().WithValues("ScheduleTask", objKey).WithValues("task", taskKey).WithName("schedule-task-processor")

			subCtx, cancel := context.WithTimeout(context.Background(), taskTimeoutInterval)
			defer cancel()

			logger.V(4).Infof("running task processor, time: %s", mostRecentTime.String())

			err := processor.Process(subCtx, obj.Name, taskKey)
			completionTime := meta.Now()
			removeFromActive(obj, taskName)

			switch {
			case err != nil:
				logger.Errorf(err, "failed to run task processor, time: %s", completionTime)

				obj.Status.LastFailedTime = &completionTime
				obj.Status.LastFailedMessage = err.Error()
			default:
				logger.V(4).Infof("finished running task processor, time: %s", completionTime)

				obj.Status.LastSuccessfulTime = &completionTime
			}

			alignFn := func(as *walrusutil.ScheduleTask) (*walrusutil.ScheduleTask, bool, error) {
				if reflect.DeepEqual(as.Status, obj.Status) {
					return as, true, nil
				}

				removeFromActive(as, taskName)

				if obj.Status.LastSuccessfulTime != nil &&
					(as.Status.LastSuccessfulTime == nil || as.Status.LastSuccessfulTime.Before(obj.Status.LastSuccessfulTime)) {
					as.Status.LastSuccessfulTime = obj.Status.LastSuccessfulTime
				}

				if obj.Status.LastFailedTime != nil &&
					(as.Status.LastFailedTime == nil || as.Status.LastFailedTime.Before(obj.Status.LastFailedTime)) {
					as.Status.LastFailedTime = obj.Status.LastFailedTime
					as.Status.LastFailedMessage = obj.Status.LastFailedMessage
				}
				return as, false, nil
			}

			_, err = kubeclientset.UpdateStatusWithCtrlClient(ctx, r.client, obj, kubeclientset.WithUpdateStatusAlign(alignFn))
			if err != nil {
				logger.Errorf(err, "failed to update status of task: %v", err)
			}
		})

		obj.Status.Schedule = obj.Spec.Schedule
		obj.Status.ActiveTaskName = &taskName
		obj.Status.ActiveTaskTimeoutTime = &meta.Time{Time: mostRecentTime.Add(taskTimeoutInterval)}
		obj.Status.LastScheduleTime = &meta.Time{Time: *mostRecentTime}
	}

	logger.V(4).Infof("task %s for schedule task %s creation completed, time: %s", taskKey, objKey, mostRecentTime.String())

	return requeueAfter, nextScheduleTime, nil
}

// updateStatus updates the status of the cronjob.
func (r *ScheduleTaskReconciler) updateStatus(ctx context.Context, obj *walrusutil.ScheduleTask) error {
	alignFn := func(as *walrusutil.ScheduleTask) (*walrusutil.ScheduleTask, bool, error) {
		if reflect.DeepEqual(as.Status, obj.Status) {
			return as, true, nil
		}

		as.Status = obj.Status
		return as, false, nil
	}

	_, err := kubeclientset.UpdateStatusWithCtrlClient(ctx, r.client, obj, kubeclientset.WithUpdateStatusAlign(alignFn))
	return err
}

func genTaskTimeoutInterval(obj *walrusutil.ScheduleTask, scheduleInterval time.Duration) time.Duration {
	// the default and minimum timeout, while the job always timeout, which means the interval should be changed.
	timeout := scheduleInterval * 3
	const (
		// The maximum timeout is 6 hours.
		atMost = 6 * time.Hour
	)

	if obj.Spec.ActiveTaskDeadlineSeconds != nil {
		timeout = time.Duration(*obj.Spec.ActiveTaskDeadlineSeconds) * time.Second
	}

	if timeout > atMost {
		timeout = atMost
	}

	return timeout
}

// genTaskName generates a job name.
func genTaskName(obj *walrusutil.ScheduleTask, scheduledTime time.Time) string {
	return fmt.Sprintf("%s-%d", obj.Name, scheduledTime.Unix())
}

// isActive checks if a task is the active one.
func isActive(obj *walrusutil.ScheduleTask, name string) bool {
	return obj != nil && obj.Status.ActiveTaskName != nil && *obj.Status.ActiveTaskName == name
}

// removeFromActive removes a task from the active one,
// while the name is empty, will reset active task name and timeout.
func removeFromActive(obj *walrusutil.ScheduleTask, name string) {
	if obj == nil {
		return
	}

	if name == "" || isActive(obj, name) {
		obj.Status.ActiveTaskName = nil
		obj.Status.ActiveTaskTimeoutTime = nil
	}
}

func (r *ScheduleTaskReconciler) SetupController(_ context.Context, opts controller.SetupOptions) error {
	r.client = opts.Manager.GetClient()

	return ctrl.NewControllerManagedBy(opts.Manager).
		For(
			&walrusutil.ScheduleTask{},
			ctrlbuilder.WithPredicates(ctrlpredicate.GenerationChangedPredicate{})).
		WithEventFilter(namespaceFilter()).
		Complete(r)
}

func namespaceFilter() ctrlpredicate.Predicate {
	return ctrlpredicate.Funcs{
		CreateFunc: func(e ctrlevent.CreateEvent) bool {
			return e.Object.GetNamespace() == systemkuberes.SystemNamespaceName
		},
		UpdateFunc: func(e ctrlevent.UpdateEvent) bool {
			return e.ObjectNew.GetNamespace() == systemkuberes.SystemNamespaceName
		},
		DeleteFunc: func(e ctrlevent.DeleteEvent) bool {
			return e.Object.GetNamespace() == systemkuberes.SystemNamespaceName
		},
	}
}
