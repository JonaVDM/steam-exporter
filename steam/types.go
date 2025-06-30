package steam

type PlayerSummariesResponse struct {
	Response struct {
		Players []struct {
			Steamid                  string `json:"steamid"`
			Communityvisibilitystate int    `json:"communityvisibilitystate"`
			Profilestate             int    `json:"profilestate"`
			Personaname              string `json:"personaname"`
			Profileurl               string `json:"profileurl"`
			Avatar                   string `json:"avatar"`
			Avatarmedium             string `json:"avatarmedium"`
			Avatarfull               string `json:"avatarfull"`
			Avatarhash               string `json:"avatarhash"`
			Lastlogoff               int    `json:"lastlogoff"`
			Personastate             int    `json:"personastate"`
			Realname                 string `json:"realname"`
			Primaryclanid            string `json:"primaryclanid"`
			Timecreated              int    `json:"timecreated"`
			Personastateflags        int    `json:"personastateflags"`
			Gameextrainfo            string `json:"gameextrainfo"`
			Gameid                   string `json:"gameid"`
		} `json:"players"`
	} `json:"response"`
}

type OwnedGamesResponse struct {
	Response struct {
		GameCount int `json:"game_count"`
		Games     []struct {
			Appid                    int    `json:"appid"`
			Name                     string `json:"name"`
			PlaytimeForever          int    `json:"playtime_forever"`
			ImgIconURL               string `json:"img_icon_url"`
			PlaytimeWindowsForever   int    `json:"playtime_windows_forever"`
			PlaytimeMacForever       int    `json:"playtime_mac_forever"`
			PlaytimeLinuxForever     int    `json:"playtime_linux_forever"`
			PlaytimeDeckForever      int    `json:"playtime_deck_forever"`
			RtimeLastPlayed          int    `json:"rtime_last_played"`
			PlaytimeDisconnected     int    `json:"playtime_disconnected"`
			HasCommunityVisibleStats bool   `json:"has_community_visible_stats,omitempty"`
			ContentDescriptorids     []int  `json:"content_descriptorids,omitempty"`
			HasLeaderboards          bool   `json:"has_leaderboards,omitempty"`
			Playtime2Weeks           int    `json:"playtime_2weeks,omitempty"`
		} `json:"games"`
	} `json:"response"`
}
