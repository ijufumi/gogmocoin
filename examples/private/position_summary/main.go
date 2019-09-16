package main

import (
	"api_client/api/common/configuration"
	"api_client/api/private"
	"log"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	client := private.New()
	response, err := client.PositionSummary(configuration.SymbolBTCJPY)
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("result:%+v", response)
}
