package main

import (
	"example/Go/controllers"
	"example/Go/initializers"
	"fmt"

	"github.com/gin-gonic/gin"
)
func init() {
	initializers.LoadEnv()
	initializers.ConnectToDB()
}

func main() {
	fmt.Println("Starting")
	 router := gin.Default()
	 router.GET("/", controllers.Greeting)
	 router.GET("/posts", controllers.GetPosts)
	 router.POST("/post", controllers.CreatePost)
	 router.GET("/posts/:id", controllers.ReadPost)
	 router.PATCH("/post/:id", controllers.EditPost)
	 router.DELETE("/post/:id", controllers.DeletePost)
	 router.Run()
	 
}