package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/maimunar/calculator-api/api"
)

func DivideHandler(c *gin.Context) {
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

	result := body.Number1 / body.Number2

	c.JSON(200, gin.H{
		"result": result,
	})
}
