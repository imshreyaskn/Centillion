package main

import (
	"fmt"
	"net/http"
	"time"
)

func fetchURL(url string) (int, error) {

	client := &http.Client{
		Timeout: 2 * time.Second,
	}

	resp, err := client.Get(url)
	if err != nil {
		return 0, err
	}

	defer resp.Body.Close()

	return resp.StatusCode, nil

}

func main() {
	urls := []string{
		"https://google.com",
		"https://golang.com",
		"https://yahoo.com",
	}

	results := map[string]int{}

	for _, url := range urls {
		statusCode, err := fetchURL(url)
		if err != nil {
			fmt.Println("Error")
			results[url] = 0
		} else {
			results[url] = statusCode
		}
	}

	fmt.Println("Result:")

	for url, code := range results {
		fmt.Printf("%s -> %d \n", url, code)
	}

}
