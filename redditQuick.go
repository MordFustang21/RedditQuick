package main

import (
	"os"
	"redditQuick/redditLib"
	"log"
)

func main() {
	//TODO: Switch to flags
	var subReddit string
	if len(os.Args) < 2 {
		println("No subreddit specified defaulting to Golang")
		subReddit = "golang"
	} else {
		subReddit = os.Args[1];
	}

	redditResponse, err := redditLib.GetSubReddit(subReddit)

	if err != nil {
		log.Fatal(err)
	}

	redditLib.PrintSubReddit(redditResponse)

}
