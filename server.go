package main

import (
	"goapi/controllers"
	"goapi/models"

	"github.com/gin-gonic/gin"
)

func main() {

	models.ConnectDatabase()

	r := gin.Default()

	r.POST("/login")
	r.POST("/register", controllers.Register)

	r.GET("/posts", controllers.FindPosts)
	r.POST("/posts", controllers.CreatePost)
	r.GET("/posts/:id", controllers.FindPost)
	r.PATCH("/posts/:id", controllers.UpdatePost)
	r.DELETE("/posts/:id", controllers.DeletePost)

	r.Run()
}

// go mod init goapi
// go mod tidy
// go run server.go
