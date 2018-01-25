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

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, syncF syncFetcher) {
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:
	if depth <= 0 || syncF.cache[url] {
		return
	}
	
	syncF.mux.Lock()
	body, urls, err := syncF.f.Fetch(url)
	syncF.cache[url] = true
	syncF.mux.Unlock()
	
	if err != nil {
		fmt.Println(err)
		return
	}
	
	fmt.Printf("found: %s %q\n", url, body)
	
	done := make(chan bool)
	for _, u := range urls {
		go func(url string) {
			Crawl(url, depth-1, syncF)
			done <- true
		}(u)
	}
	for range urls {
		<- done
	}
	return
}

func main() {
	Crawl("http://golang.org/", 4, syncFet)
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type syncFetcher struct {
	f fakeFetcher
	cache map[string]bool
	mux sync.Mutex
}

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	res, ok := f[url];
	if ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"http://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}

var syncFet = syncFetcher{f: fetcher, cache: make(map[string]bool)}

