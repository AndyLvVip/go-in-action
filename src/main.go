package main

import (
	_ "config"
	_ "matchers"
	"search"
)

func main() {
	search.Run("president")
}
