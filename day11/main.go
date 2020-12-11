package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

const (
	emptySeat    = "L"
	floor        = "."
	max          = 4 // this was 3 for question 1
	occupiedSeat = "#"
)

func main() {
	f, _ := os.Open("./in")
	scanner := bufio.NewScanner(f)

	a := [][]string{}
	for scanner.Scan() {
		t := scanner.Text()

		a = append(a, strings.Split(t, ""))
	}

	b := cop(a)

	loop := 0
	count := 0
	for {
		for x, _ := range a {
			for y, z := range a[x] {
				if z == floor {
					continue
				}

				adj := countAdjacent(a, x, y)
				if adj == 0 {
					b[x][y] = occupiedSeat
				} else if adj > max {
					b[x][y] = emptySeat
				}
			}
		}

		loop++
		if occu(b) == count {
			fmt.Println("Occupied seats,", count)
			break
		}
		count = occu(b)
		a = cop(b)

		fmt.Println("\n\n\n")
		p(a)
		time.Sleep(time.Second / 2)
	}

	fmt.Println("Loops,", loop)
}

func p(s [][]string) {
	for _, x := range s {
		for _, y := range x {
			fmt.Printf("%s", y)
		}
		fmt.Println("")
	}
}

func occu(s [][]string) int {
	c := 0
	for _, x := range s {
		for _, y := range x {
			if y == occupiedSeat {
				c++
			}
		}
	}

	return c
}

// this should be replaced into countAdjacent to run problem 1
func countAdjacentFirst(s [][]string, x, y int) int {
	c := 0
	if g(s, x-1, y-1) {
		c++
	}
	if g(s, x, y-1) {
		c++
	}
	if g(s, x+1, y-1) {
		c++
	}
	if g(s, x-1, y) {
		c++
	}
	if g(s, x+1, y) {
		c++
	}
	if g(s, x-1, y+1) {
		c++
	}
	if g(s, x, y+1) {
		c++
	}
	if g(s, x+1, y+1) {
		c++
	}

	return c
}

func countAdjacent(s [][]string, x, y int) int {
	c := 0
	if g2(s, x-1, y-1, 0) {
		c++
	}
	if g2(s, x, y-1, 1) {
		c++
	}
	if g2(s, x+1, y-1, 2) {
		c++
	}
	if g2(s, x-1, y, 3) {
		c++
	}
	if g2(s, x+1, y, 4) {
		c++
	}
	if g2(s, x-1, y+1, 5) {
		c++
	}
	if g2(s, x, y+1, 6) {
		c++
	}
	if g2(s, x+1, y+1, 7) {
		c++
	}

	return c
}

func g2(s [][]string, x, y, d int) bool {
	if x < 0 || y < 0 || x >= len(s) || y >= len(s[x]) {
		return false
	}

	if s[x][y] == occupiedSeat {
		return true
	}

	if s[x][y] == floor {
		switch d {
		case 0:
			return g2(s, x-1, y-1, 0)
		case 1:
			return g2(s, x, y-1, 1)
		case 2:
			return g2(s, x+1, y-1, 2)
		case 3:
			return g2(s, x-1, y, 3)
		case 4:
			return g2(s, x+1, y, 4)
		case 5:
			return g2(s, x-1, y+1, 5)
		case 6:
			return g2(s, x, y+1, 6)
		case 7:
			return g2(s, x+1, y+1, 7)
		}
	}

	return false
}

func g(s [][]string, x, y int) bool {
	if x < 0 || y < 0 || x >= len(s) || y >= len(s[x]) {
		return false
	}

	if s[x][y] == occupiedSeat {
		return true
	}
	return false
}

func cop(a [][]string) [][]string {
	b := make([][]string, len(a))
	for x, _ := range a {
		b[x] = make([]string, len(a[x]))
		for y, z := range a[x] {
			b[x][y] = z
		}
	}
	return b
}
