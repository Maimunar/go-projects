package handlers

import (
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/maimunar/calculator-api/api"
	"github.com/maimunar/calculator-api/internal/database"
)

func SumHandler(r *database.SQLiteRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		var body api.NumberArraysBody

		err := c.BindJSON(&body)
		if err != nil {
			c.JSON(400, gin.H{
				"message": "Invalid input",
			})
			return
		}

		inputBuilder := strings.Builder{}

		sum := 0
		for _, number := range body.Numbers {
			sum += number

			inputBuilder.WriteString(strconv.Itoa(number))
			inputBuilder.WriteString("+")
		}
		input := inputBuilder.String()[:inputBuilder.Len()-1]

		r.AddCalculation(input, sum)

		c.JSON(200, gin.H{
			"result": sum,
		})
	}
}
