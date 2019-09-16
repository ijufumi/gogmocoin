package model

import (
	"time"
)

// ResponseCommon ...
type ResponseCommon struct {
	Messages     []map[string]string `json:"messages,omitempty"`
	Status       int                 `json:"status"`
	ResponseTime time.Time           `json:"responsetime"`
}

// Pagination ...
type Pagination struct {
	CurrentPage int `json:"currentPage"`
	Count       int `json:"count"`
}
