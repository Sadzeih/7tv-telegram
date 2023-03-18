package bttv

import (
	"fmt"
	"github.com/Sadzeih/bttv-telegram/pkg/emote"
)

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

// Convert transforms a BTTV emote into standardized emote struct
func (e Emote) Convert() emote.Emote {
	return emote.Emote{
		ID:   e.ID,
		Name: e.Code,
		Type: e.Type,
		URL:  e.URL,
	}
}
