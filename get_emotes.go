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

	return emotes, nil
}
