package models

import "gorm.io/gorm"

type Employee struct {
	gorm.Model

	Name     string `json:"name"`
	Email    string `json:"email"`
	Position string `json:"position"`
}