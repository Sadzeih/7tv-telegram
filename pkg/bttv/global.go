package bttv

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	globalEndpoint = "https://api.betterttv.net/3/cached/emotes/global"
)

// GlobalEmotes returns BTTV global emotes
func GlobalEmotes() (Emotes, error) {
	resp, err := http.Get(globalEndpoint)
	if err != nil {
		return nil, fmt.Errorf("could not fetch bttv global emotes: %w", err)
	}

	emotes := make(Emotes, 0)
	err = json.NewDecoder(resp.Body).Decode(&emotes)
	if err != nil {
		return nil, fmt.Errorf("could not decode emotes: %w", err)
	}

	computeURLs(emotes)

	return emotes, nil
}
