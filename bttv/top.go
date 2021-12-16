package bttv

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

const (
	trendingEndpoint = "https://api.betterttv.net/3/emotes/shared/trending"
	trendingLimit    = 50
)

type trendingResponse struct {
	Emote Emote `json:"emote"`
}

// TrendingEmotes gets 200 trending emotes from the last week
func TrendingEmotes() (Emotes, error) {
	emotes := make(Emotes, 0)

	url, err := url.Parse(trendingEndpoint)
	if err != nil {
		return nil, fmt.Errorf("could not create trending url: %w", err)
	}

	for i := 0; i < 4; i++ {
		offset := i * trendingLimit

		queryValues := url.Query()
		queryValues.Set("limit", fmt.Sprintf("%d", trendingLimit))
		queryValues.Set("offset", fmt.Sprintf("%d", offset))
		url.RawQuery = queryValues.Encode()

		resp, err := http.Get(url.String())
		if err != nil {
			return nil, fmt.Errorf("could not fetch %d-%d trending emotes: %w", offset, offset+i*trendingLimit, err)
		}

		if resp.StatusCode != 200 {
			return nil, fmt.Errorf("bttv returned non-200: %d", resp.StatusCode)
		}

		trendingEmotes := make([]trendingResponse, 0)
		if err := json.NewDecoder(resp.Body).Decode(&trendingEmotes); err != nil {
			return nil, fmt.Errorf("could not decode emotes: %w", err)
		}

		for _, t := range trendingEmotes {
			emotes = append(emotes, t.Emote)
		}
	}

	computeURLs(emotes)

	return emotes, nil
}
