package main

import (
	"log"
	"math"
	"os"

	tb "gopkg.in/tucnak/telebot.v2"
)

func main() {
	b, err := tb.NewBot(tb.Settings{
		// TODO: better config handling (with viper)
		Token: os.Getenv("TOKEN"),
		Poller: &tb.Webhook{
			Listen: os.Getenv("LISTEN_ADDR"),
			Endpoint: &tb.WebhookEndpoint{
				PublicURL: os.Getenv("PUBLIC_URL"),
			},
		},
	})

	if err != nil {
		log.Fatal(err)
	}

	b.Handle(tb.OnQuery, func(q *tb.Query) {
		emotes, err := getEmotes()
		if err != nil {
			log.Println(err)
			return
		}

		if q.Text != "" {
			emotes = SearchEmotes(q.Text, emotes)
		}
		emotes = emotes[:int(math.Min(50, float64(len(emotes))))]

		results := make(tb.Results, 0)
		for _, emote := range emotes {
			var result tb.Result
			switch emote.Type {
			case "png":
				result = &tb.PhotoResult{
					Title:       emote.Name,
					Description: emote.Name,
					URL:         emote.URL,
					ThumbURL:    emote.URL,
				}

			case "gif":
				result = &tb.GifResult{
					Title:    emote.Name,
					URL:      emote.URL,
					ThumbURL: emote.URL,
				}

			default:
				result = nil
			}

			if result == nil {
				continue
			}
			result.SetResultID(emote.ID)
			results = append(results, result)
		}

		err = b.Answer(q, &tb.QueryResponse{
			Results:   results,
			CacheTime: 60,
		})
		if err != nil {
			log.Println(err)
		}
	})

	b.Start()
}
