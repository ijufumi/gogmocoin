package api

import (
	"encoding/json"
	"log"
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
			log.Printf("[%v] received:%v", name, string(v))
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

func EncodeJSON(v any) string {
	b, e := json.MarshalIndent(v, "", "")
	if e != nil {
		return ""
	}
	return string(b)
}
