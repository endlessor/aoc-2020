package main

import (
	"fmt"
)

func main() {
	first("1", 2020)
	first("2", 30000000)
}

var (
	all      []int
	lastNum  int
	lastTurn map[int]int
)

func first(round string, max int) {
	lastTurn = make(map[int]int)

	all = []int{9, 3, 1, 0, 8, 4}
	for a, b := range all {
		lastTurn[b] = a
		lastNum = b
	}

	for x := len(all); x < max; x++ {
		var newNum int

		if prev, seen := lastTurn[lastNum]; seen {
			newNum = x - 1 - prev
		} else {
			newNum = 0
		}

		if lastNum >= 0 {
			lastTurn[lastNum] = x - 1
		}

		lastNum = newNum
	}

	fmt.Printf("#%s: %d\n", round, lastNum)
}
