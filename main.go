package main

import (
	"github.com/Sadzeih/bttv-telegram/pkg/emote"
	"log"
	"math"
	"os"
	"strconv"

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
		var emotes []emote.Emote
		if q.Text == "" {
			return
		}
		emotes, err := getEmotes(q.Text)
		if err != nil {
			log.Println(err)
			return
		}
		emotes = emotes[:int(math.Min(50, float64(len(emotes))))]

		results := make(tb.Results, len(emotes))
		for i, e := range emotes {
			var result tb.Result
			switch e.Type {
			case "png":
				result = &tb.PhotoResult{
					URL:      e.URL,
					ThumbURL: e.URL,
					Width:    e.Width,
					Height:   e.Height,
				}
			case "gif", "webp":
				result = &tb.GifResult{
					URL:       e.URL,
					ThumbURL:  e.URL,
					ThumbMIME: "image/gif",
					Width:     e.Width,
					Height:    e.Height,
				}

			default:
				result = nil
			}

			if result == nil {
				continue
			}
			result.SetResultID(strconv.Itoa(i + 1))
			results[i] = result
		}

		err = b.Answer(q, &tb.QueryResponse{
			Results:   results,
			CacheTime: 0,
		})
		if err != nil {
			log.Println(err)
		}
	})

	b.Start()
}
