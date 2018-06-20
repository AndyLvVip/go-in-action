package search

import (
	"fmt"
	"log"
)

type Matcher interface {
	Search(feed *Feed, searchTerm string) ([]*Result, error)
}

type Result struct {
	Field   string
	Content string
}

func Display(results chan *Result) {
	for result := range results {
		fmt.Printf("%s\n%s\n", result.Field, result.Content)
	}
}

func Match(matcher Matcher, feed *Feed, searchTerm string, results chan *Result) {
	searchResults, err := matcher.Search(feed, searchTerm)
	if nil != err {
		log.Println(err)
		return
	}

	for _, result := range searchResults {
		results <- result
	}
}
