package exporter

import (
	"fmt"

	"github.com/jonavdm/steam-exporter/steam"
	"github.com/prometheus/client_golang/prometheus"
)

type SteamExporter struct {
	Steam            *steam.Steam
	Player           string
	currentlyPlaying *prometheus.Desc
	playtime         *prometheus.Desc
}

func NewSteamExporter(steam *steam.Steam, player string) *SteamExporter {
	return &SteamExporter{
		Steam:  steam,
		Player: player,
		currentlyPlaying: prometheus.NewDesc(
			"steam_user_currently_playing",
			"Whether a user is currently playing a game",
			[]string{"user_id", "user_name", "game_id", "game_name"},
			nil,
		),
		playtime: prometheus.NewDesc(
			"steam_games_playtime_minutes",
			"How much time has been wasted",
			[]string{"user_id", "user_name", "game_id", "game_name"},
			nil,
		),
	}
}

func (e *SteamExporter) Describe(ch chan<- *prometheus.Desc) {
	ch <- e.currentlyPlaying
}

func (e *SteamExporter) Collect(ch chan<- prometheus.Metric) {
	users, _ := e.Steam.GetPlayerSummaries(e.Player)

	if len(users.Response.Players) == 0 {
		return
	}

	username := users.Response.Players[0].Personaname

	for _, user := range users.Response.Players {
		if user.Gameid == "" {
			continue
		}

		ch <- prometheus.MustNewConstMetric(
			e.currentlyPlaying,
			prometheus.GaugeValue,
			1,
			user.Steamid, user.Personaname, user.Gameid, user.Gameextrainfo,
		)
	}

	games, _ := e.Steam.GetOwnedGames(e.Player)
	for _, game := range games.Response.Games {
		ch <- prometheus.MustNewConstMetric(
			e.playtime,
			prometheus.CounterValue,
			float64(game.PlaytimeForever+game.PlaytimeDisconnected),
			e.Player, username, fmt.Sprint(game.Appid), game.Name,
		)
	}
}
