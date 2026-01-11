package database

import (
	"fmt"
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/minh20051202/ticket-system-backend/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	database = os.Getenv("DB_DATABASE")
	password = os.Getenv("DB_PASSWORD")
	username = os.Getenv("DB_USERNAME")
	port     = os.Getenv("DB_PORT")
	host     = os.Getenv("DB_HOST")
)

func New() {
	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Ho_Chi_Minh", host, username, password, database, port)

	db, err := gorm.Open(postgres.Open(connStr))
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	fmt.Println(">> Connected to PostgreSQL successfully")
	err = db.AutoMigrate(&models.User{}, &models.Event{}, &models.Order{}, &models.Ticket{})

	if err != nil {
		log.Fatal("Error occurred trying to create schemas:", err)
	}
	fmt.Println(">> Created schemas")
}
