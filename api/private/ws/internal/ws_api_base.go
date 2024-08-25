package internal

import "github.com/ijufumi/gogmocoin/v2/api/internal/api"

type PrivateWSAPIBase struct {
	wsAPIBase   *api.WSAPIBase
	restAPIBase api.RestAPIBase
}

func NewPrivateWSAPIBase(apiKey, secretKey string, requestFactory api.RequestFactoryFunc) *PrivateWSAPIBase {
	return &PrivateWSAPIBase{
		restAPIBase: api.NewPrivateRestAPIBase(apiKey, secretKey),
		wsAPIBase:   api.NewWSAPIBase(requestFactory),
	}
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
