package main

import (
	"example/Go/initializers"
	"example/Go/models"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDB()
}
func main() {
	initializers.DB.AutoMigrate(&models.Post{})
	initializers.DB.AutoMigrate(&models.User{})
	
}