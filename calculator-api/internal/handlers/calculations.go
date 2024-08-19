package handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/maimunar/calculator-api/internal/database"
)

func CalculationsHandler(r *database.SQLiteRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		calculations, err := r.GetCalculations()
		if err != nil {
			c.JSON(500, gin.H{
				"message": "Internal server error",
			})
		}

		fmt.Println(calculations)
		c.JSON(200, calculations)
	}
}
