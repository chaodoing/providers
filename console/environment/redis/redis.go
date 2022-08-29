package redis

import (
	`encoding/xml`
)

type Redis struct {
	XMLName  xml.Name  `xml:"redis"`
	Comment  string    `xml:"comment,attr"`
	Host     *Host     `xml:"host"`
	Port     *Port     `xml:"port"`
	Db       *Db       `xml:"db"`
	Password *Password `xml:"password"`
	TTL      *TTL      `xml:"ttl"`
}

func NewRedis() (data *Redis, err error) {
	var (
		host     *Host
		port     *Port
		db       *Db
		password *Password
		ttl      *TTL
	)
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
	password, err = NewPassword()
	if err != nil {
		return
	}
	ttl, err = NewTTL()
	if err != nil {
		return
	}
	data = &Redis{
		Comment:  "Redis配置",
		Host:     host,
		Port:     port,
		Db:       db,
		Password: password,
		TTL:      ttl,
	}
	return
}
