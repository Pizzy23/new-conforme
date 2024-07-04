package main

import (
	"conforme/db"
	"conforme/middleware"
	"log"
)

// @title           ConformeTec
// @version         1.0
// @description     This is a server for app.

// @host      3.138.100.192:8080

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/

func main() {
	log.Println("Starting application...")

	r := middleware.SetupRouter()

	log.Println("Connecting to database...")
	db.ConnectDatabase()

	log.Println("Running migrations...")
	if err := db.Migrate(db.Repo); err != nil {
		log.Fatal("Failed to migrate database: ", err)
	}

	log.Println("Starting server...")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
