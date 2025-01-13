package main

import (
	"log"

	"github.com/souzasbella/golang-microservices/internal/database"
	"github.com/souzasbella/golang-microservices/internal/server"
)

func main() {
	db, err := database.NewDatabaseClient()
	if err != nil {
		log.Fatalf("error creating database client: %s", err)
	}

	srv := server.NewEchoServer(db)
	if err := srv.Start(); err != nil {
		log.Fatal(err.Error())
	}
}
