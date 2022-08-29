package encrypt

import (
	`github.com/google/uuid`
)

// UUID 生成UUID
//  @return string uuid 字符串
func UUID() string {
	return uuid.New().String()
}
