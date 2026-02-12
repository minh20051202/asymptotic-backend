package main

import (
	"log"

	"github.com/minh20051202/ticket-system-backend/internal/database"
	"github.com/minh20051202/ticket-system-backend/internal/server"
)

func main() {
	db, err := database.New()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Init(); err != nil {
		log.Fatal(err)
	}

	server := server.NewAPIServer(":8080", db)
	server.Run()
}
