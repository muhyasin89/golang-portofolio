package config

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error

	db_host := os.Getenv("POSTGRES_READ_HOST")
	db_user := os.Getenv("POSTGRES_READ_USER")
	db_password := os.Getenv("POSTGRES_READ_PASSWORD")
	db_name := os.Getenv("POSTGRES_READ_DB")
	db_port := os.Getenv("POSTGRES_READ_PORT")

	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", db_host, db_user, db_password, db_name, db_port)
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect Database")
	}

	DB = database
}
