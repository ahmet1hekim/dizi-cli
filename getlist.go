package main

import (
	"log"
	"strings"
	"time"

	"github.com/gocolly/colly/v2"
)

func getlist() map[string]string {
	TitletoURL := make(map[string]string)
	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36"),
	)

	c.Limit(&colly.LimitRule{
		DomainGlob:  "*dizibox.plus*",
		Parallelism: 1,
		Delay:       2 * time.Second,
	})

	selector := "ul.alphabetical-category-list"

	c.OnHTML(selector, func(e *colly.HTMLElement) {
		e.ForEach("li, il", func(_ int, el *colly.HTMLElement) {
			// Extract from both valid <li> and typo <il> elements
			link := el.ChildAttr("a", "href")
			title := el.ChildAttr("a", "title")

			// Clean and validate data
			if link != "" && title != "" {
				fullURL := e.Request.AbsoluteURL(link)
				cleanTitle := strings.TrimSpace(title)
				TitletoURL[cleanTitle] = fullURL

				// fmt.Printf("Title: %s\nURL: %s\n\n", cleanTitle, fullURL)
			}
		})

	})

	// Error handling
	c.OnError(func(r *colly.Response, err error) {
		log.Printf("Error fetching %s: %v", r.Request.URL, err)
	})

	// Start scraping
	err := c.Visit("https://www.dizibox.plus/")
	if err != nil {
		log.Fatal("Visit failed:", err)
	}
	return TitletoURL

}
