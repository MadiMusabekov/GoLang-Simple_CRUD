package main

import (
	"Go-CRUD/initializers"
	"Go-CRUD/models"
)

func init() {
	initializers.ConnectToDB()
	initializers.LoadEnvVariables()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
