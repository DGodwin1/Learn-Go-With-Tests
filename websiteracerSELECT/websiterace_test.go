package websiteracerSELECT

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	// A server takes a ResponseWriter and a Request.
	slowServer := makeDelayedServer(5 * time.Millisecond)
	fastServer := makeDelayedServer(0)

	defer slowServer.Close()
	defer fastServer.Close()

	// You're setting up the slowURL to use the slow server
	// and the fastURL to use the fast one. http.GET will call
	// on the server you've put here.
	slowURL := slowServer.URL
	fastURL := fastServer.URL

	want := fastURL
	got := Racer(slowURL, fastURL)

	if got != want {
		t.Errorf("Got %q, wanted %q", got, want)
	}

}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
