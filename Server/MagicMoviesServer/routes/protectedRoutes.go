package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/swapno963/MovieDegital/MagicMoviesServer/ServerMagicMoviesServer/MagicMoviesServer/controllers"
	"github.com/swapno963/MovieDegital/MagicMoviesServer/ServerMagicMoviesServer/MagicMoviesServer/middleware"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func SetupProtectedRoutes(router *gin.Engine, client *mongo.Client) {
	router.Use(middleware.AuthMiddleWare())

	router.GET("/movie/:imdb_id", controllers.GetMovie(client))
	router.POST("/addmovie", controllers.AddMovie(client))
	router.PATCH("/updatereview/:imdb_id", controllers.AdminReviewUpdate(client))
	router.GET("/recommendedmovies", controllers.GetRecommendedMovies(client))

}
