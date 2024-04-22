package v1

import (
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// ScheduleTask is the schema for the scheduletasks API.
//
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:crd-gen:resource:scope="Namespaced",subResources=["status"]
type ScheduleTask struct {
	meta.TypeMeta   `json:",inline"`
	meta.ObjectMeta `json:"metadata,omitempty"`

	Spec   ScheduleTaskSpec   `json:"spec"`
	Status ScheduleTaskStatus `json:"status,omitempty"`
}

var _ runtime.Object = (*ScheduleTask)(nil)

// ScheduleTaskSpec defines the desired state of ScheduleTask.
type ScheduleTaskSpec struct {
	// Schedule is a schedule to run the task in Cron format,
	// see https://en.wikipedia.org/wiki/Cron.
	Schedule string `json:"schedule"`

	// Suspend, when set to true, prevents new task from starting,
	// it does not affect already started task.
	Suspend bool `json:"suspend,omitempty"`

	// Await is a flag that, when set to true,
	// indicates that the task should wait until the scheduled time to execute when creating the scheduled task.
	Await bool `json:"await,omitempty"`

	// TimeZone specifies the timezone used for calculating the cron schedule,
	// if not provided, the default is the local time of the machine,
	// see https://en.wikipedia.org/wiki/List_of_tz_database_time_zones.
	TimeZone *string `json:"timeZone,omitempty"`

	// ActiveTaskDeadlineSeconds specifies the active task timeout duration in seconds,
	// if not specified, the default timeout is three times the schedule interval, with a maximum timeout of 6 hours.
	ActiveTaskDeadlineSeconds *int64 `json:"activeTaskDeadlineSeconds,omitempty"`
}

// ScheduleTaskStatus defines the observed state of ScheduleTask.
type ScheduleTaskStatus struct {
	// Schedule is the current schedule to run the task in Cron format,
	// see https://en.wikipedia.org/wiki/Cron.
	Schedule string `json:"schedule,omitempty"`

	// ActiveTaskName is the currently running task names.
	ActiveTaskName *string `json:"activeTaskName,omitempty"`

	// ActiveTaskTimeoutTime is the time currently running task timeout.
	ActiveTaskTimeoutTime *meta.Time `json:"activeTaskTimeoutTime,omitempty"`

	// LastScheduleTime is the last time the task was successfully scheduled.
	LastScheduleTime *meta.Time `json:"lastScheduleTime,omitempty"`

	// LastSuccessfulTime is the last time the task successfully completed.
	LastSuccessfulTime *meta.Time `json:"lastSuccessfulTime,omitempty"`

	// LastFailedTime is the last time the task failed.
	LastFailedTime *meta.Time `json:"lastFailedTime,omitempty"`

	// LastFailedMessage is the last error message of the failed task.
	LastFailedMessage string `json:"lastFailedMessage,omitempty"`
}

// ScheduleTaskList holds the list of ScheduleTask.
//
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type ScheduleTaskList struct {
	meta.TypeMeta `json:",inline"`
	meta.ListMeta `json:"metadata,omitempty"`

	Items []ScheduleTask `json:"items"`
}

var _ runtime.Object = (*ScheduleTaskList)(nil)
