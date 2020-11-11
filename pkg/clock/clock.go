package clock

import (
	"time"
)

// GetWallClock returns the current wall clock time. This is fetched
// from the CLOCK_REALTIME clock on Linux and GetSystemTimeAsFileTime API
// on Windows.
func GetWallClock() time.Time {
	return time.Now().UTC()
}
