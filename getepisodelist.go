package main

import (
	"log"

	"github.com/gocolly/colly/v2"
)

func getseasandeplist(showURL string) map[string]string {
	seasonandepmap := make(map[string]string)
	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36"),
	)

	// Debugging: Print visited URLs
	c.OnRequest(func(r *colly.Request) {
		// fmt.Println("Visiting:", r.URL)
	})

	// Error handling
	c.OnError(func(r *colly.Response, err error) {
		log.Printf("Error fetching %s: %v\n", r.Request.URL, err)
	})

	// First OnHTML: Extract season links
	c.OnHTML(`div.content-wrapper > section#single-diziler > div.full-width.pull-left > div.bg-dark.p-b-0 > a.btn.btn-s.btn-default-light`, func(e *colly.HTMLElement) {
		href := e.Attr("href")
		if href != "" {
			fullURL := e.Request.AbsoluteURL(href)
			// fmt.Println("Found season link:", fullURL)

			// Create a new request with context
			ctx := colly.NewContext()
			ctx.Put("isSeasonPage", "true")

			// Use Request.Visit with context
			err := c.Request("GET", fullURL, nil, ctx, nil)
			if err != nil {
				log.Printf("Failed to visit season page: %v\n", err)
			}
		}
	})

	// Second OnHTML: Process episode data (only for season pages)
	c.OnHTML(`article.grid-box.grid-four`, func(e *colly.HTMLElement) {
		// Check if this request is from a season page
		if e.Request.Ctx.Get("isSeasonPage") != "true" {
			return // Skip if not a season page
		}

		// Extract episode data
		postTitle := e.DOM.Find("div.post-title")
		title := postTitle.Find("a.season-episode").Text()
		link := postTitle.Find("a.season-episode").AttrOr("href", "")
		fullLink := e.Request.AbsoluteURL(link)
		// date := postTitle.Find("small.date").Text()
		seasonandepmap[title] = fullLink
		// fmt.Printf("Episode: %s\nLink: %s\nDate: %s\n\n", title, fullLink, date)
	})

	// Start scraping
	err := c.Visit(showURL)
	if err != nil {
		log.Fatal("Failed to start scraping:", err)
	}
	c.Wait()
	return seasonandepmap
}
