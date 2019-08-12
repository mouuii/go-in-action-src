package search

import (
	"log"
	"sync"
)

var matchers = make(map[string]Matcher)

func Run(searchTerm string) {
	feeds, err := RetriveFeeds()
	if err != nil {
		log.Fatal(err)
	}
	results := make(chan *Result)
	var wg sync.WaitGroup

	wg.Add(len(feeds))

	for _, feed := range feeds {
		matcher, exists := matchers[feed.Type]
		if !exists {
			matcher = matchers["default"]
		}
		go func(matcher Matcher, feed *Feed) {
			Match(matcher, feed, searchTerm, results)
			wg.Done()
		}(matcher, feed)
	}
	go func() {
		wg.Wait()
		close(results)
	}()
	display(results)

}

func Register(feedType string, matcher Matcher) {
	if _, exists := matchers[feedType]; exists {
		log.Fatal(feedType, "Matcher already registered")
	}
	log.Println("register", feedType, "matcher")
	matchers[feedType] = matcher
}

func display(results chan *Result) {
	for result := range results {
		log.Printf("%s-------%s:\n", result.Field, result.Content)
	}
}
