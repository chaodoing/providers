package calendar

import (
	`time`
)

// NewCalendars 创建日历实例
func NewCalendars() *Calendar {
	return &Calendar{
		calendars: nil,
		year:      time.Now().Year(),
		month:     int(time.Now().Month()),
	}
}

// Year 设置年
func (c *Calendar) Year(year int) *Calendar {
	c.year = year
	return c
}

// Month 设置月
func (c *Calendar) Month(month int) *Calendar {
	c.month = month
	return c
}

// Get 获取日历
func (c *Calendar) Get() []Datetime {
	var (
		begin = time.Date(c.year, time.Month(c.month), 1, 0, 0, 0, 0, time.Local)
		month = time.Date(c.year, time.Month(c.month+1), 1, 0, 0, 0, 0, time.Local)
	)
	c.calendars = append(c.calendars, Datetime(begin))
	month = month.Add(-24 * time.Hour) // 减少一天
	var index int = 1
	for index = 1; index < month.Day(); index++ {
		c.calendars = append(c.calendars, Datetime(begin.Add(time.Duration(index)*24*time.Hour)))
	}
	return c.calendars
}

// Calendars 生成日历
func (c *Calendar) Calendars(Year int, Month int) []Datetime {
	c.year = Year
	c.month = Month
	return c.Get()
}
