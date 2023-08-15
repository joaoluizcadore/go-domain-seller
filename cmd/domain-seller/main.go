package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joaoluizcadore/domain-seller/config"
	"github.com/joho/godotenv"
)

func initializeServer() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Pong!",
		})
	})

	router.Run(fmt.Sprintf(":%v", config.GetConfig().ServerPort))
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("Error loading .env file: %v", err)
		panic("Error loading .env file")
	}

	log.Println("Starting server...")
	initializeServer()

}
