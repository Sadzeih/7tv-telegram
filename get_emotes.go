package main

import (
	"github.com/Sadzeih/bttv-telegram/bttv"
	"github.com/Sadzeih/bttv-telegram/twitch"
)

func getEmotes(query string) (Emotes, error) {
	emotes := make(Emotes, 0)

	twitchEmotes, err := twitch.GlobalEmotes()
	if err != nil {
		return nil, err
	}
	emotes = append(emotes, ConvertTwitchEmotes(twitchEmotes)...)
	emotes = SearchEmotes(query, emotes)

	searchBTTV, err := bttv.SearchEmotes(query)
	if err != nil {
		return nil, err
	}
	emotes = append(emotes, ConvertBTTVEmotes(searchBTTV)...)

	globalBttv, err := bttv.GlobalEmotes()
	if err != nil {
		return nil, err
	}
	emotes = append(emotes, ConvertBTTVEmotes(globalBttv)...)
	emotes = SearchEmotes(query, emotes)

	return emotes, nil
}
