package main

import (
	"fmt"
	"ioutils"
	"net/http"
)

func main() {
	for i, url := range sys.Args[1:] {
		resp, err := http.Get(url)

	}
}
