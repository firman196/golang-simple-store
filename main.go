package main

import (
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

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	/*server := gin.Default()
	routes.ApiRoutes(server)
	server.Run(":" + port)*/
}
