package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "steam-exporter",
	Short: "A exporter for prometheus that lets you stalk your steam things",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	envKey := os.Getenv("STEAM_EXPORTER_KEY")
	envUser := os.Getenv("STEAM_EXPORTER_USER")

	rootCmd.PersistentFlags().StringP("key", "k", envKey, "The steam api key")
	rootCmd.PersistentFlags().StringP("user", "u", envUser, "The steam user")
}
