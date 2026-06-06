package rest

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ijufumi/gogmocoin/v2/api/internal/api"
	"github.com/ijufumi/gogmocoin/v2/api/private/rest/model"
	"net/url"
	"strconv"
)

// Executions is the client interface for the executions (fills) endpoint.
type Executions interface {
	Executions(orderID, executionID int64) (*model.ExecutionsRes, error)
	ExecutionsWithContext(ctx context.Context, orderID, executionID int64) (*model.ExecutionsRes, error)
	ExecutionsByOrderID(orderID int64) (*model.ExecutionsRes, error)
	ExecutionsByOrderIDWithContext(ctx context.Context, orderID int64) (*model.ExecutionsRes, error)
	ExecutionsByExecutionID(executionID int64) (*model.ExecutionsRes, error)
	ExecutionsByExecutionIDWithContext(ctx context.Context, executionID int64) (*model.ExecutionsRes, error)
}

func newExecutions(apiKey, secretKey string) executions {
	return executions{
		RestAPIBase: api.NewPrivateRestAPIBase(apiKey, secretKey),
	}
}

type executions struct {
	api.RestAPIBase
}

// ExecutionsByOrderID retrieves executions for the given order ID using a background context.
func (e *executions) ExecutionsByOrderID(orderID int64) (*model.ExecutionsRes, error) {
	return e.Executions(orderID, 0)
}

// ExecutionsByOrderIDWithContext retrieves executions for the given order ID using the provided context.
func (e *executions) ExecutionsByOrderIDWithContext(ctx context.Context, orderID int64) (*model.ExecutionsRes, error) {
	return e.ExecutionsWithContext(ctx, orderID, 0)
}

// ExecutionsByExecutionID retrieves the execution with the given execution ID using a background context.
func (e *executions) ExecutionsByExecutionID(executionID int64) (*model.ExecutionsRes, error) {
	return e.Executions(0, executionID)
}

// ExecutionsByExecutionIDWithContext retrieves the execution with the given execution ID using the provided context.
func (e *executions) ExecutionsByExecutionIDWithContext(ctx context.Context, executionID int64) (*model.ExecutionsRes, error) {
	return e.ExecutionsWithContext(ctx, 0, executionID)
}

// Executions retrieves executions filtered by order ID or execution ID using a background context.
func (e *executions) Executions(orderID, executionID int64) (*model.ExecutionsRes, error) {
	return e.ExecutionsWithContext(context.Background(), orderID, executionID)
}

// ExecutionsWithContext retrieves executions filtered by order ID or execution ID using the provided context.
func (e *executions) ExecutionsWithContext(ctx context.Context, orderID, executionID int64) (*model.ExecutionsRes, error) {
	param := url.Values{}

	if orderID > 0 {
		param.Set("orderId", strconv.FormatInt(orderID, 10))
	}

	if executionID > 0 {
		param.Set("executionId", strconv.FormatInt(executionID, 10))
	}

	res, err := e.Get(ctx, param, "/v1/executions")
	if err != nil {
		return nil, err
	}

	executionsRes := new(model.ExecutionsRes)
	err = json.Unmarshal(res, executionsRes)
	if err != nil {
		return nil, fmt.Errorf("[Executions]error:%v,body:%s", err, res)
	}

	if !executionsRes.Success() {
		return nil, executionsRes.Error()
	}

	return executionsRes, nil
}
