package main

import (
	"fmt"
	"sync"
)

type Visit struct {
	mu      sync.Mutex
	visited map[string]bool
}

type Fetcher interface {
	// Fetch возвращает тело страницы по URL
	// и срез URL-адресов, найденных на этой странице.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl использует fetcher для рекурсивного обхода
// страниц, начиная с указанного URL, до максимальной глубины depth.
func Crawl(url string, depth int, fetcher Fetcher, state *Visit, wg *sync.WaitGroup) {
	// TODO: Загружать URL-адреса параллельно.
	// TODO: Не загружать один и тот же URL дважды.
	// Эта реализация пока не делает ни того, ни другого:
	defer wg.Done()
	if depth <= 0 {
		return
	}

	state.mu.Lock()
	if state.visited[url] == true {
		state.mu.Unlock()
		return
	} else {
		state.visited[url] = true
		state.mu.Unlock()
	}

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		wg.Add(1)
		go Crawl(u, depth-1, fetcher, state, wg)
	}
	return
}

func main() {
	state := Visit{visited: make(map[string]bool)}
	var wg sync.WaitGroup
	wg.Add(1)
	go Crawl("https://golang.org/", 4, fetcher, &state, &wg)
	wg.Wait()
}

// fakeFetcher — это реализация Fetcher, которая возвращает заранее заданные результаты.
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

// fetcher — это заполненный fakeFetcher с тестовыми данными.
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
