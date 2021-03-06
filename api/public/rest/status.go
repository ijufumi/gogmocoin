package rest

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/ijufumi/gogmocoin/api/public/rest/internal/connect"
	"github.com/ijufumi/gogmocoin/api/public/rest/model"
)

// Status ...
type Status interface {
	Status() (*model.StatusRes, error)
}

type status struct {
	con connect.Connection
}

func (s *status) Status() (*model.StatusRes, error) {
	res, err := s.con.Get(url.Values{}, "/v1/status")
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
