package cmd

import (
	"net/http"
	"os"

	"github.com/jonavdm/steam-exporter/exporter"
	"github.com/jonavdm/steam-exporter/steam"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run the exporter",
	RunE: func(cmd *cobra.Command, args []string) error {
		key, _ := cmd.Flags().GetString("key")
		user, _ := cmd.Flags().GetString("user")
		host, _ := cmd.Flags().GetString("host")
		api := steam.NewSteam(key)

		exporter := exporter.NewSteamExporter(&api, user)
		prometheus.MustRegister(exporter)

		http.Handle("/metrics", promhttp.Handler())
		return http.ListenAndServe(host, nil)
	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	host := os.Getenv("STEAM_EXPORTER_HOST")
	if host == "" {
		host = ":6718"
	}

	runCmd.Flags().String("host", host, "The interface/port to listen on")
}
