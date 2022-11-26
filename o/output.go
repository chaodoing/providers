package o

import (
	`github.com/chaodoing/providers/response`
	`github.com/kataras/iris/v12`
)

func O(ctx iris.Context, status uint32, message string, data interface{}) {
	r := response.Response{}
	r.O(ctx, status, message, data)
}

func Set(ctx iris.Context, data interface{}) {
	r := response.Response{}
	r.Set(ctx, data)
}

func Pagination(ctx iris.Context, status uint32, message string, data interface{}, total uint64, page uint, limit uint) {
	r := response.Response{}
	r.Pagination(ctx, status, message, data, total, page, limit)
}
