package api

import (
	"context"

	"github.com/bububa/easyliao/core"
	"github.com/bububa/easyliao/model/invoke"
)

// VisitorCard 历史名片
func VisitorCard(ctx context.Context, clt *core.SDKClient, req *invoke.VisitorCardRequest) (*invoke.VisitorCardResponse, error) {
	var resp invoke.VisitorCardResponse
	if err := clt.Post(ctx, "base/opendata/api/invoke", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// VisitorCardByVsID 根据访客静态id查询名片
func VisitorCardByVsID(ctx context.Context, clt *core.SDKClient, req *invoke.VisitorCardByVsIDRequest) ([]invoke.VisitorCard, error) {
	var resp invoke.VisitorCardResponse
	if err := clt.Post(ctx, "base/opendata/api/invoke", req, &resp); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
