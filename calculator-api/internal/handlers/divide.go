package handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/maimunar/calculator-api/api"
	"github.com/maimunar/calculator-api/internal/database"
)

func DivideHandler(r *database.SQLiteRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		var body api.TwoNumbersBody

		err := c.BindJSON(&body)
		if err != nil {
			c.JSON(400, gin.H{
				"message": "Invalid input",
			})
			return
		}

		if body.Number2 == 0 {
			c.JSON(500, gin.H{
				"message": "Cannot divide by zero",
			})
			return
		}

		input := fmt.Sprintf("%d/%d", body.Number1, body.Number2)
		result := body.Number1 / body.Number2
		r.AddCalculation(input, result)

		c.JSON(200, gin.H{
			"result": result,
		})
	}
}
