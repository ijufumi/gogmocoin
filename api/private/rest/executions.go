package rest

import (
	"encoding/json"
	"fmt"
	"github.com/ijufumi/gogmocoin/v2/api/common/api"
	"github.com/ijufumi/gogmocoin/v2/api/private/rest/model"
	"net/url"
	"strconv"
)

// Executions ...
type Executions interface {
	Executions(orderID, executionID int64) (*model.ExecutionsRes, error)
	ExecutionsByOrderID(orderID int64) (*model.ExecutionsRes, error)
	ExecutionsByExecutionID(executionID int64) (*model.ExecutionsRes, error)
}

func newExecutions(apiKey, secretKey string) executions {
	return executions{
		RestAPIBase: api.NewPrivateRestAPIBase(apiKey, secretKey),
	}
}

type executions struct {
	api.RestAPIBase
}

func (e *executions) ExecutionsByOrderID(orderID int64) (*model.ExecutionsRes, error) {
	return e.Executions(orderID, 0)
}

func (e *executions) ExecutionsByExecutionID(executionID int64) (*model.ExecutionsRes, error) {
	return e.Executions(0, executionID)
}

func (e *executions) Executions(orderID, executionID int64) (*model.ExecutionsRes, error) {
	param := url.Values{}

	if orderID > 0 {
		param.Set("orderId", strconv.FormatInt(orderID, 10))
	}

	if executionID > 0 {
		param.Set("executionId", strconv.FormatInt(executionID, 10))
	}

	res, err := e.Get(param, "/v1/executions")
	if err != nil {
		return nil, err
	}

	executionsRes := new(model.ExecutionsRes)
	err = json.Unmarshal(res, executionsRes)
	if err != nil {
		return nil, fmt.Errorf("[Executions]error:%v,body:%s", err, res)
	}

	if len(executionsRes.Messages) != 0 {
		return nil, fmt.Errorf("%v", executionsRes.Messages)
	}

	return executionsRes, nil
}
