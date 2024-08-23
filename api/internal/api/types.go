package api

import (
	"net/http"
	"time"
)

type GetHostFunc func() string

type MakeHeaderFunc func(apiKey, secretKey string, systemDatetime time.Time, r *http.Request, method httpMethod, path, body string)
