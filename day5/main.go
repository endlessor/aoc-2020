package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	f, _ := os.Open("./in")
	scanner := bufio.NewScanner(f)
	highest := 0

	seats := make(map[int]map[int]bool, 128)

	for scanner.Scan() {
		t := scanner.Text()

		r := t[0:7]
		s := t[7:10]

		x := (bin2dec(r, "B", 7) * 8) + bin2dec(s, "R", 3)
		if x > highest {
			highest = x
		}

		if seats[bin2dec(r, "B", 7)] == nil {
			seats[bin2dec(r, "B", 7)] = make(map[int]bool, 128)
		}
		seats[bin2dec(r, "B", 7)][bin2dec(s, "R", 3)] = true
	}

	for y, x := range seats {
		fmt.Println(y, x)
	}

	fmt.Printf("Highest, %d\n", highest)
}

func bin2dec(s, e string, a int) int {
	t := 0
	for _, x := range strings.Split(s, "") {
		if x == e {
			t += (int(math.Pow(2, float64(a))) / 2)
		}
		a--
	}

	return t
}
