package invoke

import (
	"encoding/json"

	"github.com/bububa/easyliao/model"
)

// ChatDetailRequest 对话记录详情查询 API Request
type ChatDetailRequest struct {
	// DataType 固定值：detailJson
	DataType string `json:"dataType,omitempty"`
	// StartTime 开始时间
	StartTime string `json:"startTime,omitempty"`
	// EndTime 结束时间
	EndTime string `json:"endTime,omitempty"`
	// VsID 访客静态ID
	VsID string `json:"vsId,omitempty"`
}

// Encode implmenents PostRequest
func (r ChatDetailRequest) Encode() []byte {
	r.DataType = "chatDetailJson"
	bs, _ := json.Marshal(r)
	return bs
}

// ChatDetailByIDRequest 对话记录详情查询 API Request
type ChatDetailByIDRequest struct {
	// DateType 固定值：detailJson
	DateType string `json:"dateType,omitempty"`
	// StartTime 开始时间
	StartTime string `json:"startTime,omitempty"`
	// ChatID 对话唯一标识
	ChatID model.JSONUint64 `json:"chatId,omitempty"`
}

// Encode implmenents PostRequest
func (r ChatDetailByIDRequest) Encode() []byte {
	r.DateType = "detailByChatId"
	bs, _ := json.Marshal(r)
	return bs
}

// ChatDetailResponse 对话记录详情查询 API Response
type ChatDetailResponse struct {
	model.BaseResponse
	// Data 数据列表
	Data ChatDetailList `json:"data,omitempty"`
}

type ChatDetailList []ChatDetail

// UnmarshalJSON implement json Unmarshal interface
func (l *ChatDetailList) UnmarshalJSON(b []byte) (err error) {
	if b[0] == '"' && b[len(b)-1] == '"' {
		var str string
		if err := json.Unmarshal(b, &str); err != nil {
			return err
		}
		b = []byte(str)
	}
	var list []ChatDetail
	if err := json.Unmarshal(b, &list); err != nil {
		return err
	}
	*l = ChatDetailList(list)
	return nil
}

// ChatDetail 对话记录详情
type ChatDetail struct {
	// SenderType 发送人（0:客服 1:访客）
	SenderType model.Int `json:"senderType,omitempty"`
	// Message 会话数据
	Message string `json:"message,omitempty"`
	// ChatID 会话ID
	ChatID model.Uint64 `json:"chatId,omitempty"`
	// CompanyID 公司ID
	CompanyID model.Uint64 `json:"companyId,omitempty"`
	// CreateTime 对话时间
	CreateTime string `json:"createTime,omitempty"`
	// Type 对话类型
	Type string `json:"type,omitempty"`
	// FromID 发送人
	FromID string `json:"fromId,omitempty"`
	// RecordID 对话详情ID
	RecordID model.Uint64 `json:"recordId,omitempty"`
}
