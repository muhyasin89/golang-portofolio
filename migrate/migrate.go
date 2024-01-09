package main

import (
	"go-crud/config"
	"go-crud/models"
)

func init() {
	config.LoadEnv()
	config.ConnectToDB()
}

func main() {
	// Migrate the schema
	config.DB.AutoMigrate(&models.User{})
	config.DB.AutoMigrate(&models.Post{})
}
