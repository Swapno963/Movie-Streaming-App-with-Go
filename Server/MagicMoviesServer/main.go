package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/swapno963/MovieDegital/MagicMoviesServer/ServerMagicMoviesServer/MagicMoviesServer/controllers"
)

func main() {
	router := gin.Default()
	router.GET("/hello", func(c *gin.Context) {
		c.String(200, "Hello Go coders")
	})

	router.GET("/movies", controllers.GetMovies())
	router.GET("/movies/:imdb_id", controllers.GetMovie())

	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Warning: unable to find .env file")
	}

	if err := router.Run(":8080"); err != nil {
		fmt.Println("Failed to start", err)
	}
}
