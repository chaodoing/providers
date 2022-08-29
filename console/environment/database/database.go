package database

import (
	`encoding/xml`
)

type Database struct {
	XMLName  xml.Name  `xml:"database"`
	Comment  string    `xml:"comment,attr"`
	Drive    *Drive    `xml:"drive"`
	Host     *Host     `xml:"host"`
	Port     *Port     `xml:"port"`
	Db       *Db       `xml:"db"`
	Username *Username `xml:"username"`
	Password *Password `xml:"password"`
	Charset  *Charset  `xml:"charset"`
}

// NewDatabase 数据库配置
func NewDatabase() (data *Database, err error) {
	data = &Database{
		Comment: "数据库配置",
	}
	var (
		drive    *Drive
		host     *Host
		port     *Port
		db       *Db
		username *Username
		password *Password
		charset  *Charset
	)
	drive, err = NewDrive()
	if err != nil {
		return
	}
	host, err = NewHost()
	if err != nil {
		return
	}
	port, err = NewPort()
	if err != nil {
		return
	}
	db, err = NewDb()
	if err != nil {
		return
	}
	username, err = NewUsername()
	if err != nil {
		return
	}
	password, err = NewPassword()
	if err != nil {
		return
	}
	charset, err = NewCharset()
	if err != nil {
		return
	}
	data.Drive = drive
	data.Host = host
	data.Port = port
	data.Db = db
	data.Username = username
	data.Password = password
	data.Charset = charset
	return
}
