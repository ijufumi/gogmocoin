package api

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/ijufumi/gogmocoin/v2/api/common/configuration"
	"github.com/ijufumi/gogmocoin/v2/api/common/consts"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type RestAPIBase struct {
	needsAuth bool
	apiKey    string
	secretKey string
}

func NewRestAPIBase() RestAPIBase {
	return RestAPIBase{}
}

func NewPrivateRestAPIBase(apiKey, secretKey string) RestAPIBase {
	return RestAPIBase{
		needsAuth: true,
		apiKey:    apiKey,
		secretKey: secretKey,
	}
}

// Post ...
func (c *RestAPIBase) Post(body interface{}, path string) ([]byte, error) {
	return c.sendRequest("POST", body, path)
}

// Put ...
func (c *RestAPIBase) Put(body interface{}, path string) ([]byte, error) {
	return c.sendRequest("PUT", body, path)
}

// Get ...
func (c *RestAPIBase) Get(param url.Values, path string) ([]byte, error) {
	queryString := param.Encode()
	urlString := path
	if len(queryString) != 0 {
		urlString = urlString + "?" + queryString
	}
	return c.sendRequest("GET", nil, urlString)
}

// Delete ...
func (c *RestAPIBase) Delete(body interface{}, path string) ([]byte, error) {
	return c.sendRequest("DELETE", body, path)
}

func (c *RestAPIBase) sendRequest(method string, bodyData interface{}, path string) ([]byte, error) {
	var body string
	if bodyData != nil {
		b, err := json.Marshal(bodyData)
		if err != nil {
			return nil, err
		}
		body = string(b)
	}
	req, err := http.NewRequest(method, c.getHost()+path, strings.NewReader(body))
	if err != nil {
		return nil, err
	}
	if c.needsAuth {
		c.makeHeader(time.Now(), req, method, path, body)
	}

	if configuration.Debug {
		fmt.Printf("[Request]Header:%v\n", req.Header)
		fmt.Printf("[Request]Body:%v\n", body)
	}
	res, err := http.DefaultClient.Do(req)
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

	if configuration.Debug {
		fmt.Printf("[Response]Body:%v\n", string(resBody))
	}

	return resBody, nil
}

func (c *RestAPIBase) makeHeader(systemDatetime time.Time, r *http.Request, method, path, body string) {
	timeStamp := systemDatetime.Unix()*1000 + int64(systemDatetime.Nanosecond())/int64(time.Millisecond)
	r.Header.Set("API-TIMESTAMP", fmt.Sprint(timeStamp))
	r.Header.Set("API-KEY", c.apiKey)
	r.Header.Set("API-SIGN", c.makeSign(c.secretKey, timeStamp, method, path, body))
}

func (c *RestAPIBase) makeSign(secretKey string, timeStamp int64, method, path, body string) string {
	h := hmac.New(sha256.New, []byte(secretKey))
	h.Write([]byte(fmt.Sprintf("%v%v%v%v", timeStamp, method, path, body)))
	return hex.EncodeToString(h.Sum(nil))
}

func (c *RestAPIBase) getHost() string {
	if c.needsAuth {
		return consts.PrivateRestAPIHost
	}
	return consts.PublicRestAPIHost
}
