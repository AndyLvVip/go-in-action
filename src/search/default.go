package search

import "fmt"

type defaultMatcher struct {}

func (matcher defaultMatcher) Search(feed *Feed, searchTem string) ([]*Result, error) {
	return nil, nil
}

func init() {
	fmt.Println("default main...")
	var matcher defaultMatcher
	Register("default", matcher)
}
