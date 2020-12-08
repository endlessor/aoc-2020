package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type ins struct {
	n   string
	up  bool
	val int
}

var (
	re = regexp.MustCompile(`(jmp|acc|nop) ([+-])(\d+)$`)
)

func main() {
	f, _ := os.Open("./in")
	scanner := bufio.NewScanner(f)

	acc := 0
	idx := map[int]ins{}
	i := 0
	visited := map[int]int{}

	for scanner.Scan() {
		t := scanner.Text()

		x := re.FindStringSubmatch(t)

		n, _ := strconv.Atoi(x[3])
		idx[i] = ins{n: x[1], up: x[2] == "+", val: n}
		i++
	}

	order := map[int]int{}

	// found the success condition to change line 205 to nop
	// for the second solution, make the above change in the input file
	//	and re-run the program

	l := 0
	i = 0
	for {
		visited[i] = l
		order[l] = i

		if _, ok := idx[i]; !ok {
			fmt.Println("Terminating normally.", acc)
			break
		}

		switch idx[i].n {
		case "jmp":
			if idx[i].up {
				i += idx[i].val
				if _, ok := visited[i]; ok {
					panic(acc)
				}
			} else {
				i -= idx[i].val
				if _, ok := visited[i]; ok {
					panic(acc)
				}
			}

		case "acc":
			if idx[i].up {
				acc += idx[i].val
			} else {
				acc -= idx[i].val
			}
			i++

		case "nop":
			i++
		}

		l++
	}
}
