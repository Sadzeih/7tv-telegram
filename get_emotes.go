package main

import (
	"github.com/Sadzeih/bttv-telegram/pkg/emote"
	"github.com/Sadzeih/bttv-telegram/pkg/seventv"
	"github.com/Sadzeih/bttv-telegram/pkg/twitch"
)

func getEmotes(query string) ([]emote.Emote, error) {
	twitchEmotes, err := twitch.GlobalEmotes()
	if err != nil {
		return nil, err
	}
	//searchBTTV, err := bttv.SearchEmotes(query)
	//if err != nil {
	//	return nil, err
	//}
	//globalBttv, err := bttv.GlobalEmotes()
	//if err != nil {
	//	return nil, err
	//}
	stv, err := seventv.SearchEmotes(query)
	if err != nil {
		return nil, err
	}

	emotes := make([]emote.Emote, len(twitchEmotes)+len(stv))
	for i, e := range twitchEmotes {
		emotes[i] = e.Convert()
	}
	//for i, e := range searchBTTV {
	//	emotes[i+len(twitchEmotes)] = e.Convert()
	//}
	//for i, e := range globalBttv {
	//	emotes[i+len(twitchEmotes)+len(searchBTTV)] = e.Convert()
	//}
	for i, e := range stv {
		emotes[i+len(twitchEmotes)] = e.Convert()
	}

	emotes = emote.SearchEmotes(query, emotes)

	return emotes, nil
}
