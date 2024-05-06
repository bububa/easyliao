package api

import (
	"context"

	"github.com/bububa/easyliao/core"
	"github.com/bububa/easyliao/model/invoke"
)

// QueryNtagCategory 查询名片标签词典
func QueryNtagCategory(ctx context.Context, clt *core.SDKClient, req *invoke.QueryNtagCategoryRequest) ([]invoke.NtagCategory, error) {
	var resp invoke.QueryNtagCategoryResponse
	if err := clt.Post(ctx, "base/opendata/api/invoke", req, &resp); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
