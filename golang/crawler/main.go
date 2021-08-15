package main

import (
	c "golang/crawler/crawler"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	c.CrawlerFilm(&wg)
	// c.CrawlerProduct(&wg)
}