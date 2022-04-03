package test

import (
	`encoding/json`
	`testing`
	
	`github.com/chaodoing/providers/calendar`
)

type Timespec struct {
	Year     calendar.Year     `json:"year" xml:"year"`
	Month    calendar.Month    `json:"month" xml:"month"`
	Date     calendar.Date     `json:"date" xml:"date"`
	Datetime calendar.Datetime `json:"datetime" xml:"datetime"`
	Time     calendar.Time     `json:"time" xml:"time"`
}

func TestTime(t *testing.T) {
	
	var dt = `{
                "year": "2022",
                "month": "2022-04",
                "date": "2022-04-03",
                "datetime": "2022-04-03 19:40:49",
                "time": "19:40:49"
        }`
	var data Timespec
	if err := json.Unmarshal([]byte(dt), &data); err != nil {
		t.Error(err)
		return
	}
	bytes, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(string(bytes))
}
