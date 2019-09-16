package main

import (
	"api_client/api/private"
	"log"
)

func main() {
	client := private.NewWithKeys("YOUR_API_KEY", "YOUR_API_SECRET")
	ordersRes, err := client.Orders(12345676879)
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("ordersRes:%+v", ordersRes)
}
