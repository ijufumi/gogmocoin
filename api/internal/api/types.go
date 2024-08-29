package api

import (
	"net/http"
	"time"
)

type HostFactoryFunc func() string

type HeaderCreationFunc func(apiKey, secretKey string, systemDatetime time.Time, r *http.Request, method httpMethod, path, body string)
