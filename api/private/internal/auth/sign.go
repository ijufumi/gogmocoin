package auth

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

// MakeSign ...
func MakeSign(secretKey string, timeStamp int64, method, path, body string) string {
	h := hmac.New(sha256.New, []byte(secretKey))
	h.Write([]byte(fmt.Sprintf("%v%v%v%v", timeStamp, method, path, body)))
	return hex.EncodeToString(h.Sum(nil))
}
