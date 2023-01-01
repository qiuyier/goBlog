package response

import (
	"github.com/mlogclub/simple/sqls"
	"github.com/mlogclub/simple/web"
	"qiuyier/blog/pkg/constants"
)

type Result struct {
	Code    int         `json:"code"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data"`
	Success bool        `json:"success"`
}

func Json(code int, msg string, data interface{}, success bool) *Result {
	return &Result{
		Code:    code,
		Msg:     msg,
		Data:    data,
		Success: success,
	}
}

func SuccessDataResult(data interface{}) *Result {
	return &Result{
		Code:    constants.DefaultSuccessCode,
		Msg:     "请求成功",
		Data:    data,
		Success: true,
	}
}

func SuccessReasonResult(msg string) *Result {
	return &Result{
		Code:    constants.DefaultSuccessCode,
		Msg:     msg,
		Data:    nil,
		Success: true,
	}
}

func SuccessNoReasonResult() *Result {
	return &Result{
		Code:    constants.DefaultSuccessCode,
		Msg:     "请求成功",
		Data:    nil,
		Success: true,
	}
}

func JsonError(err error) *Result {
	if e, ok := err.(*CodeError); ok {
		return &Result{
			Code:    e.Code,
			Msg:     e.Message,
			Data:    e.Data,
			Success: false,
		}
	}
	return &Result{
		Code:    0,
		Msg:     err.Error(),
		Data:    nil,
		Success: false,
	}
}

func FailMsgResult(msg string) *Result {
	return &Result{
		Code:    0,
		Msg:     msg,
		Data:    nil,
		Success: false,
	}
}

func FailNoMsgResult() *Result {
	return &Result{
		Code:    0,
		Msg:     "请求失败",
		Data:    nil,
		Success: false,
	}
}

func NewErrorResult(code int, msg string) *Result {
	return &Result{
		Code:    code,
		Msg:     msg,
		Data:    nil,
		Success: false,
	}
}

func ErrorCodeResult(code int) *Result {
	return &Result{
		Code:    code,
		Msg:     "请求失败",
		Data:    nil,
		Success: false,
	}
}

func PageDataResult(results interface{}, page *sqls.Paging) *Result {
	return &Result{
		Code: constants.DefaultSuccessCode,
		Msg:  "请求成功",
		Data: &web.PageResult{
			Results: results,
			Page:    page,
		},
		Success: true,
	}
}

func CursorDataResult(results interface{}, cursor string, hasMore bool) *Result {
	return &Result{
		Code: constants.DefaultSuccessCode,
		Msg:  "请求成功",
		Data: &web.CursorResult{
			Results: results,
			Cursor:  cursor,
			HasMore: hasMore,
		},
		Success: true,
	}
}

type RspBuilder struct {
	Data map[string]interface{}
}

func NewEmptyRspBuilder() *RspBuilder {
	return &RspBuilder{Data: make(map[string]interface{})}
}

func (builder *RspBuilder) Put(key string, value interface{}) *RspBuilder {
	builder.Data[key] = value
	return builder
}

func (builder *RspBuilder) GetStruct() map[string]interface{} {
	return builder.Data
}

func (builder *RspBuilder) BuildResult() *Result {
	return SuccessDataResult(builder.Data)
}
