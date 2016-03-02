package main

import (
	"os"
	"redditQuick/redditLib"
	"log"
)

func main() {
	if len(os.Args) < 2 {
		println("No subreddit specified")
		os.Exit(1)
	}

	subReddit := os.Args[1];

	redditResponse, err := redditLib.GetSubReddit(subReddit)

	if err != nil {
		log.Fatal(err)
	}

	redditLib.PrintSubReddit(redditResponse)

}