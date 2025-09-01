package main

import (
	"to-do-list-api/database"
	"to-do-list-api/handlers"
	"to-do-list-api/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	// Инициализация базы данных
	database.InitDB()
	defer database.CloseDB()

	r := gin.Default()

	// Публичные маршруты (не требуют авторизации)
	r.POST("/login", handlers.Login)

	// Группа защищенных маршрутов с middleware
	api := r.Group("/")
	api.Use(middleware.AuthMiddleware()) // Middleware применяется ко ВСЕМ маршрутам в группе
	{
		api.GET("/tasks", handlers.GetTasks)
		api.GET("/tasks/:id", handlers.GetTaskByID)
		api.POST("/tasks", handlers.CreateTask)
		api.PATCH("/tasks/:id", handlers.UpdateTaskByID)
		api.DELETE("/tasks/:id", handlers.DeleteTaskByID)
	}

	r.Run(":8080")
}
