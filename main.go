package main

import (
	"employee-api/database"
	"employee-api/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()

	r := gin.Default()

	r.GET("/employees", handlers.GetEmployees)
	r.POST("/employees", handlers.CreateEmployee)
	r.GET("/employees/:id", handlers.GetEmployee)
	r.PATCH("/employees/:id", handlers.UpdateEmployee)
	r.DELETE("/employees/:id", handlers.DeleteEmployee)

	r.Run(":8080")
}