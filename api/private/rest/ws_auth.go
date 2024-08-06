package rest

import (
	"encoding/json"
	"github.com/ijufumi/gogmocoin/v2/api/common/api"
	"github.com/ijufumi/gogmocoin/v2/api/private/rest/model"
	"net/url"
)

// WSAuth ...
type WSAuth interface {
	CreateWSAuthToken() (string, error)
	ExtendWSAuthToken(token string) error
	RevokeWSAuthToken(token string) error
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
	res, err := w.Post(url.Values{}, "/v1/ws-auth")
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
	req := model.WSAuthReq{Token: token}
	res, err := w.Put(req, "/v1/ws-auth")
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
	req := model.WSAuthReq{Token: token}
	res, err := w.Delete(req, "/v1/ws-auth")
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
