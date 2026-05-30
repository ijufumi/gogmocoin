package api

import (
	"encoding/json"
	"log"

	"github.com/ijufumi/gogmocoin/v2/api/common/configuration"
)

// RetrieveStream ...
func RetrieveStream[T any](name string, rawStream <-chan []byte) <-chan *T {
	stream := make(chan *T, 10)
	go func() {
		for {
			v := <-rawStream
			if v == nil {
				return
			}
			if configuration.IsDebug() {
				log.Printf("[%v] received:%v", name, string(v))
			}
			res := new(T)
			err := json.Unmarshal(v, res)
			if err != nil {
				log.Printf("[%v] unmarshal error:%v", name, err)
				continue
			}
			stream <- res
		}
	}()
	return stream
}
