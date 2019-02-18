package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

var cached_urls = make(map[string]bool)
var cached_urls_mutex sync.Mutex

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, ch chan error) {
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:
	if depth <= 0 {
		ch <- fmt.Errorf("too deep")
		return
	}

	cached_urls_mutex.Lock()
	_, cached := cached_urls[url]
	if cached {
		cached_urls_mutex.Unlock()
		ch <- fmt.Errorf("already fetched: %v", url)
		return
	}

	cached_urls[url] = true
	cached_urls_mutex.Unlock()

	fmt.Printf("fetch: %v\n", url)
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		ch <- err
		return
	}

	fmt.Printf("found: %s %q\n", url, body)
	ch_to_child := make(chan error)
	for _, u := range urls {
		go Crawl(u, depth-1, fetcher, ch_to_child)
	}
	for i := 0; i < len(urls); i++ {
		_ = <-ch_to_child
	}

	ch <- nil
	return
}

func main() {
	ch := make(chan error)
	go Crawl("https://golang.org/", 4, fetcher, ch)
	_ = <-ch
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
