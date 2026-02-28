package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	_ "github.com/joho/godotenv/autoload"
	"github.com/minh20051202/ticket-system-backend/internal/db"
	"github.com/minh20051202/ticket-system-backend/internal/identity"
	"github.com/minh20051202/ticket-system-backend/internal/ledger"
)

var (
	database = os.Getenv("DB_DATABASE")
	password = os.Getenv("DB_PASSWORD")
	username = os.Getenv("DB_USERNAME")
	port     = os.Getenv("DB_PORT")
	host     = os.Getenv("DB_HOST")
)

func main() {
	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Ho_Chi_Minh", host, username, password, database, port)

	db, err := db.NewConnection(connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	identityRepo := identity.NewPostgresRepository(db)
	if err := identityRepo.Init(); err != nil {
		log.Fatal("Failed to init identity tables:", err)
	}
	identityService := identity.NewService(identityRepo)
	identityHandler := identity.NewHandler(identityService)

	ledgerRepo := ledger.NewPostgresRepository(db)
	if err := ledgerRepo.Init(); err != nil {
		log.Fatal("Failed to init ledger tables:", err)
	}
	ledgerService := ledger.NewService(ledgerRepo)
	ledgerHandler := ledger.NewHandler(ledgerService)

	router := mux.NewRouter()

	identityHandler.RegisterRoutes(router)
	ledgerHandler.RegisterRoutes(router)

	log.Println("Server running on :8080")
	http.ListenAndServe(":8080", router)
}
