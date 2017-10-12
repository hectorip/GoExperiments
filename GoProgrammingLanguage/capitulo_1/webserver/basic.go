package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", basicHandler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func basicHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "Hello Writer Go!")
	return
}
