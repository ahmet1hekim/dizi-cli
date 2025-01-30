package main

import (
	"context"
	"strings"
	"sync"
	"time"

	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
)

func getURL(epURL string) string {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	// Configure Chrome options (headless-friendly)
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", true), // Visual debug
		chromedp.Flag("disable-web-security", true),
		chromedp.Flag("disable-blink-features", "AutomationControlled"),
		chromedp.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36"),
	)
	allocCtx, cancel := chromedp.NewExecAllocator(ctx, opts...)
	defer cancel()
	ctx, cancel = chromedp.NewContext(allocCtx)
	defer cancel()

	// Channel to signal URL detection
	urlFound := make(chan struct{})
	var targetURL string
	var mu sync.Mutex
	var once sync.Once
	// Listen for network events
	// chromedp.ListenTarget(ctx, func(ev interface{}) {
	// 	switch ev := ev.(type) {
	// 	case *network.EventRequestWillBeSent:
	// 		// Match URL pattern (adjust as needed)
	// 		if strings.Contains(ev.Request.URL, "vidmoly") {
	// 			log.Printf("[FOUND] Subdocument Request: %s", ev.Request.URL)
	// 			targetRequests = append(targetRequests, ev.Request.URL)
	// 			// Optional: Print headers or other details
	// 			// log.Printf("Headers: %v", ev.Request.Headers)
	// 		}
	// 	}
	// })
	chromedp.ListenTarget(ctx, func(ev interface{}) {
		switch ev := ev.(type) {
		case *network.EventRequestWillBeSent:
			// Check if the request matches the target URL
			if strings.Contains(ev.Request.URL, "vidmoly") {
				mu.Lock()
				targetURL = ev.Request.URL
				mu.Unlock()
				once.Do(func() {
					close(urlFound)
				}) // Signal that the URL is found
			}
		}
	})

	// Tasks to run
	tasks := chromedp.Tasks{
		network.Enable(),
		network.SetCacheDisabled(true),
		chromedp.Navigate(epURL), // Replace with your URL
		// Wait for the URL to be detected (or timeout)
		chromedp.ActionFunc(func(ctx context.Context) error {
			select {
			case <-urlFound:
				// log.Printf("Target URL detected: %s", targetURL)
				return nil
			case <-ctx.Done():
				return ctx.Err()
			}
		}),
	}

	// Execute tasks
	err := chromedp.Run(ctx, tasks...)
	if err != nil {
		// log.Fatalf("Failed: %v", err)
		return "error geting epURL"
	} else {
		return targetURL
	}
}
