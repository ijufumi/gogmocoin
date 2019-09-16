package main

import (
	"gogmocoin/api/common/configuration"
	"gogmocoin/api/private"
	"log"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	client := private.New()
	executionsRes, err := client.LastExecutions(configuration.SymbolBTCJPY, 0, 0)
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("result:%+v", executionsRes)
}
