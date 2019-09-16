package main

import (
	"gogmocoin/api/common/configuration"
	"gogmocoin/api/public/rest"
	"log"
)

func main() {
	client := rest.New()

	orderbooks, err := client.OrderBooks(configuration.SymbolBCHJPY)

	if err != nil {
		log.Println(err)
		return
	}

	log.Printf("[result]%+v", orderbooks)
}
