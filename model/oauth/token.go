package oauth

import (
	"time"
)

// AccessToken 授权token Respnse
type AccessToken struct {
	// Token token
	Token string `json:"token,omitempty"`
	// Expire 失效时间（单位：秒）
	Expire int64 `json:"expire,omitempty"`
	// ExpireAt 失效时间
	ExpireAt int64 `json:"expire_at,omitempty"`
}

func (t AccessToken) Valid() bool {
	return t.Token != "" && t.ExpireAt > time.Now().Unix()
}

func (t AccessToken) IsError() bool {
	return t.Valid()
}

func (t AccessToken) Error() string {
	return "invalid token"
}
