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

	queryValues := url.Query()
	queryValues.Set("query", query)
	queryValues.Set("limit", fmt.Sprintf("%d", limit))
	queryValues.Set("offset", "0")
	url.RawQuery = queryValues.Encode()

	resp, err := http.Get(url.String())
	if err != nil {
		return nil, fmt.Errorf("could not search emotes with query \"%s\": %w", query, err)
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("bttv returned non-200: %d", resp.StatusCode)
	}

	if err := json.NewDecoder(resp.Body).Decode(&emotes); err != nil {
		return nil, fmt.Errorf("could not decode emotes: %w", err)
	}

	computeURLs(emotes)

	return emotes, nil
}
