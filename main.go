package main

import (
	"employee-api/database"
	"employee-api/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()

	r := gin.Default()

	r.POST("/employees", handlers.CreateEmployee)

	r.Run(":8080")
}