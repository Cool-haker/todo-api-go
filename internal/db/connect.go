package db

import (
	"fmt"
	"log"
	"os"

	"github.com/Cool-haker/todo-api-go/internal/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	_ = godotenv.Load("internal/db/.env")

	host := os.Getenv("POSTGRES_DBHOST")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DBNAME")
	port := os.Getenv("POSTGRES_DBPORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Cannot connect to database")
	}
	DB = db

	if err := DB.AutoMigrate(&models.Todo{}, &models.User{}); err != nil {
		log.Fatalf("Failed to migrate models: %v", err)
	}

	fmt.Println("Connection to database successful")
}
