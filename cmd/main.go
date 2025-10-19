package main

import (
	"fmt"
	"os"
	"to-do-list-api/internal/database"
	"to-do-list-api/internal/handlers"
	"to-do-list-api/internal/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	// Инициализация базы данных
	database.InitDB()
	defer database.CloseDB()

	r := gin.Default()

	// Health check
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

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

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if err := r.Run(":" + port); err != nil {
		// log to stderr so container runtimes capture it
		fmt.Fprintf(os.Stderr, "failed to start server: %v\n", err)
	}
}
