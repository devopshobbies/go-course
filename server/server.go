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

func (s *server) Serve(port int) {
	http.Handle("/", http.FileServer(http.Dir("./static")))

	http.HandleFunc("/hello", s.incrementCounter)
	http.HandleFunc("/counter", s.incrementCounter)
	http.HandleFunc("/bmi", s.calculateBMI)

	http.Handle("/sample", &sample{})

	addr := fmt.Sprintf(":%d", port)
	fmt.Println("listening on", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

type sample struct {
}

func (s *sample) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Sample handler")
}
