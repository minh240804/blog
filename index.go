package main

import (
	"blog/model"
	"blog/router"
	//"fmt"
	"log"
)

func main() {
	err := model.DbConnect()
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	} else {
		log.Print("connect complete")
	}
	// defer model.DB.Close()
	router.InitRouter()
}
