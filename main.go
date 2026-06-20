package main

import (
	"employee-api/database"
	"employee-api/handlers"
	"employee-api/repositories"
	"employee-api/services"

	"github.com/gin-gonic/gin"
)

func main() {

	database.Connect()

	repo := repositories.NewEmployeeRepository(
		database.DB,
	)

	service := services.NewEmployeeService(
		repo,
	)

	handler := handlers.NewEmployeeHandler(
		service,
	)

	r := gin.Default()

	r.POST("/employees", handler.Create)

	r.GET("/employees", handler.GetAll)

	r.GET("/employees/:id", handler.GetByID)

	r.PATCH("/employees/:id", handler.Update)

	r.DELETE("/employees/:id", handler.Delete)

	r.Run(":8080")
}
