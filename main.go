package main

import (
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", getRootURLHandler)
	http.HandleFunc("/zoom", getZoomURLHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
