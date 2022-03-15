package util

import (
	`encoding/xml`
	`io/ioutil`
	`os`
)

// ReadXML 读取XML文件
func ReadXML(file string, data interface{}) (err error) {
	var content []byte
	content, err = os.ReadFile(os.ExpandEnv(file))
	if err != nil {
		return err
	}
	err = xml.Unmarshal(content, &data)
	if err != nil {
		return err
	}
	return nil
}

// SaveXML 存储XML文件
func SaveXML(data interface{}, file string) error {
	xmlByte, err := xml.MarshalIndent(data, "","\t")
	if err != nil {
		return err
	}
	if err := ioutil.WriteFile(os.ExpandEnv(file), xmlByte, 0666); err != nil {
		return err
	}
	return nil
}
