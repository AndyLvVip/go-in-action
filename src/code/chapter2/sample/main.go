package main

import (
	_ "code/chapter2/sample/matchers"
	"code/chapter2/sample/search"
	"log"
	"os"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	search.Run("president")
}
