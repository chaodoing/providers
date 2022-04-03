package calendar

func (m Month) MarshalText() ([]byte, error) {
	return marshalText(m)
}

func (m *Month) UnmarshalText(data []byte) error {
	t, err := unmarshalText(m, data)
	if err != nil {
		return err
	}
	*m = Month(t)
	return nil
}
// Value 为 Time 实现 Value 方法，写入数据库时会调用该方法将自定义时间类型转换并写入数据库
func (m Month) Value() (interface{}, error) {
	return value(m)
}
