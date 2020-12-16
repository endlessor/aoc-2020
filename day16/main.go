package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	ranges = regexp.MustCompile(`(\d+)-(\d+)`)
	rules  = []rule{}
)

type rule struct {
	name  string
	taken bool
	using int

	valid []int

	minA int
	maxA int

	minB int
	maxB int
}

func (r rule) has(x int) bool {
	if x >= r.minA && x <= r.maxA {
		return true
	}
	if x >= r.minB && x <= r.maxB {
		return true
	}
	return false
}

func main() {
	f, _ := os.Open("./in")
	scanner := bufio.NewScanner(f)

	myTicket := []int{}
	validTickets := [][]int{}
	allTypes := make(map[string][]int)

	out := 0
	types := make(map[int]struct{})
	group := 0
	missed := []int{}

	for scanner.Scan() {
		t := scanner.Text()
		if len(t) < 1 {
			group++
			scanner.Scan()
			continue
		}

		switch group {
		case 0:
			b := strings.Split(t, ":")
			x := ranges.FindAllStringSubmatch(strings.TrimSpace(b[1]), -1)

			r := rule{name: b[0]}

			for y, a := range x {
				low, _ := strconv.Atoi(a[1])
				high, _ := strconv.Atoi(a[2])
				for y := low; y <= high; y++ {
					types[y] = struct{}{}
					allTypes[b[0]] = append(allTypes[b[0]], y)
				}

				if y == 0 {
					r.minA = low
					r.maxA = high
				} else {
					r.minB = low
					r.maxB = high
				}
			}

			rules = append(rules, r)

		case 1:
			x := strings.Split(t, ",")
			for _, a := range x {
				n, _ := strconv.Atoi(a)
				myTicket = append(myTicket, n)
			}

		case 2:
			m := ticketMisses(types, t)
			missed = append(missed, m...)
			if len(m) > 0 {
				// skip a bad ticket
				continue
			}

			x := []int{}
			for _, a := range strings.Split(t, ",") {
				b, _ := strconv.Atoi(a)
				x = append(x, b)
			}
			validTickets = append(validTickets, x)

		}
	}

	for _, x := range missed {
		out += x
	}

	fmt.Println("#1:", out)

	for offset := 19; offset >= 0; offset-- {
		for a, b := range rules {
			v := true
			for _, x := range validTickets {
				if !b.has(x[offset]) {
					v = false
				}
			}

			if b.has(myTicket[offset]) && v {
				rules[a].taken = true
				rules[a].using = offset
				rules[a].valid = append(rules[a].valid, offset)
			}
		}
	}

	ans := make(map[int]rule)
	count := 0
	for {
		lowIdx := 0
		least := 20

		for a, b := range rules {
			x := len(b.valid)
			if x < least && x > 0 {
				lowIdx = a
				least = x
			}
		}

		num := rules[lowIdx].valid[0]
		ans[num] = rules[lowIdx]
		rules[lowIdx].valid = []int{}

		for a, b := range rules {
			l := len(b.valid)
			for x := 0; x < l; x++ {
				if b.valid[x] == num {
					s := []int{}

					if l == 1 {
						// nop
					} else if x == 0 {
						s = append(s, b.valid[1:]...)
					} else if x == l {
						s = append(s, b.valid[:x-1]...)
					} else {
						s = append(s, b.valid[:x]...)
						s = append(s, b.valid[x+1:]...)
					}

					rules[a].valid = s
					break
				}
			}
		}

		count++
		if count == 20 {
			break
		}
	}

	out = 1
	for a, b := range ans {
		if strings.HasPrefix(b.name, "departure ") {
			out *= myTicket[a]
		}
	}

	fmt.Println("#2:", out)
}

func ticketMisses(t map[int]struct{}, i string) []int {
	r := []int{}

	for _, x := range strings.Split(i, ",") {
		n, _ := strconv.Atoi(x)

		if _, ok := t[n]; !ok {
			r = append(r, n)
		}
	}

	return r
}
