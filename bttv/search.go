package bttv

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

const (
	searchEndpoint = "https://api.betterttv.net/3/emotes/shared/search"
)

// SearchEmotes searches for emotes from the query
func SearchEmotes(query string) (Emotes, error) {
	emotes := make(Emotes, 0)

	url, err := url.Parse(searchEndpoint)
	if err != nil {
		return nil, fmt.Errorf("could not create search URL: %w", err)
	}

	for i := 0; i < 4; i++ {
		offset := i * trendingLimit

		queryValues := url.Query()
		queryValues.Set("limit", fmt.Sprintf("%d", trendingLimit))
		queryValues.Set("offset", fmt.Sprintf("%d", offset))
		url.RawQuery = queryValues.Encode()

		resp, err := http.Get(url.String())
		if err != nil {
			return nil, fmt.Errorf("could not search emotes with query \"%s\": %w", query, err)
		}

		if resp.StatusCode != 200 {
			return nil, fmt.Errorf("bttv returned non-200: %d", resp.StatusCode)
		}

		searchedEmotes := make(Emotes, 0)
		if err := json.NewDecoder(resp.Body).Decode(&searchedEmotes); err != nil {
			return nil, fmt.Errorf("could not decode emotes: %w", err)
		}
	}

	return emotes, nil
}
