package utils

import (
	"fmt"
	"time"
)

func GetTime(millis int64) time.Time {
	return time.Unix(millis, 0)
}

func GetTimeStampByString(str string) int64 {
	t, err := time.Parse(time.RFC3339, str)

	if err != nil {
		fmt.Println(err)
	}
	return t.Unix()
}

func GetTimeByString(str string) time.Time {
	t, err := time.Parse(time.RFC3339, str)

	if err != nil {
		fmt.Println(err)
	}
	return t
}

func GetCurrentTimeStamp() int64 {
	return time.Now().Unix()
}
