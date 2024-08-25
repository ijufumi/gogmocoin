package internal

import (
	"fmt"
	"github.com/ijufumi/gogmocoin/v2/api/common/consts"
	"github.com/ijufumi/gogmocoin/v2/api/internal/api"
	"sync/atomic"
	"time"
)

type PrivateWSAPIBase struct {
	wsAPIBase   *api.WSAPIBase
	restAPIBase api.RestAPIBase
	token       *atomic.Value
}

type token struct {
	Token      string
	ExpireTime time.Time
}

func NewPrivateWSAPIBase(apiKey, secretKey string, requestFactory api.RequestFactoryFunc) *PrivateWSAPIBase {
	wsAPIBase := &PrivateWSAPIBase{
		restAPIBase: api.NewPrivateRestAPIBase(apiKey, secretKey),
		wsAPIBase:   api.NewWSAPIBase(requestFactory),
	}
	wsAPIBase.SetHostFactoryFunc(wsAPIBase.hostFactory)
	return wsAPIBase
}

func (w *PrivateWSAPIBase) SetHostFactoryFunc(f api.HostFactoryFunc) {
	w.wsAPIBase.SetHostFactoryFunc(f)
}

func (w *PrivateWSAPIBase) Subscribe() error {
	return w.wsAPIBase.Subscribe()
}

func (w *PrivateWSAPIBase) Unsubscribe() error {
	return w.wsAPIBase.Unsubscribe()
}

func (w *PrivateWSAPIBase) hostFactory() string {
	rawValue := w.token.Load()
	if rawValue == nil {
		return consts.PrivateWSAPIHost
	}
	tokenValue, ok := rawValue.(token)
	if !ok {
		return consts.PrivateWSAPIHost
	}
	return fmt.Sprintf("%s/%s", consts.PrivateWSAPIHost, tokenValue.Token)
}
