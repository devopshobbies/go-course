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
	var wg sync.WaitGroup

	file := "../websites.txt"
	reader, _ := os.Open(file)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		go getWebsite(scanner.Text(), &wg)
		wg.Add(1)
	}

	wg.Wait()
	fmt.Println("totall time: ", time.Since(startTime))
}

func getWebsite(website string, wg *sync.WaitGroup) {
	defer wg.Done()
	if res, err := http.Get(website); err != nil {
		fmt.Println(website, "is down")
	} else {
		fmt.Printf("[%d] %s is up\n", res.StatusCode, website)
	}
}
