package main

import (
	"fmt"
	"sync"
)

type webFinder struct {
	urls map[string]bool
	mx   sync.Mutex
}

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

func (find *webFinder) IsThere(url string) (isThere bool) {
	find.mx.Lock()
	defer find.mx.Unlock()

	_, isThere = find.urls[url]
	return
}

func (find *webFinder) Insert(url string) {
	find.mx.Lock()
	defer find.mx.Unlock()

	if find.urls == nil {
		find.urls = make(map[string]bool)
	}
	find.urls[url] = true
	return
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func (find *webFinder) Crawl(url string, depth int, fetcher Fetcher, wg *sync.WaitGroup) {
	defer wg.Done()
	if depth <= 0 || find.IsThere(url) {
		return
	}
	find.Insert(url)

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		wg.Add(1)
		go find.Crawl(u, depth-1, fetcher, wg)
	}
	return
}

func crawl() {
	var wg sync.WaitGroup // wait function call was required so that when all the children have completed then only the parent can exit()
	// as in go is a thread call so not wait only WaitGroup is there

	wg.Add(1)
	webfinder := &webFinder{urls: make(map[string]bool)}
	go webfinder.Crawl("https://golang.org/", 4, fetcher, &wg)
	wg.Wait() // all the thread have done then exit out
}

func main() {
	crawl()
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
