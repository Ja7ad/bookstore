// Package date is a collection of date utilities
package date

import "time"

const (
	apiDateLayout = "2006-01-02T15:01:05Z"
)

// GetNow returns the current time
func GetNow() time.Time {
	return time.Now()
}

// GetNowString returns the current time
func GetNowString() string {
	return GetNow().Format(apiDateLayout)
}
