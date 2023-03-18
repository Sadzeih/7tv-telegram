package twitch

import "github.com/Sadzeih/bttv-telegram/pkg/emote"

type twitchResponse struct {
	Data []Emote `json:"data"`
}

// Emote represents an emote from Twitch
type Emote struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Images struct {
		URL1x string `json:"url_1x"`
		URL2x string `json:"url_2x"`
		URL3x string `json:"url_3x"`
	} `json:"images"`
}

// Emotes is a slice of Emotes
type Emotes []Emote

// Convert transforms a Twitch emote into standardized emote struct
func (e Emote) Convert() emote.Emote {
	ee := emote.Emote{
		ID:     e.ID,
		Name:   e.Name,
		Width:  128,
		Height: 128,
		Type:   "png",
	}
	if e.Images.URL3x != "" {
		ee.URL = e.Images.URL3x
	} else if e.Images.URL2x != "" {
		ee.URL = e.Images.URL2x
	} else {
		ee.URL = e.Images.URL1x
	}
	return ee
}
