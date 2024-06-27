package private

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ijufumi/gogmocoin/api/private/internal/connect"
)

// AccessToken ...
type AccessToken interface {
	// AccessToken required targetToken for extend/delete
	AccessToken(method string, targetToken ...string) (string, error)
}

type accessToken struct {
	con *connect.Connection
}

type RequestForAccessToken struct {
	Token string `json:"token"`
}

// AccessToken ... choose method to get as post/extend as put/revoke as delete access token
// required targetToken for extend/delete
func (c *accessToken) AccessToken(method string, targetToken ...string) (string, error) {
	req := RequestForAccessToken{}

	var (
		res []byte
		err error
	)

	switch method {
	case http.MethodPost:
		// get access token for websocket 60mins
		res, err = c.con.Post(req, "/v1/ws-auth")

	case http.MethodPut:
		// extend access token for websocket extend 60mins
		if len(targetToken) == 0 {
			return "", fmt.Errorf("extend token is not found")
		}
		req.Token = targetToken[0]
		res, err = c.con.Put(req, "/v1/ws-auth")

	case http.MethodDelete:
		// delete access token for websocket
		if len(targetToken) == 0 {
			return "", fmt.Errorf("extend token is not found")
		}
		req.Token = targetToken[0]
		res, err = c.con.Delete(req, "/v1/ws-auth")

	default:
		return "", fmt.Errorf("method:%s is not supported", method)
	}
	if err != nil {
		return "", err
	}

	result := new(map[string]interface{})
	err = json.Unmarshal(res, result)
	if err != nil {
		return "", fmt.Errorf("[order]error:%v,body:%s", err, res)
	}

	token, isThere := (*result)["data"]
	if !isThere {
		// for put/delete
		status, isThere := (*result)["status"]
		if !isThere {
			return "", fmt.Errorf("access token is not found in response")
		}

		if fmt.Sprintf("%v", status) != "0" {
			return "", fmt.Errorf("status is not success")
		}
		return targetToken[0], nil
	}

	return token.(string), nil
}
