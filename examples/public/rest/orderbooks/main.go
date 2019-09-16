package main

import (
	"api_client/api/common/configuration"
	"api_client/api/public/rest"
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
