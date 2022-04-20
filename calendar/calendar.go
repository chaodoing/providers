package calendar

import (
	"time"
)


// Calendar 日历月数据
type Calendar []time.Time

// Calendars 获取日历
func Calendars(Year int, Month int64) (date Calendar) {
	var (
		begin time.Time = time.Date(Year, time.Month(Month), 1, 0, 0, 0, 0, time.Local)
		month time.Time = time.Date(Year, time.Month(Month+1), 1, 0, 0, 0, 0, time.Local)
	)
	date = append(date, begin)
	month = month.Add(-24 * time.Hour) // 减少一天
	var index int = 1
	for index = 1; index < month.Day(); index++ {
		date = append(date, begin.Add(time.Duration(index)*24*time.Hour))
	}
	return
}
