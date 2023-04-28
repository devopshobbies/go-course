package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"
)

func main() {
	startTime := time.Now()
	jobs := make(chan string, 5)
	var wg sync.WaitGroup

	const workers = 3
	for worker := 0; worker < workers; worker++ {
		go getWebsite(jobs, &wg, worker)
	}

	file := "../websites.txt"
	reader, _ := os.Open(file)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		jobs <- scanner.Text()
		wg.Add(1)
	}

	wg.Wait()
	fmt.Println("totall time: ", time.Since(startTime))
}

// workers
func getWebsite(websites chan string, wg *sync.WaitGroup, worker int) {
	for website := range websites {
		if res, err := http.Get(website); err != nil {
			fmt.Printf("%s is down, proccessed by worker %d", website, worker)
		} else {
			fmt.Printf("[%d] %s is up, proccessed by worker %d\n", res.StatusCode, website, worker)
		}
		wg.Done()
	}
}
