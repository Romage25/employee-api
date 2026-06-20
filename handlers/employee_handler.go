package handlers

import (
	"employee-api/models"
	"employee-api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type EmployeeHandler struct {
	service *services.EmployeeService
}

func NewEmployeeHandler(
	service *services.EmployeeService,
) *EmployeeHandler {

	return &EmployeeHandler{
		service: service,
	}
}

func (h *EmployeeHandler) Create(c *gin.Context) {

	var employee models.Employee

	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	h.service.Create(&employee)

	c.JSON(http.StatusCreated, employee)
}

func (h *EmployeeHandler) GetAll(c *gin.Context) {

	employees, err := h.service.GetAll()

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, employees)
}

func (h *EmployeeHandler) GetByID(c *gin.Context) {

	id := c.Param("id")

	employee, err := h.service.GetByID(id)

	if err != nil {
		c.JSON(404, gin.H{
			"error": "Employee not found",
		})
		return
	}

	c.JSON(200, employee)
}

func (h *EmployeeHandler) Update(c *gin.Context) {

	id := c.Param("id")

	var data map[string]interface{}

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	employee, err := h.service.Update(
		id,
		data,
	)

	if err != nil {
		c.JSON(404, gin.H{
			"error": "Employee not found",
		})
		return
	}

	c.JSON(200, employee)
}

func (h *EmployeeHandler) Delete(c *gin.Context) {

	id := c.Param("id")

	err := h.service.Delete(id)

	if err != nil {
		c.JSON(404, gin.H{
			"error": "Employee not found",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Deleted successfully",
	})
}
