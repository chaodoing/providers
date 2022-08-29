package request

import (
	`github.com/kataras/iris/v12`
)

type PageQuery struct {
	Page      uint  `json:"page" xml:"page"`
	Limit     uint  `json:"limit" xml:"limit"`
	Offset    int64 `json:"offset" xml:"offset"`
	Condition struct {
		Where     string
		Arguments []interface{}
	}
}

func Query(ctx iris.Context, limit int) (q PageQuery) {
	q.Page = uint(ctx.URLParamIntDefault("page", 0))
	q.Limit = uint(ctx.URLParamIntDefault("limit", limit))
	if q.Page > 0 {
		q.Offset = int64((q.Page - 1) * q.Limit)
	} else {
		q.Page = 1
	}
	return q
}

// Where 设置查询条件
func (q PageQuery) Where(where string, query ...interface{}) PageQuery {
	q.Condition = struct {
		Where     string
		Arguments []interface{}
	}{Where: where, Arguments: query}
	return q
}
