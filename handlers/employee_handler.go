package handlers

import (
	"employee-api/database"
	"employee-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetEmployees(c *gin.Context) {
	var employees []models.Employee

	database.DB.Find(&employees)

	c.JSON(http.StatusOK, employees)
}

func CreateEmployee(c *gin.Context) {
	var employee models.Employee

	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	database.DB.Create(&employee)

	c.JSON(http.StatusCreated, employee)
}

func UpdateEmployee(c *gin.Context) {
	var employee models.Employee

	id := c.Param("id")

	if err := database.DB.First(&employee, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Employee not found",
		})
		return
	}

	var input map[string]interface{}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	database.DB.Model(&employee).Updates(input)

	c.JSON(http.StatusOK, employee)
}

func DeleteEmployee(c *gin.Context) {
	var employee models.Employee

	id := c.Param("id")

	if err := database.DB.First(&employee, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Employee not found",
		})
		return
	}

	database.DB.Delete(&employee)

	c.JSON(http.StatusOK, gin.H{
		"message": "Employee deleted successfully",
	})
}

func GetEmployee(c *gin.Context) {
	var employee models.Employee

	id := c.Param("id")

	if err := database.DB.First(&employee, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Employee not found",
		})
		return
	}

	c.JSON(http.StatusOK, employee)
}