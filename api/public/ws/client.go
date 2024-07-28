package ws

import (
	"encoding/json"
	"github.com/ijufumi/gogmocoin/api/common/consts"
	"log"
)

func NewOrderBooks(symbol consts.Symbol) OrderBooks {
	return newOrderBooks(symbol)
}

func NewTicker(symbol consts.Symbol) Ticker {
	return newTicker(symbol)
}

func NewTrades(symbol consts.Symbol, option *consts.Option) Trades {
	return newTrades(symbol, option)
}

func retrieveStream[T any](name string, rawStream <-chan []byte) <-chan *T {
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
