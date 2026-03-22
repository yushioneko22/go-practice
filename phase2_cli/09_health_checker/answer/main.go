package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func checkSite(url string, wg *sync.WaitGroup) {
	defer wg.Done()

	client := http.Client{Timeout: 3 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		fmt.Printf("[DOWN] %s (Error: %v)\n", url, err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		fmt.Printf("[ UP ] %s (%d)\n", url, resp.StatusCode)
	} else {
		fmt.Printf("[ERRO] %s (%d)\n", url, resp.StatusCode)
	}
}

func main() {
	urls := []string{
		"https://google.com",
		"https://github.com",
		"https://invalid-url-example.com",
	}

	var wg sync.WaitGroup
	fmt.Println("--- Site Health Check Start ---")
	for _, u := range urls {
		wg.Add(1)
		go checkSite(u, &wg)
	}
	wg.Wait()
	fmt.Println("--- Site Health Check Finished ---")
}
