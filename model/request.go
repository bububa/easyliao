package model

// GetRequest Get请求接口
type GetRequest interface {
	Encode() string
}

// PostRequest Post请求接口
type PostRequest interface {
	Encode() []byte
}
