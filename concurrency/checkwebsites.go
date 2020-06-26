package concurrency

import "time"

type WebsiteChecker func(string) bool

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)

	for _, url := range urls {
		// To tell Go to start a new goroutine we turn a function call into a go
		// statement by putting the keyword go in front of it: go doSomething().
		// Here, we're doing that with an anonymous function before we then perform
		// the checking.
		go func(u string) {
		results[u] = wc(u)
		}(url)
		// The () means the function is executed at the SAME time
		// that it is declared. All the variables that are available at the
		// point when you declare the anonymous function are also available
		// in the body of the function. (so, that's results, url and wc)

	}

	// Essentially give the goroutine some time.
	time.Sleep(2 *time.Second)

	return results
}