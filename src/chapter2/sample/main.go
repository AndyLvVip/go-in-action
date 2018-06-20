package main

import (
	_ "chapter2/sample/config"
	_ "chapter2/sample/matchers"
	"chapter2/sample/search"
)

func main() {
	search.Run("president")
}
