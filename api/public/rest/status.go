package rest

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ijufumi/gogmocoin/v2/api/common/api"
	"github.com/ijufumi/gogmocoin/v2/api/public/rest/model"
	"net/url"
)

// Status ...
type Status interface {
	Status() (*model.StatusRes, error)
	StatusWithContext(ctx context.Context) (*model.StatusRes, error)
}

func newStatus() status {
	return status{
		RestAPIBase: api.NewRestAPIBase(),
	}
}

type status struct {
	api.RestAPIBase
}

// Status ...
func (s *status) Status() (*model.StatusRes, error) {
	return s.StatusWithContext(context.Background())
}

// StatusWithContext ...
func (s *status) StatusWithContext(ctx context.Context) (*model.StatusRes, error) {
	res, err := s.Get(ctx, url.Values{}, "/v1/status")
	if err != nil {
		return nil, err
	}

	statusRes := new(model.StatusRes)
	err = json.Unmarshal(res, statusRes)
	if err != nil {
		return nil, fmt.Errorf("[Status]error:%v,body:%s", err, res)
	}
	if !statusRes.Success() {
		return nil, statusRes.Error()
	}

	return statusRes, nil
}
