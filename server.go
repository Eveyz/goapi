package main

import (
	"github.com/gin-gonic/gin"
	"znz.com/web/controllers"
	"znz.com/web/models"
)

func main() {

	models.ConnectDatabase()

	r := gin.Default()

	r.POST("/login")
	r.POST("/register", controllers.Register)

	r.GET("/books", controllers.FindBooks)
	r.POST("/books", controllers.CreateBook)
	r.GET("/books/:id", controllers.FindBook)
	r.PATCH("/books/:id", controllers.UpdateBook)
	r.DELETE("/books/:id", controllers.DeleteBook)

	r.GET("/students", controllers.FindStudents)
	r.POST("/students", controllers.CreateStudent)
	r.GET("/students/:id", controllers.FindStudent)
	r.PATCH("/students/:id", controllers.UpdateStudent)
	r.DELETE("/students/:id", controllers.DeleteStudent)

	r.Run()
}