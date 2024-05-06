package api

import (
	"context"

	"github.com/bububa/easyliao/core"
	"github.com/bububa/easyliao/model/invoke"
)

// ChatRecord 访客对话记录查询
func ChatRecord(ctx context.Context, clt *core.SDKClient, req *invoke.ChatRecordRequest) (*invoke.ChatRecordResponse, error) {
	var resp invoke.ChatRecordResponse
	if err := clt.Post(ctx, "base/opendata/api/invoke", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChatRecordByID 根据chatId查询对话记录
func ChatRecordByID(ctx context.Context, clt *core.SDKClient, req *invoke.ChatRecordByIDRequest) (*invoke.ChatRecordResponse, error) {
	var resp invoke.ChatRecordResponse
	if err := clt.Post(ctx, "base/opendata/api/invoke", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChatRecordByTalkQuality 根据对话质量查询对话记录
func ChatRecordByTalkQuality(ctx context.Context, clt *core.SDKClient, req *invoke.ChatRecordByTalkQualityRequest) (*invoke.ChatRecordResponse, error) {
	var resp invoke.ChatRecordResponse
	if err := clt.Post(ctx, "base/opendata/api/invoke", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
