package utils

import "time"

func NormalizeDay(year int, month time.Month, day int) int {
	lastDayOfMonth := time.Date(year, month+1, 0, 0, 0, 0, 0, time.UTC).Day()

	if day > lastDayOfMonth {
		return lastDayOfMonth
	}

	if day < 1 {
		return 1
	}

	return day
}

func CreateDateWithNormalizedDay(year int, month time.Month, day int) time.Time {
	normalizedDay := NormalizeDay(year, month, day)
	return time.Date(year, month, normalizedDay, 0, 0, 0, 0, time.UTC)
}
