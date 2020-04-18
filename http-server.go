package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("You are reading a simple text.\n"))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
