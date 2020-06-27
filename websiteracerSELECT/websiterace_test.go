package websiteracerSELECT

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {

	t.Run("Fast servers beat slow ones", func(t *testing.T) {
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
		got, err := Racer(slowURL, fastURL)

		if err != nil{
			t.Fatalf("did not expect an error but got one %v", err)
		}

		if got != want {
			t.Errorf("Got %q, wanted %q", got, want)
		}

	})

	t.Run("Timeouts return an error", func(t *testing.T){
		serverA := makeDelayedServer(25*time.Millisecond)

		defer serverA.Close()

		// We can just use serverA.URL for both URL arguments
		// because we're not really interested in that part, just
		// that a timeout happens and is reported properly.
		_, err := ConfigurableRacer(serverA.URL, serverA.URL, 20*time.Millisecond)

		if err == nil{
			t.Error("expected error because of timeout but didn't get one")
		}
	})

}



func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
