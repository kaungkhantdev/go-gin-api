package database

import (
	"fmt"
	"os"

	"go-gin-api/internal/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	_ "github.com/joho/godotenv/autoload"
)


var (
	dbname     = os.Getenv("DB_DATABASE")
	password   = os.Getenv("DB_PASSWORD")
	username   = os.Getenv("DB_USERNAME")
	port       = os.Getenv("DB_PORT")
	host       = os.Getenv("DB_HOST")
	DB *gorm.DB
)


func Connect() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, dbname);

	var err error
    DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to the database:", err)
    }

    DB.AutoMigrate(&models.Test{})
}