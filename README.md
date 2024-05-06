# 易聊 Golang SDK

[![Go Reference](https://pkg.go.dev/badge/github.com/bububa/easyliao.svg)](https://pkg.go.dev/github.com/bububa/easyliao)
[![Go](https://github.com/bububa/easyliao/actions/workflows/go.yml/badge.svg)](https://github.com/bububa/easyliao/actions/workflows/go.yml)
[![goreleaser](https://github.com/bububa/easyliao/actions/workflows/goreleaser.yml/badge.svg)](https://github.com/bububa/easyliao/actions/workflows/goreleaser.yml)
[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/bububa/easyliao.svg)](https://github.com/bububa/easyliao)
[![GoReportCard](https://goreportcard.com/badge/github.com/bububa/easyliao)](https://goreportcard.com/report/github.com/bububa/easyliao)
[![GitHub license](https://img.shields.io/github/license/bububa/easyliao.svg)](https://github.com/bububa/easyliao/blob/master/LICENSE)
[![GitHub release](https://img.shields.io/github/release/bububa/easyliao.svg)](https://GitHub.com/bububa/easyliao/releases/)

- 获取token [ (c *SDKClient) GetToken(ctx context.Context, token *oauth.AccessToken) error ]
- 访客对话记录查询 [ ChatRecord(ctx context.Context, clt *core.SDKClient, req *invoke.ChatRecordRequest) (*invoke.ChatRecordResponse, error)  ]
- 根据chatId查询对话记录 [ ChatRecordByID(ctx context.Context, clt *core.SDKClient, req *invoke.ChatRecordByIDRequest) (*invoke.ChatRecordResponse, error) ]
- 根据对话质量查询对话记录 [ ChatRecordByTalkQuality(ctx context.Context, clt *core.SDKClient, req *invoke.ChatRecordByTalkQualityRequest) (*invoke.ChatRecordResponse, error) ]
- 对话记录详情 [ ChatDetail(ctx context.Context, clt *core.SDKClient, req *invoke.ChatRecordRequest) ([]invoke.ChatDetail, error) ]
- 根据chatId对话记录详情 [ ChatDetailByID(ctx context.Context, clt *core.SDKClient, req *invoke.ChatDetailByIDRequest) ([]invoke.ChatDetail, error) ]
- 历史名片 [ VisitorCard(ctx context.Context, clt *core.SDKClient, req *invoke.VisitorCardRequest) (*invoke.VisitorCardResponse, error) ]
- 根据访客静态id查询名片 [ VisitorCardByVsID(ctx context.Context, clt *core.SDKClient, req *invoke.VisitorCardByVsIDRequest) ([]invoke.VisitorCard, error) ]
- 查询名片标签词典 [ QueryNtagCategory(ctx context.Context, clt *core.SDKClient, req *invoke.QueryNtagCategoryRequest) ([]invoke.NtagCategory, error) ]
