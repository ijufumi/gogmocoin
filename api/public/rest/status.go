package rest

import (
	"encoding/json"
	"fmt"
	"github.com/ijufumi/gogmocoin/api/common/api"
	"net/url"

	"github.com/ijufumi/gogmocoin/api/public/rest/model"
)

// Status ...
type Status interface {
	Status() (*model.StatusRes, error)
}

func newStatus() status {
	return status{
		RestAPIBase: api.NewRestAPIBase(),
	}
}

type status struct {
	api.RestAPIBase
}

func (s *status) Status() (*model.StatusRes, error) {
	res, err := s.Get(url.Values{}, "/v1/status")
	if err != nil {
		return nil, err
	}

	statusRes := new(model.StatusRes)
	err = json.Unmarshal(res, statusRes)
	if err != nil {
		return nil, fmt.Errorf("[Status]error:%v,body:%s", err, res)
	}
	if len(statusRes.Messages) != 0 {
		return nil, fmt.Errorf("%v", statusRes.Messages)
	}

	return statusRes, nil
}
