package main

import (
	"golang-store/controller"
	"golang-store/database"
	"golang-store/repository"
	"golang-store/service"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

func main() {
	//load environtment
	envErr := godotenv.Load(".env")
	if envErr != nil {
		log.Fatal().Err(envErr).Msg("cannot load environment")
	}

	appPort := os.Getenv("PORT")
	if appPort == "" {
		appPort = "8080"
	}

	var dbUsername = os.Getenv("DB_USERNAME")
	var dbPassword = os.Getenv("DB_PASSWORD")
	var dbName = os.Getenv("DB_DATABASE")
	var dbHost = os.Getenv("DB_HOST")
	var dbPort = os.Getenv("DB_PORT")
	db, dbErr := database.ConnectDB(dbUsername, dbPassword, dbHost, dbPort, dbName)

	if dbErr != nil {
		log.Fatal().Err(dbErr).Msg("cannot connect to database")
	}

	//repository layer
	categoryRepository := repository.NewCategoryRepositoryImpl(db)

	//service layer
	categoryService := service.NewCategoryServiceImpl(categoryRepository)

	//controller layer
	categoryController := controller.NewCategoryControllerImpl(categoryService)

	router := gin.Default()
	router.Use(cors.Default())
	api := router.Group("/api/v1")

	//route category
	api.POST("/category", categoryController.Create)
	api.PUT("/category", categoryController.Update)
	api.GET("/category/:id", categoryController.GetById)
	api.GET("/category", categoryController.GetAll)

	router.Run(":" + appPort)
}
