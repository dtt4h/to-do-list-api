package main

import (
	"to-do-list-api/database"
	"to-do-list-api/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	// Инициализация базы данных
	database.InitDB()
	defer database.CloseDB()

	r := gin.Default()

	r.GET("/tasks", handlers.GetTasks)
	r.GET("/tasks/:id", handlers.GetTaskByID)

	r.POST("/tasks", handlers.CreateTask)

	r.PATCH("/tasks/:id", handlers.UpdateTaskByID)

	r.DELETE("/tasks/:id", handlers.DeleteTaskByID)
	r.Run()
}
