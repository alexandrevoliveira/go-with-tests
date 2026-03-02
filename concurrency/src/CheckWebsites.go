package concurrency

// In making CheckWebsites faster we learned about
// goroutines, the basic unit of concurrency in Go, which let us manage more than one website check request.
// anonymous functions, which we used to start each of the concurrent processes that check websites.
// channels, to help organize and control the communication between the different processes, allowing us to avoid a race condition bug.

type WebsiteChecker func(string) bool
type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultsChannel := make(chan result)

	for _, url := range urls {
		go func() {
			resultsChannel <- result{url, wc(url)}
		}()
	}

	for i := 0; i < len(urls); i++ {
		r := <-resultsChannel
		results[r.string] = r.bool
	}

	return results
}
