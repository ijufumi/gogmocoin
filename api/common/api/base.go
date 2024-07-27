package api

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/ijufumi/gogmocoin/api/common/configuration"
	"github.com/ijufumi/gogmocoin/api/common/consts"
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
	b, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", c.getHost()+path, strings.NewReader(string(b)))
	if err != nil {
		return nil, err
	}
	c.makeHeader(time.Now(), req, "POST", path, string(b))

	if configuration.Debug {
		fmt.Printf("[Request]Header:%v\n", req.Header)
		fmt.Printf("[Request]Body:%v\n", string(b))
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

// Get ...
func (c *RestAPIBase) Get(param url.Values, path string) ([]byte, error) {
	queryString := param.Encode()
	urlString := c.getHost() + path
	if len(queryString) != 0 {
		urlString = urlString + "?" + queryString
	}
	req, err := http.NewRequest("GET", urlString, nil)
	if err != nil {
		return nil, err
	}
	c.makeHeader(time.Now(), req, "GET", path, "")

	if configuration.Debug {
		fmt.Printf("[Request]Header:%v\n", req.Header)
		fmt.Printf("[Request]URL:%v\n", req.URL)
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
