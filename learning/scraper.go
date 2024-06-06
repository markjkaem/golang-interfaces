package learning

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"sync"
	"time"
)

// Scraper defines the interface for scraping web pages
type Scraper interface {
	Scrape(ctx context.Context, url string) ([]byte, error)
}

// SimpleScraper implements the Scraper interface
type SimpleScraper struct {
	Client *http.Client
}

// NewSimpleScraper creates a new SimpleScraper
func NewSimpleScraper(timeout time.Duration) *SimpleScraper {
	return &SimpleScraper{
		Client: &http.Client{Timeout: timeout},
	}
}

// Scrape fetches the contents of a URL
func (s *SimpleScraper) Scrape(ctx context.Context, url string) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := s.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch url %s: %w", url, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("bad status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	return body, nil
}

// Result holds the result of a scraping operation
type Result struct {
	URL  string
	Data []byte
	Err  error
}

// Worker is a function that processes a single URL
func Worker(ctx context.Context, scraper Scraper, url string, results chan<- Result) {
	data, err := scraper.Scrape(ctx, url)
	results <- Result{URL: url, Data: data, Err: err}
}

// ConcurrentScraper manages concurrent scraping of multiple URLs
type ConcurrentScraper struct {
	Scraper    Scraper
	NumWorkers int
}

// NewConcurrentScraper creates a new ConcurrentScraper
func NewConcurrentScraper(scraper Scraper, numWorkers int) *ConcurrentScraper {
	return &ConcurrentScraper{Scraper: scraper, NumWorkers: numWorkers}
}

// Scrape concurrently scrapes multiple URLs
func (c *ConcurrentScraper) Scrape(ctx context.Context, urls []string) []Result {
	var wg sync.WaitGroup
	results := make(chan Result, len(urls))

	semaphore := make(chan struct{}, c.NumWorkers)
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			semaphore <- struct{}{}
			defer func() { <-semaphore }()
			Worker(ctx, c.Scraper, url, results)
		}(url)
	}

	wg.Wait()
	close(results)

	var finalResults []Result
	for result := range results {
		finalResults = append(finalResults, result)
	}

	return finalResults
}

// processResults processes the results of scraping
func ProcessResults(results []Result) {
	for _, result := range results {
		if result.Err != nil {
			fmt.Printf("Failed to fetch %s: %v\n", result.URL, result.Err)
			continue
		}
		fmt.Printf("Fetched %d bytes from %s\n", len(result.Data), result.URL)
	}
}

func ScraperExec() {
	urls := []string{
		"https://example.com",
		"https://golang.org",
		"https://github.com",
	}

	scraper := NewSimpleScraper(10 * time.Second)
	concurrentScraper := NewConcurrentScraper(scraper, 5)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	results := concurrentScraper.Scrape(ctx, urls)
	ProcessResults(results)
}
