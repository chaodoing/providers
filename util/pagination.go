package util

import (
	`encoding/xml`
	
	`github.com/kataras/iris/v12`
	`gorm.io/gorm`
)

type PageQuery struct {
	Page       uint   `json:"page" xml:"page"`
	Limit      uint16 `json:"limit" xml:"limit"`
	Offset     int    `json:"offset" xml:"offset"`
	parameters map[string]string
}

// Pagination 分页数据返回
type Pagination struct {
	XMLName xml.Name    `xml:"root" json:"-"`
	Page    uint        `json:"page" xml:"page"`             // Page 当前页码
	Limit   uint16      `json:"limit" xml:"limit"`           // Limit 每页条数
	Total   int64       `json:"total" xml:"total"`           // Total 总计条数
	Status  uint8       `json:"error_code" xml:"error_code"` // Status 状态错误代码
	Message string      `json:"message" xml:"message"`       // Message 错误消息
	Data    interface{} `json:"data" xml:"data"`             // Data 数据内容
}

// Invoke 分页查询回调
type Invoke func(tx *gorm.DB, query *PageQuery) (error, interface{})

// Query 查询条件
func Query(ctx iris.Context) (query *PageQuery) {
	query.Page = uint(ctx.URLParamIntDefault("page", 0))
	query.Limit = uint16(ctx.URLParamIntDefault("limit", 15))
	if query.Page > 0 {
		query.Offset = (int(query.Page) - 1) * int(query.Limit)
	} else {
		query.Page = 1
	}
	query.parameters = ctx.URLParams()
	delete(query.parameters, "limit")
	delete(query.parameters, "page")
	return
}

// Params 获取请求查询参数
func (query *PageQuery) Params() map[string]string {
	return query.parameters
}

// Pagination 查找分页数据
func (query *PageQuery) Pagination(db *gorm.DB, invoke Invoke) (pagination Pagination) {
	pagination.Page = query.Page
	pagination.Limit = query.Limit
	var err error
	err = db.Count(&pagination.Total).Error
	if err != nil {
		pagination.Status = 1
		pagination.Message = err.Error()
		return
	}
	err, pagination.Data = invoke(db, query)
	if err != nil {
		pagination.Status = 1
		pagination.Message = err.Error()
		return
	}
	return
}
