package invoke

import (
	"encoding/json"

	"github.com/bububa/easyliao/model"
)

// ChatRecordRequest 访客对话记录查询 API Request
type ChatRecordRequest struct {
	// DataType 固定值：chatRecord
	DataType string `json:"dataType,omitempty"`
	// StartTime 开始时间
	StartTime string `json:"startTime,omitempty"`
	// EndTime 结束时间
	EndTime string `json:"endTime,omitempty"`
	// PageNumber 指定查询页
	PageNumber model.JSONInt64 `json:"pageNumber,omitempty"`
}

// Encode implmenents PostRequest
func (r ChatRecordRequest) Encode() []byte {
	r.DataType = "chatRecord"
	bs, _ := json.Marshal(r)
	return bs
}

// ChatRecordByIDRequest 根据chatId查询对话记录 API Request
type ChatRecordByIDRequest struct {
	// DataType 固定值：chatRecord
	DataType string `json:"dataType,omitempty"`
	// StartTime 开始时间
	StartTime string `json:"startTime,omitempty"`
	// EndTime 结束时间
	EndTime string `json:"endTime,omitempty"`
	// ChatID 对话ID
	ChatID model.JSONUint64 `json:"chatId,omitempty"`
}

// Encode implmenents PostRequest
func (r ChatRecordByIDRequest) Encode() []byte {
	r.DataType = "chatRecordByChatId"
	bs, _ := json.Marshal(r)
	return bs
}

// 根据对话质量查询对话记录 API Request
type ChatRecordByTalkQualityRequest struct {
	// DataType 固定值：chatRecord
	DataType string `json:"dataType,omitempty"`
	// StartTime 开始时间
	StartTime string `json:"startTime,omitempty"`
	// EndTime 结束时间
	EndTime string `json:"endTime,omitempty"`
	// PageNumber 指定查询页
	PageNumber model.JSONInt64 `json:"pageNumber,omitempty"`
	// TalkQuality "对话质量
	// 0：全部1：有效2：无效"
	TakQuality model.JSONInt `json:"talkQuality,omitempty"`
}

// Encode implmenents PostRequest
func (r ChatRecordByTalkQualityRequest) Encode() []byte {
	r.DataType = "chatRecordByTalkQuality"
	bs, _ := json.Marshal(r)
	return bs
}

// ChatRecordResponse 访客对话记录查询 API Response
type ChatRecordResponse struct {
	model.BaseResponse
	// Total 数据总数
	Total int64 `json:"total,omitempty"`
	// PageNumber 当前页数
	PageNumber int64 `json:"pageNumber,omitempty"`
	// PageCount 数据总页数
	PageCount int64 `json:"pageCount,omitempty"`
	// PageSize 每页数量[默认1000暂不提供修改]
	PageSize int64 `json:"pageSize,omitempty"`
	// Data 数据列表
	Data ChatRecordList `json:"data,omitempty"`
}

type ChatRecordList []ChatRecord

// UnmarshalJSON implement json Unmarshal interface
func (l *ChatRecordList) UnmarshalJSON(b []byte) (err error) {
	if b[0] == '"' && b[len(b)-1] == '"' {
		var str string
		if err := json.Unmarshal(b, &str); err != nil {
			return err
		}
		b = []byte(str)
	}
	var list []ChatRecord
	if err := json.Unmarshal(b, &list); err != nil {
		return err
	}
	*l = ChatRecordList(list)
	return nil
}

