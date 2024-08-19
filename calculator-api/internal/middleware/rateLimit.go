package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

func RateLimiter(c *gin.Context) {
	limiter := rate.NewLimiter(1, 4)
	if limiter.Allow() {
		c.Next()
	} else {
		c.JSON(http.StatusTooManyRequests, gin.H{
			"message": "Limit exceed",
		})
	}
}
