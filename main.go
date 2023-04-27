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
	curser := &element{data: '|'}

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
			if curser.prev != nil {
				element.prev = curser.prev
				curser.prev.next = element
			}
			curser.prev = element
			element.next = curser

		case LEFT:
			if curser.prev != nil {
				curser.prev.data, curser.data = curser.data, curser.prev.data
				curser = curser.prev
			}

		case RIGHT:
			if curser.next != nil {
				curser.next.data, curser.data = curser.data, curser.next.data
				curser = curser.next
			}

		case BACKSPACE:
			if curser.prev != nil {
				if curser.prev.prev != nil {
					curser.prev.prev.next = curser
					curser.prev = curser.prev.prev
				} else {
					curser.prev = nil
				}
			}
		default:
			panic("invalid command")
		}
	}

	curser.print()
}
