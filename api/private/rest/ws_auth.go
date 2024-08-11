package rest

import (
	"context"
	"encoding/json"
	"github.com/ijufumi/gogmocoin/v2/api/common/api"
	"github.com/ijufumi/gogmocoin/v2/api/private/rest/model"
	"net/url"
)

// WSAuth ...
type WSAuth interface {
	CreateWSAuthToken() (string, error)
	CreateWSAuthTokenWithContext(ctx context.Context) (string, error)
	ExtendWSAuthToken(token string) error
	ExtendWSAuthTokenWithContext(ctx context.Context, token string) error
	RevokeWSAuthToken(token string) error
	RevokeWSAuthTokenWithContext(ctx context.Context, token string) error
}

func newWSAuth(apiKey, secretKey string) wsAuth {
	return wsAuth{
		RestAPIBase: api.NewPrivateRestAPIBase(apiKey, secretKey),
	}
}

type wsAuth struct {
	api.RestAPIBase
}

// CreateWSAuthToken ...
func (w *wsAuth) CreateWSAuthToken() (string, error) {
	return w.CreateWSAuthTokenWithContext(context.Background())
}

// CreateWSAuthTokenWithContext ...
func (w *wsAuth) CreateWSAuthTokenWithContext(ctx context.Context) (string, error) {
	res, err := w.Post(ctx, url.Values{}, "/v1/ws-auth")
	if err != nil {
		return "", err
	}
	var wsTokenRes *model.WSAuthRes
	err = json.Unmarshal(res, &wsTokenRes)
	if err != nil {
		return "", err
	}
	if !wsTokenRes.Success() {
		return "", wsTokenRes.Error()
	}

	return wsTokenRes.Data, nil
}

// ExtendWSAuthToken ...
func (w *wsAuth) ExtendWSAuthToken(token string) error {
	return w.ExtendWSAuthTokenWithContext(context.Background(), token)
}

// ExtendWSAuthTokenWithContext ...
func (w *wsAuth) ExtendWSAuthTokenWithContext(ctx context.Context, token string) error {
	req := model.WSAuthReq{Token: token}
	res, err := w.Put(ctx, req, "/v1/ws-auth")
	if err != nil {
		return err
	}
	var wsTokenRes *model.WSAuthRes
	err = json.Unmarshal(res, &wsTokenRes)
	if err != nil {
		return err
	}
	if !wsTokenRes.Success() {
		return wsTokenRes.Error()
	}

	return nil
}

// RevokeWSAuthToken ...
func (w *wsAuth) RevokeWSAuthToken(token string) error {
	return w.RevokeWSAuthTokenWithContext(context.Background(), token)
}

// RevokeWSAuthTokenWithContext ...
func (w *wsAuth) RevokeWSAuthTokenWithContext(ctx context.Context, token string) error {
	req := model.WSAuthReq{Token: token}
	res, err := w.Delete(ctx, req, "/v1/ws-auth")
	if err != nil {
		return err
	}
	var wsTokenRes *model.WSAuthRes
	err = json.Unmarshal(res, &wsTokenRes)
	if err != nil {
		return err
	}
	if !wsTokenRes.Success() {
		return wsTokenRes.Error()
	}

	return nil
}
