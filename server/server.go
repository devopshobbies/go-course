package server

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

type server struct {
	counter int
	mutex   *sync.Mutex
}

func New() *server {
	return &server{
		mutex:   &sync.Mutex{},
		counter: 0,
	}
}

func (s *server) Serve() {
	http.Handle("/", http.FileServer(http.Dir("./static")))

	http.HandleFunc("/hello", s.incrementCounter)
	http.HandleFunc("/counter", s.incrementCounter)
	http.HandleFunc("/bmi", s.calculateBMI)

	http.Handle("/sample", &sample{})

	log.Fatal(http.ListenAndServe(":8081", nil))
}

type sample struct {
}

func (s *sample) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Sample handler")
}
