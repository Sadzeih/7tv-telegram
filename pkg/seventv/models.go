package seventv

import (
	"github.com/Sadzeih/bttv-telegram/pkg/emote"
	"strings"
)

type Emotes []Emote

// Emote represents a 7TV emote model: yoinked from https://github.com/SevenTV/API/blob/dev/internal/rest/v2/model/emote.go#L13
type Emote struct {
	ID               string   `json:"id"`
	Name             string   `json:"name"`
	Visibility       int32    `json:"visibility"`
	VisibilitySimple []string `json:"visibility_simple"`
	Mime             string   `json:"mime"`
	Status           int8     `json:"status"`
	Tags             []string `json:"tags"`
	Width            int      `json:"width"`
	Height           int      `json:"height"`
	URL              string   `json:"url"`
}

// Convert transforms a BTTV emote into standardized emote struct
func (e Emote) Convert() emote.Emote {
	return emote.Emote{
		ID:     e.ID,
		Name:   e.Name,
		Type:   strings.ToLower(e.Mime),
		Width:  e.Width,
		Height: e.Height,
		URL:    e.URL,
	}
}
