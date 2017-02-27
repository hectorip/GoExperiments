// Servidor para responder mensajes de Fb
package main

import (
    "fmt"
    "log"
    "net/http"
)

func main() {
    http.HandleFunc("/webhook", Verify)
    // http.HandleFunc("/", Index)
    fmt.Println("Starting Server in port :9090")
    log.Fatal(http.ListenAndServe(":9090", nil))
}

func Verify(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/plain")
    if r.Method == "GET" {
        challenge := r.URL.Query().Get("hub.challenge")
        token := r.URL.Query().Get("hub.verify_token")
        if len(token) > 0 && len(challenge) > 0 && token == "my-token"{
            fmt.Fprintf(w, challenge)
            return
        }
    }
    w.WriteHeader(400)
    fmt.Fprintf(w, "Bad Request")
}