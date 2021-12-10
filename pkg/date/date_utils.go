// Package date is a collection of date utilities
package date

import "time"

const (
	apiDateLayout = "2006-01-02T15:01:05Z"
	apiDBLayout   = "2006-01-02 15:01:05"
)

// GetNow returns the current time
func GetNow() time.Time {
	return time.Now()
}

// GetNowString returns the current time
func GetNowString() string {
	return GetNow().Format(apiDateLayout)
}

// GetNowDBFormat returns the current time with db format
func GetNowDBFormat() string {
	return GetNow().Format(apiDBLayout)
}
