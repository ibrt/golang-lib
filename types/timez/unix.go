package timez

import "time"

// Unix converts the time to a Unix timestamp in the given resolution (e.g. time.Millisecond).
func Unix(t time.Time, resolution time.Duration) int64 {
	return t.UnixNano() / int64(resolution)
}

// FromUnix converts a Unix timestamp in the given resolution to Time.
func FromUnix(sec, frac int64, resolution time.Duration) time.Time {
	return time.Unix(sec, frac*int64(resolution))
}

// UnixMilli converts the time to a Unix timestamp in milliseconds.
func UnixMilli(t time.Time) int64 {
	return Unix(t, time.Millisecond)
}

// FromUnixMilli converts a Unix timestamp in milliseconds to Time.
func FromUnixMilli(sec, millisec int64) time.Time {
	return FromUnix(sec, millisec, time.Millisecond)
}

// UnixMicro converts the time to a Unix timestamp in microseconds.
func UnixMicro(t time.Time) int64 {
	return t.UnixNano() / int64(time.Microsecond)
}

// FromUnixMicro converts a Unix timestamp in microseconds to Time.
func FromUnixMicro(sec, microsec int64) time.Time {
	return FromUnix(sec, microsec, time.Microsecond)
}
