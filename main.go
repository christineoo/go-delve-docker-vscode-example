package main

import (
	"log"
	"net"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		message := r.URL.Path
		message = strings.TrimPrefix(message, "/")
		message = "Hello, " + message + "!"

		w.Write([]byte(message))
	})

	log.Print("starting web server")
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Start listening: %v", listener.Addr().String())

	if err := http.Serve(listener, nil); err != nil {
		log.Fatal(err)
	}
}