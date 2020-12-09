package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
//re = regexp.MustCompile(`(jmp|acc|nop) ([+-])(\d+)$`)
)

func main() {
	f, _ := os.Open("./in")
	scanner := bufio.NewScanner(f)

	all := []int{}
	l := map[int]int{}
	i := 0
	low := 0

	ans := 0

	for scanner.Scan() {
		t := scanner.Text()

		n, _ := strconv.Atoi(t)
		all = append(all, n)
		l[i] = n

		if i > 25 {
			// validate the number now
			_, _, c := findPair(l, n)
			if !c {
				fmt.Println("First answer is,", n)
				ans = n
			} else {
				delete(l, low)
				low++
			}
		}

		i++
	}

	for y, x := range all {
		min := x
		max := x

		// start counting, to get to ans
		counter := 0
		for _, z := range all[y:] {
			if z < min {
				min = z
			} else if z > max {
				max = z
			}

			counter += z
			if counter == ans {
				fmt.Println("Second answer is,", min+max)

				return
			}

			if counter > ans {
				break
			}
		}
	}
}

func findPair(l map[int]int, n int) (int, int, bool) {
	for _, x := range l {
		for _, y := range l {
			if x != y && x+y == n {
				return x, y, true
			}
		}
	}
	return 0, 0, false
}
