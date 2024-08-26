package internal

import (
	"context"
	"fmt"
	"github.com/ijufumi/gogmocoin/v2/api/common/consts"
	"github.com/ijufumi/gogmocoin/v2/api/internal/api"
	"github.com/ijufumi/gogmocoin/v2/api/private/rest"
	"sync/atomic"
	"time"
)

type PrivateWSAPIBase struct {
	wsAPIBase               *api.WSAPIBase
	restAPIClient           rest.Client
	ctx                     context.Context
	cancelFunc              context.CancelFunc
	token                   *atomic.Value
	tokenAutomaticExtension bool
}

type tokenData struct {
	Token      string
	ExpireTime time.Time
}

func NewPrivateWSAPIBase(apiKey, secretKey string, tokenAutomaticExtension bool, requestFactory api.RequestFactoryFunc) *PrivateWSAPIBase {
	wsAPIBase := &PrivateWSAPIBase{
		restAPIClient:           rest.NewWithKeys(apiKey, secretKey),
		wsAPIBase:               api.NewWSAPIBase(requestFactory),
		tokenAutomaticExtension: tokenAutomaticExtension,
	}
	wsAPIBase.SetHostFactoryFunc(wsAPIBase.hostFactory)
	return wsAPIBase
}

func (w *PrivateWSAPIBase) SetHostFactoryFunc(f api.HostFactoryFunc) {
	w.wsAPIBase.SetHostFactoryFunc(f)
}

func (w *PrivateWSAPIBase) Subscribe() error {
	if w.tokenAutomaticExtension {
		ctx, cancelFunc := context.WithCancel(context.Background())
		w.ctx = ctx
		w.cancelFunc = cancelFunc

		w.automaticExtension()
	}
	return w.wsAPIBase.Subscribe()
}

func (w *PrivateWSAPIBase) Unsubscribe() error {
	if w.tokenAutomaticExtension {
		w.cancelFunc()
	}
	return w.wsAPIBase.Unsubscribe()
}

func (w *PrivateWSAPIBase) automaticExtension() {
	ticker := time.NewTicker(50 * time.Minute)
	go func() {
		for {
			select {
			case <-ticker.C:

			case <-w.ctx.Done():
				return
			}
		}
	}()
}

func (w *PrivateWSAPIBase) createWSToken() error {
	token, err := w.restAPIClient.CreateWSAuthTokenWithContext(w.ctx)
	if err != nil {
		return err
	}
	w.token.Store(&tokenData{
		Token: token,
	})

	return nil
}

func (w *PrivateWSAPIBase) extendWSToken() error {
	return w.restAPIClient.ExtendWSAuthTokenWithContext(w.ctx, w.getToken())
}

func (w *PrivateWSAPIBase) revokeWSToken() error {
	return w.restAPIClient.RevokeWSAuthTokenWithContext(w.ctx, w.getToken())
}

func (w *PrivateWSAPIBase) hostFactory() string {
	return fmt.Sprintf("%s/%s", consts.PrivateWSAPIHost, w.getToken())
}

func (w *PrivateWSAPIBase) getToken() string {
	rawValue := w.token.Load()
	if rawValue == nil {
		return ""
	}
	tokenValue, ok := rawValue.(tokenData)
	if !ok {
		return ""
	}

	return tokenValue.Token
}
