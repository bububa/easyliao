package api

import (
	"context"

	"github.com/bububa/easyliao/core"
	"github.com/bububa/easyliao/model/invoke"
)

// ChatDetail 对话记录详情
func ChatDetail(ctx context.Context, clt *core.SDKClient, req *invoke.ChatDetailRequest) ([]invoke.ChatDetail, error) {
	var resp invoke.ChatDetailResponse
	if err := clt.Post(ctx, "base/opendata/api/invoke", req, &resp); err != nil {
		return nil, err
	}
	return resp.Data, nil
}

// ChatDetailByID 根据chatId对话记录详情
func ChatDetailByID(ctx context.Context, clt *core.SDKClient, req *invoke.ChatDetailByIDRequest) ([]invoke.ChatDetail, error) {
	var resp invoke.ChatDetailResponse
	if err := clt.Post(ctx, "base/opendata/api/invoke", req, &resp); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
