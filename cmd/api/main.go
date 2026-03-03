package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	_ "github.com/joho/godotenv/autoload"
	"github.com/minh20051202/ticket-system-backend/internal/catalog"
	"github.com/minh20051202/ticket-system-backend/internal/db"
	"github.com/minh20051202/ticket-system-backend/internal/identity"
	"github.com/minh20051202/ticket-system-backend/internal/ledger"
	"github.com/minh20051202/ticket-system-backend/internal/proxy"
	"github.com/minh20051202/ticket-system-backend/internal/proxy/adapters"
	"github.com/minh20051202/ticket-system-backend/internal/utils"
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

	catalogRepo := catalog.NewPostgresRepository(db)
	if err := catalogRepo.Init(); err != nil {
		log.Fatal("Failed to init catalog tables:", err)
	}
	catalogService := catalog.NewService(catalogRepo)
	catalogHandler := catalog.NewHandler(catalogService)

	proxyService := proxy.NewService(catalogService, ledgerService)
	proxyService.RegisterAdapter("openai", &adapters.OpenAIAdapter{})
	proxyHandler := proxy.NewHandler(proxyService)
	router := mux.NewRouter()

	identityHandler.RegisterRoutes(router)
	ledgerHandler.RegisterRoutes(router)
	catalogHandler.RegisterRoutes(router)
	router.HandleFunc("/v1/run",
		identity.WithApiKeyAuth(identityService)(
			utils.MakeHTTPHandleFunc(proxyHandler.HandleRun),
		),
	).Methods("POST")

	log.Println("Server running on :8080")
	http.ListenAndServe(":8080", router)
}
