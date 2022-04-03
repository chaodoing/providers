package calendar

func (y Year) MarshalText() ([]byte, error) {
	return marshalText(y)
}
func (y *Year) UnmarshalText(data []byte) error {
	t, err := unmarshalText(y, data)
	if err != nil {
		return err
	}
	*y = Year(t)
	return nil
}
// Value 为 Time 实现 Value 方法，写入数据库时会调用该方法将自定义时间类型转换并写入数据库
func (t Year) Value() (interface{}, error) {
	return value(t)
}