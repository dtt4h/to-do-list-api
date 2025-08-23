package main

import (
	"to-do-list-api/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/tasks", handlers.GetTasks)

	r.POST("/tasks", handlers.CreateTask)

	r.GET("/tasks/:id", handlers.GetTaskByID)

	r.Run()
}
