package main

import (
	"bufio"
	"fmt"
	"os"
)

func d(o, l int) int {
	offset := o
	tree := 0
	f, _ := os.Open("./in")
	scanner := bufio.NewScanner(f)

	scanner.Scan()
	for scanner.Scan() {
		if l == 2 {
			scanner.Scan()
		}

		t := scanner.Text()
		if o >= len(t) {
			o -= len(t)
		}

		if t[o] == byte('#') {
			tree++
		}

		o += offset
	}

	return tree
}

//
func main() {
	offset := 3
	tree := 0
	f, _ := os.Open("./in")
	scanner := bufio.NewScanner(f)

	// skip the first line
	scanner.Scan()

	for scanner.Scan() {
		t := scanner.Text()
		if offset >= len(t) {
			offset = offset - len(t)
		}

		if t[offset] == byte('#') {
			tree++
		}

		offset += 3
	}

	fmt.Println("Found,", d(1, 1))
	fmt.Println("Found,", d(3, 1))
	fmt.Println("Found,", d(5, 1))
	fmt.Println("Found,", d(7, 1))
	fmt.Println("Found,", d(1, 2))

	fmt.Println("Big,", d(1, 1)*d(3, 1)*d(5, 1)*d(7, 1)*d(1, 2))
}
