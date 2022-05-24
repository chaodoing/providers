package util

import (
	`encoding/xml`
	
	`github.com/kataras/iris/v12`
	`gorm.io/gorm`
)

// PageQuery 分页数据请求
type PageQuery struct {
	ctx          iris.Context
	Page         uint   // 当前分页
	Limit        uint16 // 每页数据条数
	Offset       int    // 总条数
	hasCondition bool
	Condition    struct {
		Where string
		Param []interface{}
	} // Conditions 数据查询条件
}

// Where 设置查询条件
func (q PageQuery) Where(query string, args ...interface{}) {
	q.Condition.Where = query
	q.Condition.Param = args
	q.hasCondition = true
}

// Page 解析分页查询条件
func Page(ctx iris.Context) (p PageQuery) {
	p.ctx = ctx
	p.Page = uint(ctx.URLParamIntDefault("page", 0))
	p.Limit = uint16(ctx.URLParamIntDefault("limit", 15))
	if p.Page > 0 {
		p.Offset = (int(p.Page) - 1) * int(p.Limit)
	} else {
		p.Page = 1
	}
	return
}

// Invoke 分页查询回调
type Invoke func(db *gorm.DB, query PageQuery) (error, interface{})

// Pagination 分页数据返回
type Pagination struct {
	ctx     iris.Context `json:"-" xml:"-"`
	XMLName xml.Name     `xml:"root" json:"-"`
	Page    uint         `json:"page" xml:"page"`             // Page 当前页码
	Limit   uint16       `json:"limit" xml:"limit"`           // Limit 每页条数
	Total   int64        `json:"total" xml:"total"`           // Total 总计条数
	Status  uint8        `json:"error_code" xml:"error_code"` // Status 状态错误代码
	Message string       `json:"message" xml:"message"`       // Message 错误消息
	Data    interface{}  `json:"data" xml:"data"`             // Data 数据内容
}

// Find 查找分页数据
func (query PageQuery) Find(db *gorm.DB, invoke Invoke) (pagination Pagination) {
	pagination.ctx = query.ctx
	pagination.Page = query.Page
	pagination.Limit = query.Limit
	var err error
	if query.hasCondition {
		err = db.Where(query.Condition.Where, query.Condition.Param...).Count(&pagination.Total).Error
	} else {
		err = db.Count(&pagination.Total).Error
	}
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

// Send 发送分页数据到web
func (p Pagination) Send() {
	NewRespond(p.ctx).SendData(p)
}
