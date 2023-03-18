package twitch

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	cc "golang.org/x/oauth2/clientcredentials"
)

const (
	globalEmoteEndpoint = "https://api.twitch.tv/helix/chat/emotes/global"
	tokenEndpoint       = "https://id.twitch.tv/oauth2/token"
)

var (
	ccConfig *cc.Config
	client   *http.Client
)

func init() {
	ccConfig = &cc.Config{
		// TODO: better config handling (with viper)
		ClientID:     os.Getenv("TWITCH_CLIENT_ID"),
		ClientSecret: os.Getenv("TWITCH_CLIENT_SECRET"),
		TokenURL:     tokenEndpoint,
	}

	client = ccConfig.Client(context.Background())
}

// GlobalEmotes returns a list of global emotes
func GlobalEmotes() (Emotes, error) {
	r, err := http.NewRequest(http.MethodGet, globalEmoteEndpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("could not create global emotes request from twitch: %w", err)
	}
	r.Header.Add("client-id", ccConfig.ClientID)

	resp, err := client.Do(r)
	if err != nil {
		return nil, fmt.Errorf("could not fetch global emotes from twitch: %w", err)
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("fetching global emotes returned non 200: %d", resp.StatusCode)
	}

	tr := &twitchResponse{
		Data: make(Emotes, 0),
	}
	if err := json.NewDecoder(resp.Body).Decode(&tr); err != nil {
		return nil, fmt.Errorf("could not parse global emotes: %w", err)
	}

	return tr.Data, nil
}
