package utils

import (
	"fmt"
	"time"
)

// GetCurrentUTCTimestamp return time.Time
func GetCurrentUTCTimestamp() time.Time {
	loc, _ := time.LoadLocation("UTC")
	return time.Now().In(loc)
}

// FormatTimestamp return string
func FormatTimestamp(timestamp time.Time) string {
	t := timestamp
	return fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
		t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
}

// ConvertDateStringToUnixEpoch returns int64
func ConvertDateStringToUnixEpoch(dateString string) int64 {
	t, err := time.Parse("2006-01-02T15:04:05", dateString)
	if err != nil {
		return 0
	}
	return t.Unix()
}

// GetCurrentDateTimeOffset returns string
func GetCurrentDateTimeOffset(RequestTTLInMinutes int) string {
	CurrentDateTime := GetCurrentUTCTimestamp()
	CurrentDateTimeOffset := CurrentDateTime.Add(time.Duration(-RequestTTLInMinutes) * time.Minute)
	CurrentDateTimeOffsetFormatted := FormatTimestamp(CurrentDateTimeOffset)
	return CurrentDateTimeOffsetFormatted
}
