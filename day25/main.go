package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	con     = 20201227
	subject = 7
)

func main() {
	f, _ := os.Open("./in")
	scanner := bufio.NewScanner(f)

	scanner.Scan()
	t := scanner.Text()
	cardPublic, _ := strconv.Atoi(t)

	scanner.Scan()
	t = scanner.Text()
	doorPublic, _ := strconv.Atoi(t)

	cardLoop := 0
	doorLoop := 0

	val := 1
	for loop := 1; loop > 0; loop++ {
		newVal := subject * val
		val = newVal % con

		if val == cardPublic {
			cardLoop = loop
			if doorLoop > 0 {
				break
			}
		}
		if val == doorPublic {
			doorLoop = loop
			if cardLoop > 0 {
				break
			}
		}
	}

	private := 1
	for loop := 1; loop <= doorLoop; loop++ {
		newPrivate := cardPublic * private
		private = newPrivate % con
	}

	fmt.Println("#1:", private)
}
