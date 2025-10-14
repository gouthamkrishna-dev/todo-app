package main

import (
	"application/todo/tododatabase"

	"application/todo/route"

	"github.com/gin-gonic/gin"
)

// Create a database

func main() {
	tododatabase.Createdatabase()
	defer tododatabase.DB.Close()
	router := gin.Default()
	router.GET("/list", route.GetTodo)
	router.POST("/list", route.AddTodo)
	router.Run("localhost:8080")
}
