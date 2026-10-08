package main

import (
	"log"
	"os"
	"tuko-gg/sentinel/db"
	"tuko-gg/sentinel/router"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// get db up to date
	db.GetDb().AutoMigrate(&db.Label{}, &db.LabelAction{})

	httpPort := os.Getenv("HTTP_PORT")

	app := gin.Default()

	router.InitRouter(app)

	app.Run(":" + httpPort)
}
