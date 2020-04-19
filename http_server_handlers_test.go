package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRootURLHandler(t *testing.T) {

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getRootURLHandler)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, rr.Result().StatusCode, 200)
}

func TestGetZoomURLHandler(t *testing.T) {

	req, err := http.NewRequest("GET", "/zoom", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getZoomURLHandler)
	handler.ServeHTTP(rr, req)

	redirectedURL := rr.Result().Header.Get("Location")

	assert.Equal(t, rr.Result().StatusCode, 302)
	assert.Equal(t, redirectedURL, "https://gojek.zoom.us/my/avirajkhare00")
}
