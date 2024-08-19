package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/maimunar/calculator-api/internal/database"
	"github.com/maimunar/calculator-api/internal/middleware"
)

func Handler(r *gin.Engine, repo *database.SQLiteRepository) {
	r.Use(middleware.RateLimiter)
	r.Use(middleware.AuthMiddleWare)
	r.Use(middleware.AddIdMiddleWare)

	r.POST("/add", AddHandler(repo))
	r.POST("/subtract", SubtractHandler(repo))
	r.POST("/divide", DivideHandler(repo))
	r.POST("/multiply", MultiplyHandler(repo))
	r.POST("/sum", SumHandler(repo))

	r.GET("/calculations", CalculationsHandler(repo))
}
