package calendar

func (d Time) MarshalText() ([]byte, error) {
	return marshalText(d)
}

func (d *Time) UnmarshalText(data []byte) error {
	t, err := unmarshalText(d, data)
	if err != nil {
		return err
	}
	*d = Time(t)
	return nil
}
// Value 为 Time 实现 Value 方法，写入数据库时会调用该方法将自定义时间类型转换并写入数据库
func (d Time) Value() (interface{}, error) {
	return value(d)
}