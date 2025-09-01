package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Секретный ключ (в production должен храниться в безопасном месте)
var jwtSecret = []byte("your-secret-key-change-in-production")

// Claims структура для хранения данных в токене
type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// GenerateToken создает JWT токен
func GenerateToken(username string) (string, error) {
	// Создаем claims
	claims := &Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // Токен действителен 24 часа
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	// Создаем токен
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Подписываем токен
	return token.SignedString(jwtSecret)
}

// ValidateToken проверяет и валидирует JWT токен
func ValidateToken(tokenString string) (*Claims, error) {
	// Парсим токен
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// Проверяем метод подписи
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	// Проверяем claims
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
