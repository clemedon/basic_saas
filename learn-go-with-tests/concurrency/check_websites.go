package concurrency

// WebsiteChecker checks a url, returning a bool.
type WebsiteChecker func(string) bool
type result struct { // has ANONYMOUS MEMBERS
	string
	bool
}

// CheckWebsites takes a WebsiteChecker and a slice of urls and returns  a map.
// of urls to the result of checking each url with the WebsiteChecker function.
func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		go func(u string) { // GOROUTINES of ANONYMOUS FUNCTION
			resultChannel <- result{u, wc(u)} // CHANNEL SEND
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		r := <-resultChannel // CHANNEL RECEIVE
		results[r.string] = r.bool
	}

	return results
}