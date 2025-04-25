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
		//fmt.Println("Failed to connect to the database:", err)
	}
	defer model.DB.Close()

	router.InitRouter()
}
