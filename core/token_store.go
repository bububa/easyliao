package core

import (
	"context"
	"sync"

	"github.com/bububa/easyliao/model/oauth"
)

type TokenStore interface {
	Get(ctx context.Context, token *oauth.AccessToken) error
	Set(ctx context.Context, token *oauth.AccessToken) error
}

type MemTokenStore struct {
	token *oauth.AccessToken
	mutex *sync.RWMutex
}

func NewMemTokenStore() *MemTokenStore {
	return &MemTokenStore{
		mutex: new(sync.RWMutex),
	}
}

func (t *MemTokenStore) Get(ctx context.Context, token *oauth.AccessToken) error {
	t.mutex.RLock()
	defer t.mutex.RUnlock()
	if t.token == nil {
		token = nil
		return nil
	}
	*token = *t.token
	return nil
}

func (t *MemTokenStore) Set(ctx context.Context, token *oauth.AccessToken) error {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	t.token = token
	return nil
}
