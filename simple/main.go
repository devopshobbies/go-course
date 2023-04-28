package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	startTime := time.Now()

	file := "../websites.txt"
	reader, _ := os.Open(file)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		getWebsite(scanner.Text())
	}

	fmt.Println("totall time: ", time.Since(startTime))
}

func getWebsite(website string) {
	if res, err := http.Get(website); err != nil {
		fmt.Println(website, "is down")
	} else {
		fmt.Printf("[%d] %s is up\n", res.StatusCode, website)
	}
}
