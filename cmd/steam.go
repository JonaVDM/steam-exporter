package cmd

import (
	"fmt"

	"github.com/jonavdm/steam-exporter/steam"
	"github.com/spf13/cobra"
)

// steamCmd represents the steam command
var steamCmd = &cobra.Command{
	Use:   "steam",
	Short: "Interact with the steam api",
}

// steamCmd represents the steam command
var steamPlayerSumCmd = &cobra.Command{
	Use:   "playersum",
	Short: "Execute the player summaries api call",
	RunE: func(cmd *cobra.Command, args []string) error {
		key, _ := cmd.Flags().GetString("key")
		user, _ := cmd.Flags().GetString("user")
		api := steam.NewSteam(key)
		data, err := api.GetPlayerSummaries(user)
		if err != nil {
			return err
		}

		fmt.Println(data.Response.Players[0].Personaname)

		return nil
	},
}

// steamCmd represents the steam command
var steamGamesCmd = &cobra.Command{
	Use:   "games",
	Short: "Execute the player summaries api call",
	RunE: func(cmd *cobra.Command, args []string) error {
		key, _ := cmd.Flags().GetString("key")
		user, _ := cmd.Flags().GetString("user")
		api := steam.NewSteam(key)
		data, err := api.GetOwnedGames(user)
		if err != nil {
			return err
		}

		fmt.Println(data.Response.GameCount)

		return nil
	},
}

func init() {
	steamCmd.PersistentFlags().StringP("key", "k", "", "The steam api key")
	steamCmd.PersistentFlags().StringP("user", "u", "", "The steam user")

	steamCmd.AddCommand(steamPlayerSumCmd)
	steamCmd.AddCommand(steamGamesCmd)

	rootCmd.AddCommand(steamCmd)
}
