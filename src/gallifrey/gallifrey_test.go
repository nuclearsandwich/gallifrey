package gallifrey

import(
	"testing"
	"time"
)

func TestRelativeTime(t *testing.T) {
	aTime := time.Date(2012, 10, 1, 0, 0, 0, 0, time.UTC)
	t.Log(aTime)
	if rel := RelativeTime(aTime); rel != "1 months ago" {
		t.Error(rel)
	}
}
