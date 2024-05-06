package model

import (
	"encoding/json"
	"strconv"
	"sync"

	"github.com/bububa/easyliao/util"
)

// Response api response interface
type Response interface {
	// IsError 是否返回错误
	IsError() bool
	// Error implement error interface
	Error() string
}

// BaseResponse API 响应
type BaseResponse struct {
	// Code 响应码
	Code int `json:"code,omitempty"`
	// Msg 响应信息
	Msg string `json:"msg,omitempty"`
}

func (r BaseResponse) IsError() bool {
	return r.Code != 0
}

// Error implments error
func (r BaseResponse) Error() string {
	return util.StringsJoin("code:", strconv.Itoa(r.Code), ", msg:", r.Msg)
}

var baseResponsePool = sync.Pool{
	New: func() any {
		return new(BaseResponse)
	},
}

func NewBaseResponse() *BaseResponse {
	return baseResponsePool.Get().(*BaseResponse)
}

func ReleaseBaseResponse(v *BaseResponse) {
	v.Code = 0
	v.Msg = ""
	baseResponsePool.Put(v)
}

type ResponseWrapper struct {
	BaseResponse
	Data json.RawMessage `json:"data,omitempty"`
}

var responsePool = sync.Pool{
	New: func() any {
		return new(ResponseWrapper)
	},
}

func NewResponseWrapper() *ResponseWrapper {
	return responsePool.Get().(*ResponseWrapper)
}

func ReleaseResponseWrapper(v *ResponseWrapper) {
	v.Code = 0
	v.Msg = ""
	v.Data = nil
	responsePool.Put(v)
}
