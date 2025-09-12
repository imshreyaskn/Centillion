package main

import (
	"fmt"
	"sync"
	"time"
)

func fetchProfile(userID string, ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Fetching Profile for user %s...\n", userID)
	time.Sleep(3 * time.Second)
	profile := fmt.Sprintf("Profile for %s\n", userID)
	ch <- profile

}

func fetchPosts(userID string, ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Fetching posts for user %s...\n", userID)
	time.Sleep(2 * time.Second)
	posts := fmt.Sprintf("Posts from %s: [Post1, Post2, Post3]\n", userID)
	ch <- posts
}

func main() {

	var wg sync.WaitGroup
	wg.Add(2)

	ch := make(chan string, 2)

	startTime := time.Now()
	userId := "userId123"

	go fetchProfile(userId, ch, &wg)
	go fetchPosts(userId, ch, &wg)

	go func() {
		wg.Wait()
		close(ch)
	}()

	for result := range ch {
		fmt.Println(result)
	}

	fmt.Printf("\nTotal Time: %v\n", time.Since(startTime))

}
