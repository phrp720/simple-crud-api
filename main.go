package main

import (
	"api/db"
	"api/router"
	"log"
)

func main() {

	db.InitPostgresDB()
	err := router.InitRouter().Run()
	if err != nil {
		log.Print("Error starting the server")
		return
	}

}
