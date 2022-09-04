package timeutil

import "time"

// NowMillis represents current time in milliseconds
func NowMillis() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

// NowSeconds represents current time in seconds
func NowSeconds() int64 {
	return time.Now().UnixNano() / int64(time.Second)
}
