package handlers

import (
	"net/http"
	"strconv"
	"to-do-list-api/models"

	"github.com/gin-gonic/gin"
)

var tasks = []models.Task{
	{ID: 1, Title: "Buy a new pc", Description: "earn money", Status: "pending"},
}

func GetTaskByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid ID format | ID must be a number",
		})
	}

	for _, task := range tasks {
		if task.ID == id {
			c.JSON(http.StatusOK, task)
			return
		}
	}
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

func UpdateTaskByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	var updateData map[string]interface{}
	if err := c.BindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	for i := range tasks {
		if tasks[i].ID == id {
			if status, exists := updateData["status"]; exists {
				tasks[i].Status = status.(string) // приведение типа
			}
			c.JSON(http.StatusOK, tasks[i])
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
}

func DeleteTaskByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	for i := range tasks {
		if tasks[i].ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
}
