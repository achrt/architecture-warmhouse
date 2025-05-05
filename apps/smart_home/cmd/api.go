package cmd

import (
	"smarthome/handlers"
	"smarthome/services"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "Запуск внешнего сервиса температуры",
	Run: func(cmd *cobra.Command, args []string) {
		port := GetEnv("PORT", ":8081")

		router := gin.Default()
		AddHealthCheck(router)

		handler := handlers.NewTemperatureHandler(services.NewTemperatureApi())
		handler.RegisterRoutes(router.Group("/"))

		ServeHTTP(router, port)
	},
}

func init() {
	rootCmd.AddCommand(apiCmd)
}
