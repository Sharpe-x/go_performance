package utils

import (
	"sync"
	"time"
)

type WaitGroupWrapper struct {
	sync.WaitGroup
}

func (w *WaitGroupWrapper) Wrap(cb func()) {
	w.Add(1)
	go func() {
		defer w.Done()
		cb()
	}()
}

// Truncate returns the result of rounding x toward zero to a multiple of m.
// If m <= 0, Truncate returns x unchanged.
func Truncate(x, m int64) int64 {
	if m <= 0 {
		return x
	}
	return x - x%m
}

// TimeToMs returns an integer number, which represents t in milliseconds.
func TimeToMs(t time.Time) int64 {
	return t.UnixNano() / int64(time.Millisecond)
}

// MsToTime returns the UTC time corresponding to the given Unix time,
// t milliseconds since January 1, 1970 UTC.
func MsToTime(t int64) time.Time {
	return time.Unix(0, t*int64(time.Millisecond)).UTC()
}
