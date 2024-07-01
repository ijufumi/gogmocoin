package connect

import (
	"log"
	"net/http"
	"time"

	"github.com/ijufumi/gogmocoin/api/private"
)

const (
	TermAccessToken = 59 * time.Minute
)

type AccessToken struct {
	token     string
	createdAt time.Time
}

func reqAccessToken() *AccessToken {
	return &AccessToken{}
}

func (a *AccessToken) Token() string {
	return a.token
}

func (a *AccessToken) CreatedAt() string {
	return a.token
}

// 60分経過した場合は、トークンを取得し直す
func (a *AccessToken) IsExpired() bool {
	if a.token == "" {
		return true
	}

	// 60分経過していない場合は、トークンを取得し直す
	if time.Since(a.createdAt) > TermAccessToken {
		return true
	}

	return false
}

func (a *AccessToken) Get() *AccessToken {
	client := private.New()
	token, err := client.AccessToken(http.MethodPost)
	if err != nil {
		log.Printf("failed to get access token: %v", err)
		return nil
	}
	a.token = token
	a.createdAt = time.Now()

	log.Printf("[success] access token: %v, expiration to %s", a.token, a.createdAt.Add(60*time.Minute).String())

	return a
}

func (a *AccessToken) Extend() *AccessToken {
	client := private.New()
	token, err := client.AccessToken(http.MethodPut, a.token)
	if err != nil {
		log.Printf("failed to get access token: %v", err)
		return nil
	}

	if token != a.token {
		log.Printf("token is not matched, failed extend token: %v", token)
		return nil
	}

	a.createdAt = time.Now()

	return a
}

func (a *AccessToken) Delete() *AccessToken {
	client := private.New()
	token, err := client.AccessToken(http.MethodDelete, a.token)
	if err != nil {
		log.Printf("failed to get access token: %v", err)
		return nil
	}

	if token != a.token {
		log.Printf("token is not matched, failed delete token: %v", token)
		return nil
	}

	a.token = ""
	a.createdAt = time.Time{}

	return a
}
