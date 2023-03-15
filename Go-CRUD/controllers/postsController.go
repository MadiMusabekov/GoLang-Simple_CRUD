package controllers

import (
	"Go-CRUD/initializers"
	"Go-CRUD/models"

	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {
	//Get data off req body
	var body struct {
		Body  string
		Title string
	}
	c.Bind(&body)
	//Create a post
	post := models.Post{Title: body.Title, Body: body.Body}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}
	//Return post
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostIndex(c *gin.Context) {
	// Get the posts
	var posts []models.Post
	initializers.DB.Find(&posts)
	//Respond with them
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostsShow(c *gin.Context) {
	//Get id off url
	id := c.Param("id")
	//Get the posts
	var posts models.Post
	initializers.DB.First(&posts, id)

	//Respond with them
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostsUpdate(c *gin.Context) {
	// get the id off the url
	id := c.Param("id")
	//get the data off req body
	var body struct {
		Body  string
		Title string
	}
	c.Bind(&body)
	// Find the post were updating
	var post models.Post
	initializers.DB.First(&post, id)
	//Update it
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})
	//Respond with it
	c.JSON(200, gin.H{
		"posts": post,
	})
}

func PostsDelete(c *gin.Context) {
	// Get the id off the url
	id := c.Param("id")
	// Delete the posts
	initializers.DB.Delete(&models.Post{}, id)
	// Respond
	c.Status(200)
}
