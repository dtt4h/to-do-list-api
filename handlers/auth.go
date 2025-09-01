package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"to-do-list-api/auth"

	"github.com/gin-gonic/gin"
)

// LoginRequest структура для запроса логина
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Login обрабатывает запрос на аутентификацию
func Login(c *gin.Context) {
	var loginReq LoginRequest

	// Логируем raw body для отладки
	body, _ := c.GetRawData()
	fmt.Printf("Raw request body: %s\n", string(body))

	// Пытаемся распарсить JSON вручную для отладки
	var rawData map[string]interface{}
	if err := json.Unmarshal(body, &rawData); err == nil {
		fmt.Printf("Parsed raw data: %+v\n", rawData)
	}

	// Восстанавливаем body для Gin
	c.Request.Body = io.NopCloser(bytes.NewBuffer(body))

	if err := c.ShouldBindJSON(&loginReq); err != nil {
		fmt.Printf("Bind error: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request: " + err.Error()})
		return
	}

	fmt.Printf("Parsed login request: Username=%s, Password=%s\n", loginReq.Username, loginReq.Password)

	// Простая проверка учетных данных
	if loginReq.Username == "admin" && loginReq.Password == "password" {
		// Генерируем токен
		token, err := auth.GenerateToken(loginReq.Username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"token":   token,
			"message": "Login successful",
		})
		return
	}

	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
}
