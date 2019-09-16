package main

import (
	"gogmocoin/api/public/rest"
	"log"
)

func main() {
	client := rest.New()
	status, err := client.Status()
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("result:%+v", status)
}
