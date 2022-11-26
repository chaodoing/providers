package o

import (
	`encoding/json`
	`os`
)

// ReadJSON 读取JSON
func ReadJSON(file string, data interface{}) (err error) {
	var content []byte
	content, err = os.ReadFile(os.ExpandEnv(file))
	if err != nil {
		return err
	}
	err = json.Unmarshal(content, &data)
	if err != nil {
		return err
	}
	return nil
}

// SaveJSON 写入JSON
func SaveJSON(data interface{}, file string) error {
	xmlByte, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}
	if err := os.WriteFile(os.ExpandEnv(file), xmlByte, 0666); err != nil {
		return err
	}
	return nil
}
