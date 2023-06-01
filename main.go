package main

import (
	"github.com/devopshobbies/go-course/server"
)

func main() {
	server := server.New()
	server.Serve(8080)
}
