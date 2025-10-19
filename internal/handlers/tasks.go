package handlers

import (
	"database/sql"
	"net/http"
	"strconv"
	"to-do-list-api/internal/database"
	"to-do-list-api/internal/models"

	"github.com/gin-gonic/gin"
)

func GetTaskByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid ID format | ID must be a number",
		})
		return
	}

	var task models.Task
	err = database.DB.QueryRow("SELECT id, title, description, status FROM tasks WHERE id = ?", id).
		Scan(&task.ID, &task.Title, &task.Description, &task.Status)

	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch task"})
		}
		return
	}

	c.JSON(http.StatusOK, task)
}

func GetTasks(c *gin.Context) {
	rows, err := database.DB.Query("SELECT id, title, description, status FROM tasks")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch tasks"})
		return
	}
	defer rows.Close()

	var tasks []models.Task
	for rows.Next() {
		var task models.Task
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Status)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan task"})
			return
		}
		tasks = append(tasks, task)
	}

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

	if newTask.Status == "" {
		newTask.Status = "pending"
	}

	result, err := database.DB.Exec("INSERT INTO tasks (title, description, status) VALUES (?, ?, ?)",
		newTask.Title, newTask.Description, newTask.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task"})
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get task ID"})
		return
	}

	newTask.ID = int(id)

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

	var task models.Task
	err = database.DB.QueryRow("SELECT id, title, description, status FROM tasks WHERE id = ?", id).
		Scan(&task.ID, &task.Title, &task.Description, &task.Status)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	if statusValue, exists := updateData["status"]; exists {
		s, ok := statusValue.(string)
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{"error": "status must be a string"})
			return
		}
		task.Status = s
	}
	if titleValue, exists := updateData["title"]; exists {
		s, ok := titleValue.(string)
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{"error": "title must be a string"})
			return
		}
		task.Title = s
	}
	if descValue, exists := updateData["description"]; exists {
		s, ok := descValue.(string)
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{"error": "description must be a string"})
			return
		}
		task.Description = s
	}

	_, err = database.DB.Exec("UPDATE tasks SET title = ?, description = ?, status = ? WHERE id = ?", task.Title, task.Description, task.Status, task.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update task"})
		return
	}

	c.JSON(http.StatusOK, task)
}

func DeleteTaskByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	result, err := database.DB.Exec("DELETE FROM tasks WHERE id = ?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete task"})
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get deletion result"})
		return
	}

	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
}
