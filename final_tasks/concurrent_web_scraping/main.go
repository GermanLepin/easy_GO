package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
)

// Create a program that concurrently scrapes data from multiple websites.
// The program should use goroutines to fetch data from different URLs simultaneously.
// Each goroutine should parse and process the HTML content to extract relevant information.
// Ensure proper synchronization and error handling. Display the collected data neatly.

func main() {
	var wg sync.WaitGroup

	urls := []string{
		"https://google.com",
		"https://linkedin.com",
		"https://github.com",
		"https://reddit.com",
	}

	countUrls := len(urls)

	results := make(chan string, countUrls)
	errors := make(chan error, countUrls)

	wg.Add(countUrls)
	for _, url := range urls {
		go fetchByUrl(&wg, url, results, errors)
	}

	wg.Wait()
	close(results)
	close(errors)

	wg.Add(1)
	go func() {
		defer wg.Done()

		for result := range results {
			fmt.Printf("received data from URL:\n%s\n", result)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		for err := range errors {
			fmt.Printf("error fetching data: %v\n", err)
		}
	}()

	wg.Wait()
	fmt.Println("about to exit")
}

func fetchByUrl(
	wg *sync.WaitGroup,
	url string,
	results chan<- string,
	errors chan<- error,
) {
	defer wg.Done()

	response, err := http.Get(url)
	if err != nil {
		errors <- err
		return
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		errors <- err
		return
	}

	results <- string(body)
}
