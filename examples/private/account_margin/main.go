package main

import (
	"api_client/api/private"
	"log"
)

func main() {
	client := private.NewWithKeys("YOUR_API_KEY", "YOUR_API_SECRET")
	accountMarginRes, err := client.AccountMargin()
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("result:%+v", accountMarginRes)
}
