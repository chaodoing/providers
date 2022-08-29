package response

import (
	`encoding/xml`
)

type Data struct {
	XMLName xml.Name    `xml:"root" json:"-"`
	Status  uint32      `json:"status" xml:"status"`
	Message string      `json:"message" xml:"message"`
	Data    interface{} `json:"data" xml:"data"`
}

type Pagination struct {
	XMLName xml.Name    `xml:"root" json:"-"`
	Status  uint32      `json:"status" xml:"status"`
	Message string      `json:"message" xml:"message"`
	Total   uint64      `json:"total" xml:"total"`
	Page    uint        `json:"page" xml:"page"`
	Limit   uint        `json:"limit" xml:"limit"`
	Data    interface{} `json:"data" xml:"data"`
}
