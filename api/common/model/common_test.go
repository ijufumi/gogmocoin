package model

import (
	"errors"
	"testing"

	"github.com/ijufumi/gogmocoin/v2/api/common/consts"
	"github.com/stretchr/testify/assert"
)

func TestResponseCommon_Success(t *testing.T) {
	assert.True(t, (&ResponseCommon{Status: consts.StatusOK}).Success())
	assert.False(t, (&ResponseCommon{Status: consts.Status(1)}).Success())
}

func TestResponseCommon_Error(t *testing.T) {
	r := &ResponseCommon{
		Status:   consts.Status(5),
		Messages: []map[string]string{{"message_code": "ERR-5201", "message_string": "invalid request"}},
	}

	err := r.Error()
	assert.Error(t, err)

	// The error should be recoverable as *APIError with the status and messages.
	var apiErr *APIError
	assert.True(t, errors.As(err, &apiErr))
	assert.Equal(t, consts.Status(5), apiErr.Status)
	assert.Equal(t, "ERR-5201", apiErr.Messages[0]["message_code"])
	assert.Contains(t, err.Error(), "status=5")
}
