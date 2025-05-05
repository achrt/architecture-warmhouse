package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "smarthome",
	Short: "Умный дом, но в консоли",
	Long:  `CLI для запуска различных сервисов в экосистеме умного дома.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Ошибка запуска:", err)
	}
}
