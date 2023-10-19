package app

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	client2 "secondTask/internal/client"
	"secondTask/internal/controllers"
	routes2 "secondTask/internal/routes"
	"time"
)

func Run() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal().Err(err).Msg("Error loading .env file")
	}
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})
	log.Info().Msg("Logger initialized")

	apiUrl := os.Getenv("API_URL")
	client := client2.NewClient(apiUrl)
	controller := controllers.NewApiController(*client)
	routes := routes2.NewRoutes(*controller)

	go func() {
		for {
			client.FetchCoinGeckoData()
			time.Sleep(10 * time.Minute)
		}
	}()

	r := gin.Default()
	routes.SetupAPIRoutes(r)
	err = r.Run(":8080")
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to start the server")
	}
}
