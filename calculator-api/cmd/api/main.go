package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/maimunar/calculator-api/internal/database"
	"github.com/maimunar/calculator-api/internal/handlers"
)

func main() {
	r := gin.Default()

	repo := database.OpenDB()
	defer repo.Close()

	handlers.Handler(r, &repo)

	fmt.Println("Starting GO API service...")

	err := r.Run(":8080")

	if err != nil {
		log.Fatalln("Error starting GO API service: ", err)
	}

}
