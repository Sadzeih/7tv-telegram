package main

import (
	"github.com/Sadzeih/bttv-telegram/bttv"
	"github.com/Sadzeih/bttv-telegram/twitch"
	"github.com/sahilm/fuzzy"
)

// Emote is a standardized struct for an emote
type Emote struct {
	ID   string
	Name string
	Type string
	URL  string
}

// Emotes is a slice of Emotes
type Emotes []Emote

// Len returns the length of Emotes
func (e Emotes) Len() int {
	return len(e)
}

// String returns the name of the emote at the given index
func (e Emotes) String(i int) string {
	return e[i].Name
}

// SearchEmotes fuzzy searches emotes from the query text
func SearchEmotes(query string, e Emotes) Emotes {
	matches := fuzzy.FindFrom(query, e)

	rankedEmotes := make(Emotes, len(matches))
	for i, match := range matches {
		rankedEmotes[i] = e[match.Index]
	}

	return rankedEmotes
}

// ConvertTwitchEmotes transforms Twitch emotes into standardized emote struct
func ConvertTwitchEmotes(t twitch.Emotes) Emotes {
	emotes := make(Emotes, len(t))

	for idx, te := range t {
		e := Emote{
			ID:   te.ID,
			Name: te.Name,
			Type: "png",
		}

		if te.Images.URL3x != "" {
			e.URL = te.Images.URL3x
		} else if te.Images.URL2x != "" {
			e.URL = te.Images.URL2x
		} else {
			e.URL = te.Images.URL1x
		}

		emotes[idx] = e
	}

	return emotes
}

// ConvertBTTVEmotes transforms Twitch emotes into standardized emote struct
func ConvertBTTVEmotes(be bttv.Emotes) Emotes {
	emotes := make(Emotes, len(be))

	for idx, b := range be {
		e := Emote{
			ID:   b.ID,
			Name: b.Code,
			Type: b.Type,
			URL:  b.URL,
		}

		emotes[idx] = e
	}

	return emotes
}
