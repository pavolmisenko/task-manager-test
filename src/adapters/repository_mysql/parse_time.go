package repository_mysql

import "time"

func StringToTime(s string) (time.Time, error) {
	layout := "2006-01-02 15:04:05"
	return time.Parse(layout, s)
}
