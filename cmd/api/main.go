package main

import (
	"log"
	"os"

	"guitaclara/internal/infrastructure/httpapi"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := gin.Default()
	httpapi.RegisterRoutes(r)

	log.Printf("API listening on :%s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
