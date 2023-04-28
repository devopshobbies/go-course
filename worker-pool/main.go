package main

import (
	"fmt"
	"net/http"
	"time"
)

var websites = []string{
	"https://google.com/",
	"https://github.com/",
	"https://www.linkedin.com/",
	"http://medium.com/",
	"https://golang.org/",
	"https://www.udemy.com/",
	"https://www.coursera.org/",
	"https://wesionary.team/",
}

func main() {
	startTime := time.Now()
	jobs := make(chan string, 5)
	counter := make(chan struct{}, len(websites))

	const workers = 3
	for worker := 0; worker < workers; worker++ {
		go getWebsite(jobs, counter, worker)
	}

	// populate jobs
	for index := 0; index < len(websites); index++ {
		jobs <- websites[index]
	}

	// sleep until all the jobs are finished
	// can also be implemented by wait-groups
	for index := 0; index < len(websites); index++ {
		<-counter
	}

	fmt.Println("totall time: ", time.Since(startTime))
}

// workers
func getWebsite(websites chan string, counter chan struct{}, worker int) {
	for website := range websites {
		if res, err := http.Get(website); err != nil {
			fmt.Printf("%s is down, proccessed by worker %d", website, worker)
		} else {
			fmt.Printf("[%d] %s is up, proccessed by worker %d\n", res.StatusCode, website, worker)
		}
		counter <- struct{}{}
	}
}
