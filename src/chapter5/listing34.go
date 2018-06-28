package main

import (
	"os"
	"fmt"
	"net/http"
	"io"
)

func init() {
	if len(os.Args) != 2 {
		fmt.Println("./listing32 <url>")
		os.Exit(-1)
	}
}

func main() {
	res, err := http.Get(os.Args[1])
	if nil != err {
		fmt.Println(err)
		return
	}

	io.Copy(os.Stdout, res.Body)
	if err = res.Body.Close(); nil != err {
		fmt.Println(err)
	}
}
