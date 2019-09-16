package main

import (
	"gogmocoin/api/common/configuration"
	"gogmocoin/api/public/rest"
	"log"
	"time"
)

func main() {
	client := rest.New()

	for i := 0; i < 5; i++ {
		tradesRes, err := client.Trades(configuration.SymbolXRPJPY, int64(i), int64(i))
		if err != nil {
			log.Println(err)
			continue
		}
		log.Printf("[%v]result:%+v", i, tradesRes)
		time.Sleep(time.Second)
	}
}
