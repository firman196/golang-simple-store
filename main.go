package main

import (
	"golang-graphql/database"
	"os"

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
	_, dbErr := database.ConnectDB(dbUsername, dbPassword, dbHost, dbPort, dbName)

	if dbErr != nil {
		log.Fatal().Err(dbErr).Msg("cannot connect to database")
	}

	/*server := gin.Default()
	routes.ApiRoutes(server)
	server.Run(":" + port)*/
}
