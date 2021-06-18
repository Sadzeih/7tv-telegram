package bttv

import "fmt"

// Emote is a BTTV emote
type Emote struct {
	ID   string `json:"id"`
	Code string `json:"code"`
	Type string `json:"imageType"`
	URL  string
}

// Emotes is a slice of Emotes
type Emotes []Emote

const (
	emoteURLFormat = "https://cdn.betterttv.net/emote/%s/3x"
)

func computeURLs(e []Emote) {
	for idx := range e {
		e[idx].URL = fmt.Sprintf(emoteURLFormat, e[idx].ID)
	}
}
