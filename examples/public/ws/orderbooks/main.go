package main

import (
	"api_client/api/common/configuration"
	"api_client/api/public/ws/channel/orderbooks"
	"log"
	"time"
)

func main() {
	client := orderbooks.New(configuration.SymbolBTCJPY)
	timeoutCnt := 0
	for {
		select {
		case v := <-client.Receive():
			log.Printf("msg:%+v", v)
		case <-time.After(180 * time.Second):
			log.Println("timeout...")
			timeoutCnt++
		}
		if timeoutCnt > 20 {
			break
		}
	}
	e := client.Unsubscribe()
	if e != nil {
		log.Println(e)
		return
	}
}
