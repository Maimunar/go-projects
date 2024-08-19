package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func AddIdMiddleWare(c *gin.Context) {
	c.Set("request_id", uuid.New().String())
}
