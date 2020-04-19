package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//TestgetZoomMeetingURL tests getZoomMeetingURL()
func TestGetZoomMeetingURL(t *testing.T) {
	zoomMeetingURL := getZoomMeetingURL()

	assert.Equal(t, zoomMeetingURL, "https://gojek.zoom.us/my/avirajkhare00", "Status code of / should be 200")
}
