package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/maimunar/calculator-api/internal/middleware"
)

func Handler(r *gin.Engine) {
	r.Use(middleware.AuthMiddleWare)

	r.POST("/add", AddHandler)
	r.POST("/subtract", SubtractHandler)
	r.POST("/divide", DivideHandler)
	r.POST("/multiply", MultiplyHandler)
	r.POST("/sum", SumHandler)
}
