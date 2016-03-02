package redditLib

import (
	"fmt"
	"log"
	"encoding/json"
	"net/http"
	"errors"
)

func GetSubReddit(subReddit string) (*RedditResponse, error) {

	req, err := http.NewRequest("GET", "https://reddit.com/r/" + subReddit + ".json", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("User-Agent", "RedditQuick Golang")

	client := http.Client{}

	res, err := client.Do(req)

	if res.StatusCode == 429 {
		return nil, errors.New("Too many requests try again later")
	}

	defer res.Body.Close()

	if err != nil {
		log.Fatal(err)
	}

	response := new(RedditResponse)

	err = json.NewDecoder(res.Body).Decode(response)

	if err != nil {
		log.Fatal(err)
	}

	return response, nil
}

func PrintSubReddit(subReddit *RedditResponse) {
	for _, child := range subReddit.Data.Children {
		fmt.Printf("%s (%d)\n", child.Data.Title, child.Data.CommentCount)
		fmt.Println(child.Data.URL)
		fmt.Println("")
	}
}

type Item struct {
	Title        string
	URL          string
	CommentCount int `json:"num_comments"`
}

type RedditResponse struct {
	Data struct {
		     Children []struct {
			     Data Item
		     }
	     }
}