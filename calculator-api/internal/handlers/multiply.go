package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/maimunar/calculator-api/api"
)

func MultiplyHandler(c *gin.Context) {
	var body api.TwoNumbersBody

	err := c.BindJSON(&body)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid input",
		})
		return
	}

	result := body.Number1 * body.Number2

	c.JSON(200, gin.H{
		"result": result,
	})
}
