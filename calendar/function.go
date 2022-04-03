package calendar

import (
	`errors`
	`time`
)

// marshalText 为 Time 重写 MarshaJSON 方法，在此方法中实现自定义格式的转换
func marshalText(t interface{}) ([]byte, error) {
	switch value := t.(type) {
	case Year:
		text := time.Time(value).Format(FORMAT_YEAR)
		return []byte(text), nil
	case Month:
		text := time.Time(value).Format(FORMAT_MONTH)
		return []byte(text), nil
	case Date:
		text := time.Time(value).Format(FORMAT_DATE)
		return []byte(text), nil
	case Time:
		text := time.Time(value).Format(FORMAT_TIME)
		return []byte(text), nil
	case Datetime:
		text := time.Time(value).Format(FORMAT_DATE_TIME)
		return []byte(text), nil
	default:
		v, ok := value.(Datetime)
		if ok {
			text := time.Time(v).Format(FORMAT_DATE_TIME)
			return []byte(text), nil
		}
	}
	return nil, errors.New("格式不支持")
}

func unmarshalText(ts interface{}, data []byte) (time.Time, error) {
	switch v := ts.(type) {
	case *Year:
		return time.Parse(FORMAT_YEAR, string(data))
	case *Month:
		return time.Parse(FORMAT_MONTH, string(data))
	case *Date:
		return time.Parse(FORMAT_DATE, string(data))
	case *Time:
		return time.Parse(FORMAT_TIME, string(data))
	case *Datetime:
		return time.Parse(FORMAT_DATE_TIME, string(data))
	default:
		if _, ok := v.(time.Time); ok {
			return time.Parse(FORMAT_DATE_TIME, string(data))
		} else {
			return time.Now(), errors.New("格式不支持:unmarshalText")
		}
	}
}

// value 为 Time 实现 Value 方法，写入数据库时会调用该方法将自定义时间类型转换并写入数据库
func value(t interface{}) (interface{}, error) {
	var zeroTime time.Time // 初始化时间 1971-01-01
	switch val := t.(type) {
	case Year:
		if time.Time(val).UnixNano() == zeroTime.UnixNano() { // 如果时间是初试时间 则放回空值
			return nil, nil
		}
		return time.Time(val).Format(FORMAT_YEAR), nil
	case Month:
		if time.Time(val).UnixNano() == zeroTime.UnixNano() { // 如果时间是初试时间 则放回空值
			return nil, nil
		}
		return time.Time(val).Format(FORMAT_MONTH), nil
	case Date:
		if time.Time(val).UnixNano() == zeroTime.UnixNano() { // 如果时间是初试时间 则放回空值
			return nil, nil
		}
		return time.Time(val).Format(FORMAT_DATE), nil
	case Time:
		if time.Time(val).UnixNano() == zeroTime.UnixNano() { // 如果时间是初试时间 则放回空值
			return nil, nil
		}
		return time.Time(val).Format(FORMAT_TIME), nil
	case Datetime:
		if time.Time(val).UnixNano() == zeroTime.UnixNano() { // 如果时间是初试时间 则放回空值
			return nil, nil
		}
		return time.Time(val).Format(FORMAT_DATE_TIME), nil
	default:
		v, ok := val.(Datetime)
		if ok {
			if time.Time(v).UnixNano() == zeroTime.UnixNano() { // 如果时间是初试时间 则放回空值
				return nil, nil
			}
			return time.Time(v).Format(FORMAT_DATE_TIME), nil
		}
	}
	return nil, errors.New("格式不支持")
}
