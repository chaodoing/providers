package log

import (
	`encoding/xml`
	
	`github.com/chaodoing/providers/console/environment/log/database`
	`github.com/chaodoing/providers/console/environment/log/iris`
)

type Log struct {
	XMLName   xml.Name           `xml:"log"`
	Comment   string             `xml:"comment,attr"`
	Console   *Console           `xml:"console"`
	Record    *Record            `xml:"record"`
	Directory *Directory         `xml:"directory"`
	Iris      *iris.Iris         `xml:"iris"`
	Database  *database.Database `xml:"database"`
}

func NewLog() (data *Log, err error) {
	var (
		console     *Console
		record      *Record
		directory   *Directory
		irisLog     *iris.Iris
		databaseLog *database.Database
	)
	console, err = NewConsole()
	if err != nil {
		return
	}
	record, err = NewRecord()
	if err != nil {
		return
	}
	directory, err = NewDirectory()
	if err != nil {
		return
	}
	irisLog, err = iris.NewIris()
	if err != nil {
		return
	}
	databaseLog, err = database.NewDatabase()
	if err != nil {
		return
	}
	data = &Log{
		Comment:   "日志配置",
		Console:   console,
		Record:    record,
		Directory: directory,
		Iris:      irisLog,
		Database:  databaseLog,
	}
	return
}
