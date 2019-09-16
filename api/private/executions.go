package private

import (
	"api_client/api/private/internal/connect"
	"api_client/api/private/model"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
)

// Executions ...
type Executions interface {
	Executions(orderID, executionID int64) (*model.ExecutionsRes, error)
	ExecutionsByOrderID(orderID int64) (*model.ExecutionsRes, error)
	ExecutionsByExecutionID(executionID int64) (*model.ExecutionsRes, error)
}

type executions struct {
	con *connect.Connection
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

	res, err := e.con.Get(param, "/v1/executions")
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
