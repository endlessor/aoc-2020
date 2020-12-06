package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	first()
	second()
}

func first() {
	f, _ := os.Open("./in")
	scanner := bufio.NewScanner(f)
	total := 0

	g := ""
	for scanner.Scan() {
		t := scanner.Text()
		if len(t) < 1 {
			x := make(map[string]bool, 26)
			y := strings.Split(g, "")

			for _, z := range y {
				x[z] = true
			}

			total += len(x)

			g = ""
		} else {
			g += t
		}
	}

	fmt.Println("Total,", total)
}

func second() {
	f, _ := os.Open("./in")
	scanner := bufio.NewScanner(f)
	total := 0

	g := ""
	loop := 0
	for scanner.Scan() {
		t := scanner.Text()
		if len(t) < 1 {
			count := 0
			x := make(map[string]int, 26)
			y := strings.Split(g, "")

			for _, z := range y {
				x[z]++
			}

			for _, z := range x {
				if z == loop {
					count++
				}
			}

			total += count

			g = ""
			loop = 0
		} else {
			g += t
			loop++
		}
	}

	fmt.Println("Total,", total)
}
