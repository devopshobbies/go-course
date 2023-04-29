package server

import (
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

	http.HandleFunc("/counter", s.incrementCounter)
	http.HandleFunc("/bmi", s.calculateBMI)

	log.Fatal(http.ListenAndServe(":8081", nil))
}
