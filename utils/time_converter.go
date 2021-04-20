package utils

import (
	"time"
)

func ConvertTimestampToTime(timestamp int64) time.Time {

	timeInUTC := time.Unix(timestamp, 0)

	return timeInUTC
}
