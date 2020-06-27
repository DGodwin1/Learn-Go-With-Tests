package websiteracerSELECT

import (
	"fmt"
	"net/http"
	"time"
)

var tenSecondTimeout = 10 * time.Second

func Racer(url1, url2 string) (winner string, error error) {
	return ConfigurableRacer(url1, url2, tenSecondTimeout)
}

func ConfigurableRacer(url1, url2 string, timeout time.Duration) (fastURL string, err error){
	select {
	case <-ping(url1):
		return url1, nil
	case <-ping(url2):
		return url2, nil
		// If neither url1 not url2 come back before 10 seconds,
		// then a timeout error will be thrown instead.
		// Adding timeout to the duration means we get away
		// from hard coding that value. Instead, we can set different
		// timeouts depending on whether we're running in production or test.
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", url1, url2)
	}}

func ping(url string) chan struct{}{
	ch := make(chan struct{})
	go func(){
		http.Get(url)
		close(ch)
	}()
	return ch
}

