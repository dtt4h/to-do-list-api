package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"` // done/in_progress/pending
}

var tasks = []Task{
	{ID: 1, Title: "Buy a new pc", Description: "earn money", Status: "pending"},
}

func getTasks(c *gin.Context) {
	c.JSON(http.StatusOK, tasks)
}

func createTask(c *gin.Context) {
	var newTask Task

	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	newTask.ID = len(tasks) + 1

	tasks = append(tasks, newTask)

	c.JSON(http.StatusCreated, gin.H{
		"task": newTask,
	})
}

func main() {
	r := gin.Default()

	r.GET("/tasks", getTasks)

	r.POST("/tasks", createTask)

	r.Run()
}
