package main

import (
	"log"

	"github.com/sanjevscet/go-microservices/internal/database"
	"github.com/sanjevscet/go-microservices/internal/server"
)

func main() {
	db, err := database.NewDatabaseClient()
	if err != nil {
		log.Fatalf("failed to initialized Database client %s", err.Error())
	}

	myServer := server.NewEchoServer(db)

	if err := myServer.Start(); err != nil {
		log.Fatalf("unable to start server %s", err.Error())
	}

}
