package engine

import (
	"eg/egolang/crawler/fetcher"
	"log"
)

func worker(r Request) (ParseResult, error) {
	log.Printf("Fetching: %s", r.Url)

	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetcher: error fetching %s: %c", r.Url, err)
		return ParseResult{}, err
	}
	return r.Parser.Parse(body, r.Url), nil

}
