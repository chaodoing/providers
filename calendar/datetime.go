package calendar

func (d Datetime) MarshalText() ([]byte, error) {
	return marshalText(d)
}

func (d *Datetime) UnmarshalText(data []byte) error {
	t, err := unmarshalText(d, data)
	if err != nil {
		return err
	}
	*d = Datetime(t)
	return nil
}
// Value 为 Time 实现 Value 方法，写入数据库时会调用该方法将自定义时间类型转换并写入数据库
func (d Datetime) Value() (interface{}, error) {
	return value(d)
}