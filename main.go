package main

import (
	"api2/models"
	"api2/server"
	"log"
)

func Init() {
	models.InitDB()
}

func main() {
	Init()
	api := server.InitRouter()

	err := api.Run(":8888")
	if err != nil {
		log.Panicln(err)
	}
}
