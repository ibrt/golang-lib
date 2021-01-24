package timez

import "time"

// UnixMilli converts the time to a Unix timestamp in milliseconds.
func UnixMilli(t time.Time) int64 {
	return t.UnixNano() / int64(time.Millisecond)
}

// FromUnixMilli converts a Unix timestamp in milliseconds to Time.
func FromUnixMilli(sec, msec int64) time.Time {
	return time.Unix(sec, msec*int64(time.Millisecond))
}
