package main

import "net/http"

func getRootURLHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("You are reading a simple text.\n"))
}

func getZoomURLHandler(w http.ResponseWriter, r *http.Request) {
	zoomPersonalMeetingURL := getZoomMeetingURL()
	http.Redirect(w, r, zoomPersonalMeetingURL, 302)
}
