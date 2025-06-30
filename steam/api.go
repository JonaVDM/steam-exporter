package steam

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Steam struct {
	Key        string
	ApiBase    string
	HttpClient http.Client
}

func NewSteam(key string) Steam {
	return Steam{
		ApiBase:    "https://api.steampowered.com",
		Key:        key,
		HttpClient: http.Client{},
	}
}

// Execute performs a GET request to the Steam API.
// It appends the API key to the query parameters.
func (s *Steam) Execute(path string, params []string) ([]byte, error) {
	query := url.Values{}
	query.Set("key", s.Key)

	for i := 0; i < len(params)-1; i += 2 {
		query.Set(params[i], params[i+1])
	}

	fullUrl := fmt.Sprintf("%s%s?%s", s.ApiBase, path, query.Encode())

	resp, err := s.HttpClient.Get(fullUrl)
	if err != nil {
		return nil, fmt.Errorf("http GET failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("bad response: %s - %s", resp.Status, string(body))
	}

	return io.ReadAll(resp.Body)
}

// GetPlayerSummaries returns the summary of a player using their SteamID.
func (s *Steam) GetPlayerSummaries(steamId string) (PlayerSummariesResponse, error) {
	response := PlayerSummariesResponse{}
	body, err := s.Execute("/ISteamUser/GetPlayerSummaries/v0002", []string{
		"steamids", steamId,
	})
	if err != nil {
		return response, err
	}

	if err := json.Unmarshal(body, &response); err != nil {
		return response, fmt.Errorf("failed to decode JSON: %w", err)
	}

	return response, nil
}

// GetOwnedGames returns the list of owned games for a given SteamID.
func (s *Steam) GetOwnedGames(steamId string) (OwnedGamesResponse, error) {
	response := OwnedGamesResponse{}
	body, err := s.Execute("/IPlayerService/GetOwnedGames/v0001", []string{
		"steamid", steamId,
		"include_appinfo", "1",
		"include_played_free_games", "1",
	})
	if err != nil {
		return response, err
	}

	if err := json.Unmarshal(body, &response); err != nil {
		return response, fmt.Errorf("failed to decode JSON: %w", err)
	}

	return response, nil
}
