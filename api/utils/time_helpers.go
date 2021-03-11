package utils

import (
	"time"
)

// ConvertDateStringToUnixEpoch returns int64
func ConvertDateStringToUnixEpoch(dateString string) int64 {
	t, err := time.Parse("2006-01-02T15:04:05", dateString)
	if err != nil {
		return 0
	}
	return t.Unix()
}

// GetCurrentDateTimeOffset returns time.Time
func GetCurrentDateTimeOffset(RequestTTLInMinutes int) time.Time {
	CurrentDateTime := time.Now()
	CurrentDateTimeOffset := CurrentDateTime.Add(time.Duration(-RequestTTLInMinutes) * time.Minute)
	return CurrentDateTimeOffset
}
