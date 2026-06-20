package main

import (
	"employee-api/database"
)

func main() {
	database.Connect()

	println("Database connected successfully!")
}