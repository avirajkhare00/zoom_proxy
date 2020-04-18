package main

import (
	"log"
	"net/http"
)

func main() {

	zoomPersonalMeetingURL := "https://gojek.zoom.us/my/avirajkhare00"

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("You are reading a simple text.\n"))
	})

	http.HandleFunc("/zoom", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, zoomPersonalMeetingURL, 302)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
