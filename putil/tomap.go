package putil

import (
	`reflect`
	`strings`
)

func IsExistInArrayString(value string, column []string) bool {
	for _, v := range column {
		if v == value {
			return true
		}
	}
	return false
}

func ToMap(data interface{}, columns []string) map[string]interface{} {
	t := reflect.TypeOf(data)
	v := reflect.ValueOf(data)
	var Columns = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		tagValues := strings.Split(t.Field(i).Tag.Get("gorm"), ";")
		value := v.Field(i).Interface()
		for _, v := range tagValues {
			if strings.HasPrefix(v, "column") {
				columnTag := strings.Split(v, ":")
				column := columnTag[1]
				if IsExistInArrayString(t.Field(i).Name, columns) {
					Columns[column] = value
				}
			}
		}
	}
	return Columns
}
