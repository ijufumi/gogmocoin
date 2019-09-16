package main

import (
	"gogmocoin/api/common/configuration"
	"gogmocoin/api/private"
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
