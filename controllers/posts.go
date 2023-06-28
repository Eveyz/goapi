package controllers

import (
	"net/http"

	"goapi/models"

	"github.com/gin-gonic/gin"
)

// GET /posts
// Get all posts
func FindPosts(c *gin.Context) {
	var posts []models.Book
	models.DB.Find(&posts)

	c.JSON(http.StatusOK, gin.H{"posts": posts})
}

// POST /posts
// Create new post
func CreatePost(c *gin.Context) {
	// Validate input
	var input models.CreatePostInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create book
	post := models.Post{Title: input.Title, Content: input.Content}
	models.DB.Create(&post)

	c.JSON(http.StatusOK, gin.H{"post": post})
}

// GET /posts/:id
// Find a post
func FindPost(c *gin.Context) { // Get model if exist
	var post models.Post

	if err := models.DB.Where("id = ?", c.Param("id")).First(&post).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"post": post})
}

// PATCH /posts/:id
// Update a post
func UpdatePost(c *gin.Context) {
	// Get model if exist
	var post models.Post
	if err := models.DB.Where("id = ?", c.Param("id")).First(&post).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input models.UpdatePostInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&post).Updates(input)

	c.JSON(http.StatusOK, gin.H{"post": post})
}

// DELETE /posts/:id
// Delete a post
func DeletePost(c *gin.Context) {
	// Get model if exist
	var post models.Post
	if err := models.DB.Where("id = ?", c.Param("id")).First(&post).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&post)

	c.JSON(http.StatusOK, gin.H{"success": true})
}
