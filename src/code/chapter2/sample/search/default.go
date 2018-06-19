package search

type defaultMatcher struct {}

func (matcher defaultMatcher) Search(feed* Feed, searchTem string) ([]*Result, error) {
	return nil, nil
}

func init() {
	var matcher defaultMatcher
	Register("default", matcher)
}
