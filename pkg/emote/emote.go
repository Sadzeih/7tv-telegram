package emote

import (
	"github.com/sahilm/fuzzy"
)

// Emote is a standardized struct for an emote
type Emote struct {
	ID     string
	Name   string
	Type   string
	Width  int
	Height int
	URL    string
}

// emotes is a slice of emotes
type emotes []Emote

// Len returns the length of emotes
func (e emotes) Len() int {
	return len(e)
}

// String returns the name of the emote at the given index
func (e emotes) String(i int) string {
	return e[i].Name
}

// SearchEmotes fuzzy searches emotes from the query text
func SearchEmotes(query string, e []Emote) []Emote {
	matches := fuzzy.FindFrom(query, emotes(e))

	rankedEmotes := make([]Emote, len(matches))
	for i, match := range matches {
		rankedEmotes[i] = e[match.Index]
	}

	return rankedEmotes
}
