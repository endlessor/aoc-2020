package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	f, _ := os.Open("./in")
	scanner := bufio.NewScanner(f)

	a := []int{}
	single := 1 // the first step
	triple := 1 // the adapter in the bag

	for scanner.Scan() {
		t := scanner.Text()

		n, _ := strconv.Atoi(t)
		a = append(a, n)
	}

	sort.Ints(a)

	for y, x := range a[:len(a)-1] {
		if a[y+1]-x == 1 {
			single++
		} else if a[y+1]-x == 3 {
			triple++
		}
	}

	fmt.Println("First answer,", single*triple)

	all := make(map[int]int)
	for i, v := range a {
		all[v] = i
	}

	o := make([]int, len(a))
	o[len(o)-1] = 1

	for x := len(o) - 2; x >= 0; x-- {
		count := 0

		for y := 1; y <= 3; y++ {
			if pos, ok := all[a[x]+y]; ok {
				count += o[pos]
			}
		}
		o[x] = count
	}

	ans := 0
	for v := 1; v <= 3; v++ {
		if pos, ok := all[v]; ok {
			ans += o[pos]
		}
	}

	fmt.Println("Second answer,", ans)
}
