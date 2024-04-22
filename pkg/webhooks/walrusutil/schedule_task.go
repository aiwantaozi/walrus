package walrusutil

import (
	"context"
	"time"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/validation/field"
	ctrlwebhook "sigs.k8s.io/controller-runtime/pkg/webhook"
	ctrladmission "sigs.k8s.io/controller-runtime/pkg/webhook/admission"

	walrusutil "github.com/seal-io/walrus/pkg/apis/walrusutil/v1"
	"github.com/seal-io/walrus/pkg/scheduletasks/cron"
	"github.com/seal-io/walrus/pkg/scheduletasks/processors/catalog"
	"github.com/seal-io/walrus/pkg/webhook"
)

// ScheduleTaskWebhook hooks a v1.ScheduleTask object.
//
// +k8s:webhook-gen:validating:group="walrusutil.seal.io",version="v1",resource="scheduletasks",scope="Namespaced"
// +k8s:webhook-gen:validating:operations=["CREATE","UPDATE"],failurePolicy="Fail",sideEffects="None",matchPolicy="Equivalent",timeoutSeconds=10
type ScheduleTaskWebhook struct {
	webhook.DefaultCustomValidator
}

func (r *ScheduleTaskWebhook) SetupWebhook(_ context.Context, opts webhook.SetupOptions) (runtime.Object, error) {
	return &walrusutil.ScheduleTask{}, nil
}

var _ ctrlwebhook.CustomValidator = (*ScheduleTaskWebhook)(nil)

func (r *ScheduleTaskWebhook) ValidateCreate(ctx context.Context, obj runtime.Object) (ctrladmission.Warnings, error) {
	return r.validateScheduleTask(obj)
}

func (r *ScheduleTaskWebhook) ValidateUpdate(ctx context.Context, oldObj, newObj runtime.Object) (ctrladmission.Warnings, error) {
	return r.validateScheduleTask(newObj)
}

func (r *ScheduleTaskWebhook) validateScheduleTask(obj runtime.Object) (ctrladmission.Warnings, error) {
	cj := obj.(*walrusutil.ScheduleTask)

	validators := make([]cron.CronValidator, 0)
	if cj.Name == catalog.Name {
		validators = append(validators, cron.AtLeastDuration(30*time.Minute))
	}

	_, err := cron.NewCron(cj, validators...)
	if err != nil {
		return nil, field.Invalid(
			field.NewPath("spec.schedule"), cj.Spec.Schedule, err.Error())
	}
	return nil, nil
}
