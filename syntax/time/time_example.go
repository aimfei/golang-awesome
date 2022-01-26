package time_example

import "time"

func TimeAdd(now time.Time, duration time.Duration) time.Time {
	add := now.Add(duration)
	return add
}
