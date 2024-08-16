package middleware

import (
	"log/slog"
	"os"

	"github.com/gin-gonic/gin"
)

func getAuthToken() string {
	correctToken, ex := os.LookupEnv("CALCULATOR_TOKEN")
	if !ex {
		slog.Warn("CALCULATOR_TOKEN environment variable not set, using default value...")
		correctToken = "default-token"
	}
	correctToken = "Bearer " + correctToken

	return correctToken
}

func AuthMiddleWare(c *gin.Context) {
	// Get token from environment variable
	correctToken := getAuthToken()

	token := c.GetHeader("Authorization")

	if token != correctToken {
		slog.Error("user is not authenticated")
		c.JSON(401, gin.H{
			"message": "Unauthorized - invalid token",
		})
		c.Abort()
		return
	}

	slog.Info("user is authenticated")
}
