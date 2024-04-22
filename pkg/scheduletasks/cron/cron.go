package cron

import (
	"fmt"
	"strings"
	"time"

	"github.com/robfig/cron/v3"

	walrusutil "github.com/seal-io/walrus/pkg/apis/walrusutil/v1"
)

var defaultCronParser = cron.NewParser(
	cron.Second | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow | cron.Descriptor,
)

// NewCron creates the cron scheduler.
func NewCron(obj *walrusutil.ScheduleTask, validators ...CronValidator) (cron.Schedule, error) {
	expr, err := formatScheduleExpression(obj)
	if err != nil {
		return nil, err
	}

	sched, err := defaultCronParser.Parse(expr)
	if err != nil {
		return nil, err
	}

	if len(validators) == 0 {
		validators = append(validators, AtLeastDuration(1*time.Second))
	}

	for i := range validators {
		err = validators[i](obj, sched)
		if err != nil {
			return nil, err
		}
	}

	return sched, nil
}

// formatScheduleExpression formats the schedule expression with timezone.
func formatScheduleExpression(obj *walrusutil.ScheduleTask) (string, error) {
	var (
		loc   string
		sched = obj.Spec.Schedule
	)
	switch {
	case strings.Contains(obj.Spec.Schedule, "TZ"):
		parts := strings.Fields(obj.Spec.Schedule)
		loc = parts[0][len("TZ="):]
	case obj.Spec.TimeZone != nil:
		loc = *obj.Spec.TimeZone
		sched = fmt.Sprintf("TZ=%s %s", *obj.Spec.TimeZone, obj.Spec.Schedule)
	}

	if loc != "" {
		if _, err := time.LoadLocation(loc); err != nil {
			return "", fmt.Errorf("could not parse location %s: %w", loc, err)
		}
	}

	return sched, nil
}

type CronValidator = func(obj *walrusutil.ScheduleTask, schedule cron.Schedule) error

// AtLeastDuration checks if the duration between two schedules is valid.
func AtLeastDuration(d time.Duration) CronValidator {
	return func(obj *walrusutil.ScheduleTask, schedule cron.Schedule) error {
		earliestTime := obj.ObjectMeta.CreationTimestamp.Time
		if obj.Status.LastScheduleTime != nil {
			earliestTime = obj.Status.LastScheduleTime.Time
		}

		next := schedule.Next(earliestTime)

		// Check if the duration two schedules is valid.
		// cron.ParseStandard("0 59 23 31 2 *") represent 11:59 PM on the 31st day of February every year,
		// which is invalid, duration will be 0.
		duration := int64(schedule.Next(next).Sub(next).Round(time.Second).Seconds())
		if duration < 1 {
			return fmt.Errorf("time difference between two schedules is less than 1 second")
		}

		if duration < int64(d.Seconds()) {
			return fmt.Errorf("cron expression %s is too short, at least %v", obj.Spec.Schedule, d)
		}

		return nil
	}
}