// ChatRecord 访客对话记录
type ChatRecord struct {
	// VisitorStaticID 访客静态ID（唯一值）
	VisitorStaticID string `json:"visitorStaticId,omitempty"`
	// CompanyID 公司ID
	CompanyID uint64 `json:"companyId,omitempty"`
	// ChatID 会话数据id
	ChatID uint64 `json:"chatId,omitempty"`
	// UserRealName 访客真实姓名
	UserRealName string `json:"userRealName,omitempty"`
	// GroupID 接待分组ID
	GroupID uint64 `json:"groupId,omitempty"`
	// VisitorMsgCount 访客说话总数
	VisitorMsgCount int64 `json:"visitorMsgCount,omitempty"`
	// VisitorFirstTime 访客首次说话时间（时间戳/毫秒）
	VisitorFirstTime int64 `json:"visitorFirstTime,omitempty"`
	// SiteName 子站点名称
	SiteName string `json:"siteName,omitempty"`
	// SearchHost 搜索引擎（网站，www.sogou.com）
	SearchHost string `json:"searchHost,omitempty"`
	// Effective 会话效果（0为无效，1为有效）
	Effective int `json:"effective,omitempty"`
	// UserMsgCount 客服说话总数
	UserMsgCount int64 `json:"userMsgCount,omitempty"`
	// UserFirstMsgTime 客服首次回复时间（时间戳/毫秒）
	UserFirstMsgTime int64 `json:"userFirstMsgTime,omitempty"`
	// SummarizeName 对话总结标签名称
	SummarizeName string `json:"summarizeName,omitempty"`
	// Keyword 搜索词
	Keyword string `json:"keyword,omitempty"`
	// BcpBiddingWord 关键词
	BcpBiddingWord string `json:"bcpBiddingWord,omitempty"`
	// VisitorLocationCountry 访客--国家
	VisitorLocationCountry string `json:"visitorLocationCountry,omitempty"`
	// VisitorLocationProvince 访客--省
	VisitorLocationProvince string `json:"visitorLocationProvince,omitempty"`
	// VisitorLocationCity 访客--市
	VisitorLocationCity string `json:"visitorLocationCity,omitempty"`
	// UserTimeoutTimes 客服回复超时次数
	UserTimeoutTimes int64 `json:"userTimeoutTimes,omitempty"`
	// ChatType 设备类型（pc、phone）
	ChatType string `json:"chatType,omitempty"`
	// VisitorIP 访客--IP
	VisitorIP string `json:"visitorIp,omitempty"`
	// FirstURL 最初访问页
	FirstURL string `json:"firstUrl,omitempty"`
	// SummarizeID 对话总结ID
	SummarizeID uint64 `json:"summarizeId,omitempty"`
	// InviteMode 发起方式（1客服发起2：访客发起3：内部会议室4：会话转移，5外部会议室）
	InviteMode int `json:"inviteMode,omitempty"`
	// LastMessageTime 最后一条消息时间（时间戳/毫秒）
	LastMessageTime int64 `json:"lastMessageTime,omitempty"`
	// ChoseType 结束方式（1：访客，2：客服，3：会话转移，4：超时）
	ChoseType int `json:"choseType,omitempty"`
	// UserID 客服ID
	UserID string `json:"userId,omitempty"`
	// AllTime 对话时间（秒）
	AllTime int64 `json:"allTime,omitempty"`
	// CreateTime 对话创建时间（时间戳/毫秒）
	CreateTime int64 `json:"createTime,omitempty"`
	// EndTime 结束时间
	EndTime int64 `json:"endTime,omitempty"`
	// EffectTime 有效沟通时间（秒）
	EffectTime int64 `json:"effectTime,omitempty"`
	// ChatCategory 对话类型
	ChatCategory uint64 `json:"chatCategory,omitempty"`
	// SiteID 子站点id
	SiteID uint64 `json:"siteId,omitempty"`
	// ChatURL 对话页
	ChatURL string `json:"chatUrl,omitempty"`
	// MsgCount 对话总数
	MsgCount int64 `json:"msgCount,omitempty"`
	// ReferPage 来源页
	ReferPage string `json:"referPage,omitempty"`
	// VisitorID 访客ID
	VisitorID string `json:"visitorId,omitempty"`
	// ExtColum1 扩展字段1
	ExtColum1 string `json:"extColum1,omitempty"`
	// ExtColum2 扩展字段2
	ExtColum2 string `json:"extColum2,omitempty"`
	// ExtColum3 扩展字段3
	ExtColum3 string `json:"extColum3,omitempty"`
	// ExtColum4 扩展字段4
	ExtColum4 string `json:"extColum4,omitempty"`
	// ExtColum5 扩展字段5
	ExtColum5 string `json:"extColum5,omitempty"`
	// ExtColum6 扩展字段6
	ExtColum6 string `json:"extColum6,omitempty"`
}
