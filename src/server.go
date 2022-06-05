package main

import (
	"log"
	"net/http"
	"os"

	"record_song_key_service/src/driver"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func loadEnv() {
	err := godotenv.Load("../.env")

	if err != nil {
		log.Println(err)
	}
}

var dbDriver *gorm.DB

func initialize() {
	loadEnv()

	dbDriver := driver.GetInstance()

	log.Printf(os.Getenv("POSTGRES_HOST"))
	dbDriver.GetConnection(
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_OUT_PORT"),
	)
}

func main() {
	initialize()

	engine := gin.Default()

	engine.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello record music key service")
	})

	engine.Run(":8000")
}
