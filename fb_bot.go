// Servidor para responder mensajes de Fb

package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/webhook", Verify)
	http.HandleFunc("/", Index)
	fmt.Println("Starting Server, port :9090")
	log.Fatal(http.ListenAndServe(":9090", nil))
}

func Index(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	// no se lo que hace Fprintf
	fmt.Fprintf(w, "<h1>Bienvenido al Bot responder</h1>")
	return
}

func Verify(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	if r.Method == "GET" {
		challenge := r.URL.Query().Get("hub.challenge")
		token := r.URL.Query().Get("hub.verify_token")
		if len(token) > 0 && len(challenge) > 0 && token == "my-token" {
			fmt.Fprintf(w, challenge)
			return
		}
	}
	w.WriteHeader(400)
	fmt.Fprintf(w, "Bad Request")
}

type MessengerPayload struct {
	Entry []struct {
		Time      uint64 `json:"time,omitempty"`
		Messaging []struct {
			Sender struct {
				Id string `json:"id"`
			} `json:"sender,omitempty"`
			Recipient struct {
				Id string `json:id`
			} `json:"recipient,omitempty"`
			Timestamp uint64 `json:"timestamp,omitempty"`
			Message   *struct {
				Mid  string `json:"mid,omitempty"`
				Seq  uint64 `json:"seq,omitempty"`
				Text string `json:"text"`
			} `json:message,omitempty`
		} `json:"messaging"`
	}
}
