package main

import (
	"example/Go/controllers"
	"example/Go/initializers"
	"example/Go/middlewares"
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
	 router.PATCH("/post/edit/:id", middlewares.ProtectRoute,middlewares.RestrictedRoute("Admin"),controllers.EditPost)
	 router.DELETE("/post/:id", middlewares.ProtectRoute,middlewares.RestrictedRoute("Admin"),controllers.DeletePost)
	 router.POST("/signup", controllers.CreateUser)
	 router.POST("/login", controllers.LoginUser)
	 router.Run()
	 
}