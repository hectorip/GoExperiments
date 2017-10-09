package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	for i, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("Error retrieving: %s, Error: %s", url, resp)
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("Error reading response from %s, Error: %s", url, err)
		}
		fmt.Printf("%d -> %s", i+1, body)
	}
}
