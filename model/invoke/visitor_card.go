package invoke

import (
	"encoding/json"

	"github.com/bububa/easyliao/model"
)

// VisitorCardRequest 历史名片查询 API Request
type VisitorCardRequest struct {
	// DataType 固定值：visitorCard
	DataType string `json:"dataType,omitempty"`
	// StartTime 查询开始时间[结束时间默认为当前时间]
	StartTime string `json:"startTime,omitempty"`
	// EndTime 查询结束时间
	EndTime string `json:"endTime,omitempty"`
	// PageNumber 页码数
	PageNumber model.JSONInt64 `json:"pageNumber,omitempty"`
}

// Encode implmenents PostRequest
func (r VisitorCardRequest) Encode() []byte {
	r.DataType = "visitorCard"
	bs, _ := json.Marshal(r)
	return bs
}

// VisitorCardByVsIDRequest 根据访客静态id查询名片 API Request
type VisitorCardByVsIDRequest struct {
	// DataType 固定值：visitorCard
	DataType string `json:"dataType,omitempty"`
	// VisitorStaticID 访客静态ID
	VisitorStaticID string `json:"visitorStaticId,omitempty"`
}

// Encode implmenents PostRequest
func (r VisitorCardByVsIDRequest) Encode() []byte {
	r.DataType = "newVisitorCardByStaticId"
	bs, _ := json.Marshal(r)
	return bs
}

// VisitorCardResponse 历史名片查询 API Response
type VisitorCardResponse struct {
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
	Data VisitorCardList `json:"data,omitempty"`
}

type VisitorCardList []VisitorCard

// UnmarshalJSON implement json Unmarshal interface
func (l *VisitorCardList) UnmarshalJSON(b []byte) (err error) {
	if b[0] == '"' && b[len(b)-1] == '"' {
		var str string
		if err := json.Unmarshal(b, &str); err != nil {
			return err
		}
		b = []byte(str)
	}
	var list []VisitorCard
	if err := json.Unmarshal(b, &list); err != nil {
		return err
	}
	*l = VisitorCardList(list)
	return nil
}

// VisitorCard 访客名片
type VisitorCard struct {
	// ID 名片id
	ID model.Uint64 `json:"id,omitempty"`
	// VisitorStaticID 访客静态ID（唯一值）
	VisitorStaticID string `json:"visitorStaticId,omitempty"`
	// CompanyID 公司ID
	CompanyID model.Uint64 `json:"companyId,omitempty"`
	// ChatID 会话数据id
	ChatID model.Uint64 `json:"chatId,omitempty"`
	// CompanyName 公司名称
	CompanyName string `json:"companyName,omitempty"`
	// Name 名称（名片）
	Name string `json:"name,omitempty"`
	// Email 电子邮箱
	Email string `json:"email,omitempty"`
	// Weixin 微信
	Weixin string `json:"weixin,omitempty"`
	// QQ
	QQ string `json:"qq,omitempty"`
	// Sex 性别
	Sex string `json:"sex,omitempty"`
	// IP ip
	IP string `json:"ip,omitempty"`
	// Mobile 手机号
	Mobile string `json:"mobile,omitempty"`
	// UserID 客服ID
	UserID string `json:"userId,omitempty"`
	// CreateTime 对话创建时间（时间戳/毫秒）
	CreateTime int64 `json:"createTime,omitempty"`
	// EditTime 访客首次说话时间
	EditTime string `json:"editTime,omitempty"`
	// Province 省
	Province string `json:"province,omitempty"`
	// City 市
	City string `json:"city,omitempty"`
	// Area 地区
	Area string `json:"area,omitempty"`
	// URL Url
	URL string `json:"url,omitempty"`
	// ChatURL 对话页
	ChatURL string `json:"chatUrl,omitempty"`
	// FirstURL 最初访问页
	FirstURL string `json:"firstUrl,omitempty"`
	// Refer 来源页
	Refer string `json:"refer,omitempty"`
	// SearchHost 搜索引擎（网站，www.sogou.com）
	SearchHost string `json:"searchHost,omitempty"`
	// Tel 电话
	Tel string `json:"tel,omitempty"`
	// Keyword 关键词
	Keyword string `json:"keyword,omitempty"`
	// Note 备注
	Note string `json:"note,omitempty"`
	// TgAdAccountID 账户id
	TgAdAccountID model.Uint64 `json:"tgAdAccountId,omitempty"`
	// TgPlanID 计划id
	TgPlanID model.Uint64 `json:"tgPlanId,omitempty"`
	// TgGroupID 单元id
	TgGroupID model.Uint64 `json:"tgGroupId,omitempty"`
	// TgCreativeID 推广创意ID
	TgCreativeID model.Uint64 `json:"tgCreativeId,omitempty"`
	// TgKeywordID 关键词id
	TgKeywordID model.Uint64 `json:"tgKeywordId,omitempty"`
	// TgAdProvider 提供商id
	TgAdProvider string `json:"tgAdProvider,omitempty"`
	// TgAdAgent 代理商
	TgAdAgent string `json:"tgAdAgent,omitempty"`
	// TgAdDevice 设备
	TgAdDevice string `json:"tgAdDevice,omitempty"`
	// AllocationTime 名片分配时间
	AllocationTime string `json:"allocationTime,omitempty"`
	// NtagID 标签id
	NtagID model.Uint64 `json:"ntagId,omitempty"`
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
	// ExtColum7 扩展字段7
	ExtColum7 string `json:"extColum7,omitempty"`
	// ExtColum8 扩展字段8
	ExtColum8 string `json:"extColum8,omitempty"`
	// ExtColum9 扩展字段9
	ExtColum9 string `json:"extColum9,omitempty"`
	// ExtColum10 扩展字段10
	ExtColum10 string `json:"extColum10,omitempty"`
}
