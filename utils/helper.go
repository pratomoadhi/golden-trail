package utils

import "time"

func Today() time.Time {
	return time.Now().Truncate(24 * time.Hour)
}

func ParseDate(dateStr string) time.Time {
	t, _ := time.Parse("2006-01-02", dateStr)
	return t
}
