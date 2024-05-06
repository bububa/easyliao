package core

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"sync"
	"time"

	"github.com/bububa/easyliao/core/internal/debug"
	"github.com/bububa/easyliao/model"
	"github.com/bububa/easyliao/model/oauth"
	"github.com/bububa/easyliao/util"
)

var (
	onceInit   sync.Once
	httpClient *http.Client
)

func defaultHttpClient() *http.Client {
	onceInit.Do(func() {
		transport := http.DefaultTransport.(*http.Transport).Clone()
		transport.MaxIdleConns = 100
		transport.MaxConnsPerHost = 100
		transport.MaxIdleConnsPerHost = 100
		httpClient = &http.Client{
			Transport: transport,
			Timeout:   time.Second * 60,
		}
	})
	return httpClient
}

// SDKClient sdk client
type SDKClient struct {
	AppID      string
	Secret     string
	tokenStore TokenStore
	client     *http.Client
	debug      bool
}

// NewSDKClient 创建SDKClient
func NewSDKClient(appID string, secret string) *SDKClient {
	return &SDKClient{
		AppID:      appID,
		Secret:     secret,
		client:     defaultHttpClient(),
		tokenStore: NewMemTokenStore(),
	}
}

// SetDebug 设置debug模式
func (c *SDKClient) SetDebug(debug bool) {
	c.debug = debug
}

// SetHttpClient 设置http.Client
func (c *SDKClient) SetHttpClient(client *http.Client) {
	c.client = client
}

func (c *SDKClient) SetTokenStore(store TokenStore) {
	c.tokenStore = store
}

// Post post api
func (c *SDKClient) Post(ctx context.Context, gw string, req model.PostRequest, resp model.Response) error {
	var reqBytes []byte
	if req != nil {
		reqBytes = req.Encode()
	}
	reqUrl := util.StringsJoin(BASE_URL, gw)
	httpReq, err := http.NewRequest("POST", reqUrl, bytes.NewReader(reqBytes))
	if err != nil {
		return err
	}
	httpReq.Header.Add("Content-Type", "application/json")
	token := new(oauth.AccessToken)
	c.tokenStore.Get(ctx, token)
	if token == nil || !token.Valid() {
		if err := c.GetToken(ctx, token); err != nil {
			return err
		}
	}
	httpReq.Header.Add("token", token.Token)
	debug.PrintJSONRequest("POST", reqUrl, httpReq.Header, reqBytes, c.debug)
	if err := c.fetch(httpReq, resp); err != nil {
		return err
	}
	return nil
}

// Get get api
func (c *SDKClient) Get(ctx context.Context, gw string, req model.GetRequest, resp model.Response) error {
	reqUrl := util.StringsJoin(BASE_URL, gw)
	if req != nil {
		reqUrl = util.StringsJoin(reqUrl, "?", req.Encode())
	}
	debug.PrintGetRequest(reqUrl, c.debug)
	httpReq, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		return err
	}
	token := new(oauth.AccessToken)
	c.tokenStore.Get(ctx, token)
	if token == nil || !token.Valid() {
		if err := c.GetToken(ctx, token); err != nil {
			return err
		}
	}
	httpReq.Header.Add("token", token.Token)
	if err := c.fetch(httpReq, resp); err != nil {
		return err
	}
	return nil
}

func (c *SDKClient) GetToken(ctx context.Context, token *oauth.AccessToken) error {
	values := util.GetUrlValues()
	values.Set("appid", c.AppID)
	values.Set("secret", c.Secret)
	reqUrl := util.StringsJoin(BASE_URL, "gettoken?", values.Encode())
	util.PutUrlValues(values)
	debug.PrintGetRequest(reqUrl, c.debug)
	httpReq, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		return err
	}
	if err := c.fetch(httpReq, token); err != nil {
		return err
	}
	token.ExpireAt = time.Now().Add(time.Second * time.Duration(token.Expire)).Unix()
	return c.tokenStore.Set(ctx, token)
}

// fetch execute http request
func (c *SDKClient) fetch(httpReq *http.Request, resp model.Response) error {
	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return err
	}
	defer httpResp.Body.Close()
	base := model.NewResponseWrapper()
	defer model.ReleaseResponseWrapper(base)
	if body, err := debug.DecodeJSONHttpResponse(httpResp.Body, base, c.debug); err != nil {
		debug.PrintError(err, c.debug)
		return errors.Join(err, model.BaseResponse{
			Code: httpResp.StatusCode,
			Msg:  string(body),
		})
	}
	if base.IsError() {
		return base
	}
	if resp != nil {
		if err := json.Unmarshal([]byte(base.Data), resp); err != nil {
			return err
		}
		if resp.IsError() {
			return resp
		}
	}
	return nil
}
