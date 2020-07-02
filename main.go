package main

import (
	"log"

	"github.com/Team-73/backend/data"
	"github.com/Team-73/backend/server"
	"github.com/Team-73/backend/service"
)

func main() {
	log.Println("Reading the intial configs...")

	db, err := data.Connect()
	if err != nil {
		panic(err)
	}
	svc := service.New(db)
	server := server.InitServer(svc)
	log.Println("About to start the application...")

	if err := server.Run(":5000"); err != nil {
		panic(err)
	}
}
