package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go timed_fetch(url, ch)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) // Recibiendo del canal, se bloquea hasta que recibe algo
	}
	fmt.Printf("%.2fs tiempo total", time.Since(start).Seconds())
}

func timed_fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	byteCount, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("Error leyendo %s, %v", url, err)
		return
	}
	elapsed := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %07d %s", elapsed, byteCount, url)
}
