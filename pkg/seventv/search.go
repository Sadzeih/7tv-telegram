package seventv

import (
	"context"
	"fmt"
	graphql "github.com/hasura/go-graphql-client"
	"net/http"
)

const (
	gqlEndpoint = "https://7tv.io/v3/gql"
	limit       = 20
)

// SearchEmotes searches for emotes from the query
func SearchEmotes(query string) ([]Emote, error) {
	q := struct {
		Emotes struct {
			Items []struct {
				Id       string
				Name     string
				Animated bool
				Host     struct {
					Files []struct {
						Width  int
						Height int
					}
					Url string
				}
			}
		} `graphql:"emotes(query: $query, limit: $limit)"`
	}{}
	client := graphql.NewClient(gqlEndpoint, http.DefaultClient)
	err := client.Query(context.Background(), &q, map[string]interface{}{
		"query": query,
		"limit": limit,
	})
	if err != nil {
		return nil, fmt.Errorf("could not search emotes with query \"%s\": %w", query, err)
	}

	emotes := make([]Emote, len(q.Emotes.Items))
	for i, e := range q.Emotes.Items {
		emotes[i].ID = e.Id
		emotes[i].Name = e.Name
		file := e.Host.Files[len(e.Host.Files)-1]
		emotes[i].Width = file.Width
		emotes[i].Height = file.Height
		emotes[i].Mime = "png"
		if e.Animated {
			emotes[i].Mime = "gif"
		}
		emotes[i].URL = fmt.Sprintf("https:%s/%s.%s", e.Host.Url, "4x", emotes[i].Mime)
	}

	return emotes, nil
}
