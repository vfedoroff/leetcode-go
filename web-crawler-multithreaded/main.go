package main

import (
	"fmt"
	"strings"
	"sync"
)

type HtmlParser interface {
	GetUrls(url string) []string
}

type HtmlParserImpl struct {
	URLS []string
}

func (h *HtmlParserImpl) GetUrls(url string) []string {
	return h.URLS
}

func crawl(url string, htmlParser HtmlParser) []string {
	visited := map[string]struct{}{}
	queue := make(chan string, 10)
	prefix := strings.Split(url, "/")[2]

	queue <- url
	visited[url] = struct{}{}

	visitedMutex := sync.Mutex{}

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {

			if len(queue) == 0 {
				return
			}
			current := <-queue
			urls := htmlParser.GetUrls(current)
			for _, url := range urls {
				if strings.Split(url, "/")[2] != prefix {
					continue
				}
				visitedMutex.Lock()
				if _, ok := visited[url]; !ok {
					queue <- url
					visited[url] = struct{}{}
				}
				visitedMutex.Unlock()
			}
		}
	}()

	wg.Wait()

	result := []string{}
	for k := range visited {
		result = append(result, k)
	}

	return result
}

func main() {
	res := crawl("http://news.yahoo.com/news/topics/", &HtmlParserImpl{URLS: []string{
		"http://news.yahoo.com",
		"http://news.yahoo.com/news",
		"http://news.yahoo.com/news/topics/",
		"http://news.google.com",
		"http://news.yahoo.com/us",
	}})
	fmt.Printf("%v\n", res)
}
