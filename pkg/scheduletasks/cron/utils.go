package cron

import (
	"time"

	"github.com/robfig/cron/v3"

	walrusutil "github.com/seal-io/walrus/pkg/apis/walrusutil/v1"
)

const (
	nextScheduleDelta = 100 * time.Millisecond
)

// CalculateScheduleTiming returns the time duration to requeue based on
// the schedule and next schedule time. It adds a 100ms padding to the next requeue to account
// for Network Time Protocol(NTP) time skews. If the time drifts the adjustment, which in most
// realistic cases should be around 100s, the job will still be executed without missing the schedule.
//
// Adapted from: https://github.com/kubernetes/kubernetes/blob/master/pkg/controller/cronjob/utils.go
func CalculateScheduleTiming(st *walrusutil.ScheduleTask, now time.Time, schedule cron.Schedule) (
	*time.Duration, *time.Duration, *time.Time, *time.Time,
) {
	scheduleInterval, mostRecentTime := mostRecentScheduleTime(st, now, schedule)

	timeToNextSchedule := schedule.Next(*mostRecentTime).Add(nextScheduleDelta).Sub(now)
	nextScheduledTime := now.Add(timeToNextSchedule)
	return &timeToNextSchedule, &scheduleInterval, mostRecentTime, &nextScheduledTime
}

// mostRecentScheduleTime returns:
//   - the duration between two schedules,
//   - the last schedule time or CronJob's creation time,
//   - the most recent time a Job should be created.
//
// Adapted from: https://github.com/kubernetes/kubernetes/blob/master/pkg/controller/cronjob/utils.go
func mostRecentScheduleTime(cj *walrusutil.ScheduleTask, now time.Time, schedule cron.Schedule) (time.Duration, *time.Time) {
	earliestTime := cj.ObjectMeta.CreationTimestamp.Time
	if cj.Status.LastScheduleTime != nil {
		earliestTime = cj.Status.LastScheduleTime.Time
	}

	t1 := schedule.Next(earliestTime)
	t2 := schedule.Next(t1)

	// It is possible for cron.ParseStandard("59 23 31 2 *") to return an invalid schedule
	// minute - 59, hour - 23, dom - 31, month - 2, and dow is optional, clearly 31 is invalid
	// In this case the timeBetweenTwoSchedules will be 0, and we error out the invalid schedule
	scheduleInterval := t2.Sub(t1)
	scheduleIntervalInSeconds := int64(t2.Sub(t1).Round(time.Second).Seconds())

	if now.Before(t1) {
		return scheduleInterval, &earliestTime
	}
	if now.Before(t2) {
		return scheduleInterval, &t1
	}

	// this logic used for calculating number of missed schedules does a rough
	// approximation, by calculating a diff between two schedules (t1 and t2),
	// and counting how many of these will fit in between last schedule and now.
	timeElapsed := int64(now.Sub(t1).Seconds())
	numberOfMissedSchedules := (timeElapsed / scheduleIntervalInSeconds) + 1

	var mostRecentTime time.Time
	// to get the most recent time accurate for regular schedules and the ones
	// specified with @every form, we first need to calculate the potential earliest
	// time by multiplying the initial number of missed schedules by its interval,
	// this is critical to ensure @every starts at the correct time, this explains
	// the numberOfMissedSchedules-1, the additional -1 serves there to go back
	// in time one more time unit, and let the cron library calculate a proper
	// schedule, for case where the schedule is not consistent, for example
	// something like  30 6-16/4 * * 1-5
	potentialEarliest := t1.Add(time.Duration((numberOfMissedSchedules-1-1)*scheduleIntervalInSeconds) * time.Second)
	for t := schedule.Next(potentialEarliest); !t.After(now); t = schedule.Next(t) {
		mostRecentTime = t
	}

	if mostRecentTime.IsZero() {
		if numberOfMissedSchedules == 0 {
			// no missed schedules since earliestTime
			mostRecentTime = earliestTime
		} else {
			// if there are missed schedules since earliestTime, always use now
			mostRecentTime = now
		}
	}
	return scheduleInterval, &mostRecentTime
}
