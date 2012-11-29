package gallifrey

import (
	"fmt"
	"time"
)

func RelativeTime(t time.Time) string {
	duration := time.Since(t)
	var (
		count      int
		unit       string
		monthHours = 24 * 31
		weekHours  = 24 * 7
		dayHours   = 24
	)

	if hours := int(duration.Hours()); hours >= monthHours {
		count = hours / monthHours
		unit = "months"
	} else if hours := int(duration.Hours()); hours >= weekHours {
		count = hours / weekHours
		unit = "weeks"
	} else if hours := int(duration.Hours()); hours >= dayHours {
		count = hours / dayHours
		unit = "days"
	} else if hours := int(duration.Hours()); hours >= 1 {
		count = hours
		unit = "hours"
	} else if mins := int(duration.Minutes()); mins >= 1 {
		count = mins
		unit = "minutes"
	} else {
		count = int(duration.Seconds())
		unit = "seconds"
	}

	return fmt.Sprintf("%d %s ago", count, unit)
}
