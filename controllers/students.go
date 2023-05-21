package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"znz.com/web/models"
)

// GET /students
// Get all students
func FindStudents(c *gin.Context) {
	var students []models.Student
	models.DB.Find(&students)

	c.JSON(http.StatusOK, gin.H{"students": students})
}

// POST /Students
// Create new Student
func CreateStudent(c *gin.Context) {
	// Validate input
	var input models.CreateStudentInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create Student
	Student := models.Student{Name: input.Name, Age: input.Age, Gender: input.Gender}
	models.DB.Create(&Student)

	c.JSON(http.StatusOK, gin.H{"student": Student})
}

// GET /students/:id
// Find a student
func FindStudent(c *gin.Context) { // Get model if exist
	var student models.Student

	if err := models.DB.Where("id = ?", c.Param("id")).First(&student).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"student": student})
}

// PATCH /students/:id
// Update a student
func UpdateStudent(c *gin.Context) {
	// Get model if exist
	var student models.Student
	if err := models.DB.Where("id = ?", c.Param("id")).First(&student).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input models.UpdateStudentInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&student).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": student})
}

// DELETE /students/:id
// Delete a student
func DeleteStudent(c *gin.Context) {
	// Get model if exist
	var student models.Student
	if err := models.DB.Where("id = ?", c.Param("id")).First(&student).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&student)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
