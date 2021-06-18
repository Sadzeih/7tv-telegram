package twitch

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
