package main

import (
	"github.com/Sadzeih/bttv-telegram/bttv"
	"github.com/Sadzeih/bttv-telegram/twitch"
)

func getEmotes() (Emotes, error) {
	emotes := make(Emotes, 0)

	twitchEmotes, err := twitch.GlobalEmotes()
	if err != nil {
		return nil, err
	}
	emotes = append(emotes, ConvertTwitchEmotes(twitchEmotes)...)

	trendingBttv, err := bttv.TrendingEmotes()
	if err != nil {
		return nil, err
	}
	emotes = append(emotes, ConvertBTTVEmotes(trendingBttv)...)

	globalBttv, err := bttv.GlobalEmotes()
	if err != nil {
		return nil, err
	}
	emotes = append(emotes, ConvertBTTVEmotes(globalBttv)...)

	return emotes, nil
}
