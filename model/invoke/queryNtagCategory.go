package invoke

import (
	"encoding/json"

	"github.com/bububa/easyliao/model"
)

// QueryNtagCategoryRequest 查询名片标签词典 API Request
type QueryNtagCategoryRequest struct {
	// DataType 固定值：queryNtagCategory
	DataType string `json:"dataType,omitempty"`
}

// Encode implmenents PostRequest
func (r QueryNtagCategoryRequest) Encode() []byte {
	r.DataType = "queryNtagCategory"
	bs, _ := json.Marshal(r)
	return bs
}

// QueryNtagCategoryResponse 查询名片标签词典 API Response
type QueryNtagCategoryResponse struct {
	model.BaseResponse
	Data []NtagCategory `json:"data,omitempty"`
}

// NtagCategory 名片标签词典
type NtagCategory struct {
	// ID 标签ID
	ID uint64 `json:"id,omitempty"`
	// Name 标签名称
	Name string `json:"name,omitempty"`
	// Type 标签类型
	Type int `json:"type,omitempty"`
}
