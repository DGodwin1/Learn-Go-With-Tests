package concurrency

type WebsiteChecker func(string) bool
type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		// To tell Go to start a new goroutine we turn a function call into a go
		// statement by putting the keyword go in front of it: go doSomething().
		// Here, we're doing that with an anonymous function before we then perform
		// the checking.
		go func(u string) {
			resultChannel <- result{u, wc(u)} //adding a url and bool
		}(url)
	}
		// The () means the function is executed at the SAME time
		// that it is declared. All the variables that are available at the
		// point when you declare the anonymous function are also available
		// in the body of the function. (so, that's results, url and wc)
		for i := 0; i < len(urls); i++ {
			// Get a value out of the resultChannel and assign to result.

			result := <-resultChannel
			// Now get info out of the struct with dot notation
			// And place into the results map.
			results[result.string] = result.bool
		}

	return results
}
