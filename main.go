package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type command string

const (
	INSERT    command = "INSERT"
	LEFT      command = "LEFT"
	RIGHT     command = "RIGHT"
	BACKSPACE command = "BACKSPACE"
)

type element struct {
	data byte
	prev *element
	next *element
}

func (e *element) getFirstElement() (current *element) {
	current = e

	for {
		if current.prev == nil {
			break
		}
		current = current.prev
	}

	return
}

func (e *element) print() {
	current := e.getFirstElement()
	var sb strings.Builder

	for {
		sb.WriteByte(current.data)
		if current.next == nil {
			break
		} else {
			current = current.next
		}
	}

	fmt.Println(sb.String())
}

func main() {
	pointer := &element{data: '|'}

	var count int
	fmt.Scan(&count)

	reader := bufio.NewReader(os.Stdin)

	for index := 0; index < count; index++ {
		line, _ := reader.ReadString('\n')
		line = strings.TrimSuffix(line, "\n")
		fields := strings.Fields(line)

		switch command(fields[0]) {
		case INSERT:
			element := &element{data: []byte(fields[1])[0]}
			if pointer.prev != nil {
				element.prev = pointer.prev
				pointer.prev.next = element
			}
			pointer.prev = element
			element.next = pointer

		case LEFT:
			if pointer.prev != nil {
				pointer.prev.data, pointer.data = pointer.data, pointer.prev.data
				pointer = pointer.prev
			}

		case RIGHT:
			if pointer.next != nil {
				pointer.next.data, pointer.data = pointer.data, pointer.next.data
				pointer = pointer.next
			}

		case BACKSPACE:
			if pointer.prev != nil {
				if pointer.prev.prev != nil {
					pointer.prev.prev.next = pointer
					pointer.prev = pointer.prev.prev
				} else {
					pointer.prev = nil
				}
			}
		default:
			panic("invalid command")
		}
	}

	pointer.print()
}
