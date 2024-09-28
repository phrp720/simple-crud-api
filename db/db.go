package db

import (
	"api/dao"
	"fmt"
	"gorm.io/driver/sqlite"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var GetDatabase *gorm.DB
var err error

func InitPostgresDB() {
	err = godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
	var (
		host     = os.Getenv("DB_HOST")
		port     = os.Getenv("DB_PORT")
		dbUser   = os.Getenv("DB_USERNAME")
		dbName   = os.Getenv("DB_NAME")
		password = os.Getenv("DB_PASSWORD")
	)
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		host,
		port,
		dbUser,
		dbName,
		password,
	)

	GetDatabase, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	err := GetDatabase.AutoMigrate(&dao.Product{})
	if err != nil {
		log.Fatal(err)
	}
}

func InitInMemoryDB() {
	var err error
	dialector := sqlite.New(sqlite.Config{
		DSN: ":memory:", // Data Source Name for in-memory database
		// Add other configurations if needed
	})
	GetDatabase, err = gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	err = GetDatabase.AutoMigrate(&dao.Product{})
	if err != nil {
		log.Fatal(err)
	}
}
