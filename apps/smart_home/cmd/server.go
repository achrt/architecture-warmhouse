package cmd

import (
	"log"
	"smarthome/db"
	"smarthome/handlers"
	"smarthome/services"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Запуск серверной части (датчики, устройства, API)",
	Run: func(cmd *cobra.Command, args []string) {
		dbURL := GetEnv("DATABASE_URL", "postgres://postgres:postgres@localhost:5432/smarthome")
		temperatureAPIURL := GetEnv("TEMPERATURE_API_URL", "http://localhost:8081")
		port := GetEnv("PORT", ":8080")

		database, err := db.New(dbURL)
		if err != nil {
			log.Fatalf("DB connection error: %v", err)
		}
		defer database.Close()

		temperatureService := services.NewTemperatureService(temperatureAPIURL)
		sensorHandler := handlers.NewSensorHandler(database, temperatureService)

		router := gin.Default()
		AddHealthCheck(router)
		sensorHandler.RegisterRoutes(router.Group("/api/v1"))

		ServeHTTP(router, port)
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
