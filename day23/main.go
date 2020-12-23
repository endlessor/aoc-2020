package main

import (
	"bufio"
	"container/ring"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	filename    = "./in"
	secondMoves = 10000000
	firstMoves  = 100
)

func main() {
	first()
	second()
}

func first() {
	input := []int{}

	f, _ := os.Open(filename)
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		t := scanner.Text()

		x := strings.Split(t, "")
		for _, a := range x {
			n, _ := strconv.Atoi(a)

			input = append(input, n)
		}
	}

	run(input, false)
}

func second() {
	input := []int{}

	f, _ := os.Open(filename)
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		t := scanner.Text()

		x := strings.Split(t, "")
		for _, a := range x {
			n, _ := strconv.Atoi(a)

			input = append(input, n)
		}
	}

	for x := 10; x <= 1000000; x++ {
		input = append(input, x)
	}

	run(input, true)
}

func run(input []int, doSecond bool) {
	moves := 0
	if doSecond {
		moves = secondMoves
	} else {
		moves = firstMoves
	}

	totalSize := len(input)
	cache := make(map[int]*ring.Ring, totalSize)
	cups := ring.New(totalSize)

	for _, c := range input {
		cups.Value = c
		cache[c] = cups
		cups = cups.Next()
	}

	current := cups
	for loop := 0; loop < moves; loop++ {
		three := current.Unlink(3)
		next := 1 + ((totalSize + current.Value.(int) - 2) % totalSize)
		removed := make(map[int]bool)
		for x := 0; x < 3; x++ {
			removed[three.Move(x).Value.(int)] = true
		}
		for removed[next] {
			next = 1 + ((totalSize + next - 2) % totalSize)
		}
		cache[next].Link(three)
		current = current.Next()
	}

	one := cache[1]
	if doSecond {
		fmt.Println("#2:", one.Move(1).Value.(int)*one.Move(2).Value.(int))
	} else {
		fmt.Printf("#1: ")
		for x := 1; x < totalSize; x++ {
			one = one.Next()
			fmt.Printf("%d ", one.Value.(int))
		}
		fmt.Println("")
	}
}
