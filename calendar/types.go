package calendar

import (
	`time`
)

type (
	Unix     int64
	Year     time.Time
	Month    time.Time
	Date     time.Time
	Datetime time.Time
	Time     time.Time
	Calendar struct {
		// calendars 当前日历
		calendars []Datetime
		// Year 当前年
		year int
		// Month 当前月
		month int
	}
)