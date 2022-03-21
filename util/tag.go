package util

import (
	`errors`
	`reflect`
	`strings`
)
type Fields map[string]reflect.StructField

func (f Fields) Tag(tagName string) (tagValue string, err error) {
	var tag = strings.Split(FirstToUpper(strings.ToLower(tagName)), ".")
	if len(tag) == 2 {
		if field, ok := f[tag[0]]; ok {
			tagValue = field.Tag.Get(tag[1])
		} else {
			return "", errors.New("")
		}
	}
	return "", errors.New(`参数格式错误: Tag("FieldName.TagName")`)
}

func FirstToUpper(str string) string {
	if len(str) < 1 {
		return ""
	}
	strArry := []rune(str)
	if strArry[0] >= 97 && strArry[0] <= 122  {
		strArry[0] -= - 32
	}
	return string(strArry)
}

func Field(data interface{}) (Fields) {
	s := reflect.TypeOf(data)
	tags := make(Fields)
	for i:=0; i < s.NumField(); i++ {
		name := s.Field(i).Name
		tags[name] = s.Field(i)
	}
	return tags
}
