package connect

import (
	"api_client/api/common/configuration"
	"api_client/api/private/internal/auth"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const host = "https://api.coin.z.com/private"

// Connection ...
type Connection struct {
	apiKey    string
	secretKey string
}

// New ...
func New(apiKey, secretKey string) *Connection {
	return &Connection{
		apiKey:    apiKey,
		secretKey: secretKey,
	}
}

// Post ...
func (c *Connection) Post(body interface{}, path string) ([]byte, error) {
	b, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", host+path, strings.NewReader(string(b)))
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
		res.Body.Close()
	}()
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if configuration.Debug {
		fmt.Printf("[Response]Body:%v\n", string(resBody))
	}

	return resBody, nil
}

// Get ...
func (c *Connection) Get(param url.Values, path string) ([]byte, error) {
	queryString := param.Encode()
	urlString := host + path
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
		res.Body.Close()
	}()
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if configuration.Debug {
		fmt.Printf("[Response]Body:%v\n", string(resBody))
	}
	return resBody, nil

}

func (c *Connection) makeHeader(systemDatetime time.Time, r *http.Request, method, path, body string) {
	timeStamp := systemDatetime.Unix()*1000 + int64(systemDatetime.Nanosecond())/int64(time.Millisecond)
	r.Header.Set("API-TIMESTAMP", fmt.Sprint(timeStamp))
	r.Header.Set("API-KEY", c.apiKey)
	r.Header.Set("API-SIGN", auth.MakeSign(c.secretKey, timeStamp, method, path, body))
}
