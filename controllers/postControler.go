package controllers

import (
	"example/Go/initializers"
	"example/Go/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)
func Greeting(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message":"hurray"})

}

func CreatePost(c *gin.Context) {
	var body struct {
		Body string
		Title string

	}
	c.Bind(&body)
	post := models.Post{Title: body.Title, Body: body.Body}
	newPost:= initializers.DB.Create(&post)

	if newPost.Error != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "error creating post"})
		log.Fatal(newPost.Error)
		return
	}
	c.JSON(201, gin.H{"post": newPost})

}

func ReadPost(c *gin.Context) {
	id:= c.Param("id")
	var post models.Post
	 initializers.DB.First(&post, id)
	
	
	c.JSON(200, gin.H{"post": post})



}
func DeletePost(c *gin.Context) {
	id:= c.Param("id")

	var post models.Post
	 initializers.DB.Delete(&post, id)
	 c.JSON(200, gin.H{"message": "Post deleted successfully"})
}

func GetPosts(c *gin.Context) {
	var posts []models.Post
	initializers.DB.Find(&posts)
	
	c.JSON(200, gin.H{"post": posts})
}

func EditPost(c *gin.Context) {
	id:= c.Param("id")

	var body struct{
		Body string
		Title string
	}
	c.Bind(body)

	var post models.Post
	 initializers.DB.First(&post, id)

	 initializers.DB.Model(&post).Updates(models.Post{Title: body.Title, Body: body.Body})

	 c.JSON(200, gin.H{"post": post})


}
