package dates

import "time"

const (
	defaultDateLayout = "2006-01-02T15:04:05Z"
	dbDateLayout      = "2006-01-02 15:04:05"
)

func GetNowUtcDefault() string {
	return GetNowUTC().Format(defaultDateLayout)
}

func GetNowUtcDbFormat() string {
	return GetNowUTC().Format(dbDateLayout)
}

func GetNowUTC() time.Time {
	return time.Now().UTC()
}
