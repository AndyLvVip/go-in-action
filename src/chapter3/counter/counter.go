package main

import (
	"fmt"
	"github.com/goinaction/code/chapter3/words"
	"io/ioutil"
	"os"
)

func main() {
	for _, fileName := range os.Args[1:] {

		file, err := os.Open(fileName)
		if nil != err {
			fmt.Println(err)
			return
		}
		data, err := ioutil.ReadAll(file)
		file.Close()
		if nil != err {
			fmt.Println(err)
			return
		}

		count := words.CountWords(string(data))
		fmt.Printf("There are %d words in %s\n", count, fileName)
	}
}
