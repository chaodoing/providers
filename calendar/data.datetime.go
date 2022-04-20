package calendar

import (
	"database/sql/driver"
	"time"
)

// Datetime 年-月-日 时:分:秒
type Datetime time.Time

// MarshalText 为 Time 重写 MarshaJSON 方法，在此方法中实现自定义格式的转换；
func (t Datetime) MarshalText() ([]byte, error) {
	text := time.Time(t).Format(FORMAT_DATE_TIME)
	return []byte(text), nil
}

func (t *Datetime) UnmarshalText(data []byte) error {
	ts, err := time.Parse(FORMAT_DATE_TIME, string(data))
	if err == nil {
		*t = Datetime(ts)
	}
	return err
}

// Value 为 Time 实现 Value 方法，写入数据库时会调用该方法将自定义时间类型转换并写入数据库
func (t Datetime) Value() (driver.Value, error) {
	var zeroTime time.Time                              // 初始化时间 1971-01-01
	if time.Time(t).UnixNano() == zeroTime.UnixNano() { // 如果时间是初试时间 则放回空值
		return nil, nil
	}
	return time.Time(t).Format(FORMAT_DATE_TIME), nil
}
