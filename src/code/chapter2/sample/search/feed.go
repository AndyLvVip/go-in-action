package search

import (
	"os"
	"encoding/json"
)

const dataFile = "G:/workspace/go-workspace/go-in-action/src/code/chapter2/sample/data/data.json"

type Feed struct {
	Name string `json:"site"`
	URI  string `json:"link"`
	Type string `json:"type"`
}

func RetrieveFeeds() ([]*Feed, error) {
	file, err := os.Open(dataFile)
	if nil != err {
		return nil, err
	}
	defer file.Close()

	var feeds []*Feed
	err = json.NewDecoder(file).Decode(feeds)
	return feeds, err
}
