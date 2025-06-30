package cmd

import (
	"net/http"

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
		api := steam.NewSteam(key)

		exporter := exporter.NewSteamExporter(&api, user)
		prometheus.MustRegister(exporter)

		http.Handle("/metrics", promhttp.Handler())
		return http.ListenAndServe(":2112", nil)
	},
}

func init() {
	runCmd.PersistentFlags().StringP("key", "k", "", "The steam api key")
	runCmd.PersistentFlags().StringP("user", "u", "", "The steam user")

	rootCmd.AddCommand(runCmd)
}
