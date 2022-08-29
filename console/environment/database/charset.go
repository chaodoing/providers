package database

import (
	`encoding/xml`
	
	`github.com/manifoldco/promptui`
)

type Charset struct {
	XMLName  xml.Name `xml:"charset"`
	XMLValue string   `xml:",innerxml"`
	Comment  string   `xml:"comment,attr"`
}

// NewCharset 生成数据库连接字符集
func NewCharset() (data *Charset, err error) {
	data = &Charset{
		Comment: "数据库连接字符集",
	}
	command := promptui.Select{
		Label: "数据库连接字符集",
		Items: []string{"armscii8", "ascii", "big5", "binary", "cp1250", "cp1251", "cp1256", "cp1257", "cp850", "cp852", "cp866", "cp932", "dec8", "eucjpms", "euckr", "gb18030", "gb2312", "gbk", "geostd8", "greek", "hebrew", "hp8", "keybcs2", "koi8r", "koi8u", "latin1", "latin2", "latin5", "latin7", "macce", "macroman", "sjis", "swe7", "tis620", "ucs2", "ujis", "utf16", "utf16le", "utf32", "utf8", "utf8mb4"},
	}
	_, value, err := command.Run()
	if err != nil {
		return nil, err
	}
	data.XMLValue = value
	return
}
