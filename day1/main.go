package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	d := []int{}
	f, err := os.Open("./in")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		d = append(d, i)
	}

	for _, x := range d {
		for _, y := range d {
			if x != y && x+y == 2020 {
				fmt.Printf("%d + %d = 2020, %d * %d = %d\n", x, y, x, y, x*y)
			}
		}
	}

	for _, a := range d {
		for _, b := range d {
			if a == b {
				continue
			}

			for _, c := range d {
				if a == c || b == c {
					continue
				}

				if a+b+c == 2020 {
					fmt.Printf("%d + %d + %d = 2020, %d * %d * %d = %d\n", a, b, c, a, b, c, a*b*c)
				}
			}
		}
	}
}
