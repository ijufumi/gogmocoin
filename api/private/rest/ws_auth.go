package rest

import (
	"context"
	"encoding/json"
	"github.com/ijufumi/gogmocoin/v2/api/internal/api"
	"github.com/ijufumi/gogmocoin/v2/api/private/rest/model"
	"net/url"
)

// WSAuth is the client interface for managing private WebSocket access tokens.
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

// CreateWSAuthToken issues a new access token for private WebSocket channels using a background context.
func (w *wsAuth) CreateWSAuthToken() (string, error) {
	return w.CreateWSAuthTokenWithContext(context.Background())
}

// CreateWSAuthTokenWithContext issues a new access token for private WebSocket channels using the provided context.
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

// ExtendWSAuthToken extends the lifetime of the given WebSocket access token using a background context.
func (w *wsAuth) ExtendWSAuthToken(token string) error {
	return w.ExtendWSAuthTokenWithContext(context.Background(), token)
}

// ExtendWSAuthTokenWithContext extends the lifetime of the given WebSocket access token using the provided context.
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

// RevokeWSAuthToken revokes the given WebSocket access token using a background context.
func (w *wsAuth) RevokeWSAuthToken(token string) error {
	return w.RevokeWSAuthTokenWithContext(context.Background(), token)
}

// RevokeWSAuthTokenWithContext revokes the given WebSocket access token using the provided context.
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
