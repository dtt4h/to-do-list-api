package handlers

import (
	"net/http"
	"to-do-list-api/models"

	"github.com/gin-gonic/gin"
)

var tasks = []models.Task{
	{ID: 1, Title: "Buy a new pc", Description: "earn money", Status: "pending"},
}

func GetTaskByID(c *gin.Context) {
	// TODO: реализовать поиск задачи по ID
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"task_id": id,
	})
}

func GetTasks(c *gin.Context) {
	c.JSON(http.StatusOK, tasks)
}

func CreateTask(c *gin.Context) {
	var newTask models.Task

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
