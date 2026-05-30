package api

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/ijufumi/gogmocoin/v2/api/common/configuration"
	"github.com/ijufumi/gogmocoin/v2/api/common/consts"
)

var httpClient = &http.Client{
	Timeout: 30 * time.Second,
}

type RestAPIBase struct {
	getHostFunc    HostFactoryFunc
	makeHeaderFunc HeaderCreationFunc
	apiKey         string
	secretKey      string
}

type httpMethod string

const (
	httpMethodPOST   = httpMethod("POST")
	httpMethodGET    = httpMethod("GET")
	httpMethodPUT    = httpMethod("PUT")
	httpMethodDELETE = httpMethod("DELETE")
)

func NewRestAPIBase() RestAPIBase {
	return RestAPIBase{
		getHostFunc: getPublicAPIHost,
		makeHeaderFunc: func(apiKey, secretKey string, systemDatetime time.Time, r *http.Request, method httpMethod, path, body string) {
		},
	}
}

func NewPrivateRestAPIBase(apiKey, secretKey string) RestAPIBase {
	return RestAPIBase{
		getHostFunc:    getPrivateAPIHost,
		makeHeaderFunc: makeAuthHeader,
		apiKey:         apiKey,
		secretKey:      secretKey,
	}
}

// Post ...
func (c *RestAPIBase) Post(ctx context.Context, body any, path string) ([]byte, error) {
	return c.sendRequest(ctx, httpMethodPOST, body, path)
}

// Put ...
func (c *RestAPIBase) Put(ctx context.Context, body any, path string) ([]byte, error) {
	return c.sendRequest(ctx, httpMethodPUT, body, path)
}

// Get ...
func (c *RestAPIBase) Get(ctx context.Context, param url.Values, path string) ([]byte, error) {
	queryString := param.Encode()
	urlString := path
	if len(queryString) != 0 {
		urlString = urlString + "?" + queryString
	}
	return c.sendRequest(ctx, httpMethodGET, nil, urlString)
}

// Delete ...
func (c *RestAPIBase) Delete(ctx context.Context, body any, path string) ([]byte, error) {
	return c.sendRequest(ctx, httpMethodDELETE, body, path)
}

func (c *RestAPIBase) sendRequest(ctx context.Context, method httpMethod, bodyData any, path string) ([]byte, error) {
	var body string
	if bodyData != nil {
		b, err := json.Marshal(bodyData)
		if err != nil {
			return nil, err
		}
		body = string(b)
	}
	req, err := http.NewRequestWithContext(ctx, string(method), c.getHostFunc()+path, strings.NewReader(body))
	if err != nil {
		return nil, err
	}
	if method == httpMethodPOST {
		c.makeHeaderFunc(c.apiKey, c.secretKey, time.Now(), req, method, path, body)
	} else {
		c.makeHeaderFunc(c.apiKey, c.secretKey, time.Now(), req, method, path, "")
	}

	if configuration.IsDebug() {
		fmt.Printf("[Request]Header:%v\n", maskSensitiveHeader(req.Header))
		fmt.Printf("[Request]Body:%v\n", body)
	}
	res, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = res.Body.Close()
	}()
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if configuration.IsDebug() {
		fmt.Printf("[Response]StatusCode:%d\n", res.StatusCode)
		fmt.Printf("[Response]Body:%v\n", string(resBody))
	}

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", res.StatusCode, string(resBody))
	}

	return resBody, nil
}

// maskSensitiveHeader returns a copy of h with credential headers redacted so
// that enabling debug output never writes the API key or request signature to
// stdout.
func maskSensitiveHeader(h http.Header) http.Header {
	masked := h.Clone()
	if masked == nil {
		return masked
	}
	for _, key := range []string{"API-KEY", "API-SIGN"} {
		if masked.Get(key) != "" {
			masked.Set(key, "[REDACTED]")
		}
	}
	return masked
}

func makeAuthHeader(apiKey, secretKey string, systemDatetime time.Time, r *http.Request, method httpMethod, path, body string) {
	timeStamp := strconv.FormatInt(systemDatetime.UnixMilli(), 10)
	r.Header.Set("API-TIMESTAMP", timeStamp)
	r.Header.Set("API-KEY", apiKey)
	r.Header.Set("API-SIGN", makeSign(secretKey, timeStamp, method, path, body))
}

func makeSign(secretKey, timeStamp string, method httpMethod, path, body string) string {
	h := hmac.New(sha256.New, []byte(secretKey))
	// Write the parts directly to avoid the reflection overhead of fmt.Fprintf.
	_, _ = io.WriteString(h, timeStamp)
	_, _ = io.WriteString(h, string(method))
	_, _ = io.WriteString(h, path)
	_, _ = io.WriteString(h, body)
	return hex.EncodeToString(h.Sum(nil))
}

func getPublicAPIHost() string {
	return consts.PublicRestAPIHost
}

func getPrivateAPIHost() string {
	return consts.PrivateRestAPIHost
}
