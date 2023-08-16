package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joaoluizcadore/domain-seller/config"
	"github.com/joaoluizcadore/domain-seller/web/controller"
	"github.com/joho/godotenv"
)

func initializeServer() {
	router := gin.Default()

	router.LoadHTMLGlob(("web/templates/*.tmpl"))

	router.GET("/", controller.IndexAction)
	router.POST("/", controller.SendMessageAction)

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