package main

import (
	"api/db"
	"api/router"
	"log"
)

// @title           Simple RESTful CRUD API
// @version         1.0

// @contact.name   Phillip Rafail Papadakis
// @contact.email  filippospapadakis1@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

func main() {

	db.InitPostgresDB()
	err := router.InitRouter().Run()
	if err != nil {
		log.Print("Error starting the server")
		return
	}

}
