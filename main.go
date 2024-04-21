package main

import (
	"api2/config"
	"api2/models"
	"api2/server"
	"log"
)

func Init() {
	cfg, err := config.ReadConfig()
	if err != nil {
		log.Panicln(err)
	}
	config.InitLog(cfg)
	models.InitDB(cfg)
}

func main() {
	Init()
	api := server.InitRouter()

	err := api.Run(":8888")
	if err != nil {
		log.Panicln(err)
	}
}
