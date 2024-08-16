package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/maimunar/calculator-api/api"
)

func SumHandler(c *gin.Context) {
	var body api.NumberArraysBody

	err := c.BindJSON(&body)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid input",
		})
		return
	}

	sum := 0
	for _, number := range body.Numbers {
		sum += number
	}

	c.JSON(200, gin.H{
		"result": sum,
	})
}
