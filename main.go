package main

import (
	"api/db"
	"api/router"
	"fmt"
)

func main() {

	db.InitPostgresDB()
	err := router.InitRouter().Run()
	if err != nil {
		fmt.Print("Error starting the server")
		return
	}

}
