package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	first()
	second()
}

func first() {
	f, _ := os.Open("./in")
	scanner := bufio.NewScanner(f)

	earliest := 0
	eb := 0
	ebDiff := 10000

	for scanner.Scan() {
		t := scanner.Text()

		if earliest == 0 {
			earliest, _ = strconv.Atoi(t)
		} else {
			x := strings.Split(t, ",")
			for _, y := range x {
				if y == "x" {
					continue
				}

				n, _ := strconv.Atoi(y)
				a := (-1 * (earliest % n)) + n
				if a < ebDiff {
					ebDiff = a
					eb = n
				}
			}
		}
	}

	fmt.Println("#1:", eb*ebDiff)
}

func second() {
	f, _ := os.Open("./in")
	scanner := bufio.NewScanner(f)

	sched := map[int]int{}

	scanner.Scan()
	scanner.Scan()
	t := scanner.Text()

	for idx, x := range strings.Split(t, ",") {
		if x == "x" {
			continue
		}

		n, _ := strconv.Atoi(x)
		sched[n] = idx
	}

	min := 0
	p := 1
	for k, v := range sched {
		for (min+v)%k != 0 {
			min += p
		}
		p *= k
	}
	fmt.Println("#2:", min)
}
