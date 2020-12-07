package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Bag struct {
	Name    string
	Listing map[string]int
}

var (
	re = regexp.MustCompile(`(\d+) ([a-zA-Z ]+) bags?[.,]`)
)

func main() {
	first()
	second()
}

func first() {
	f, _ := os.Open("./in")
	scanner := bufio.NewScanner(f)

	within := map[string]struct{}{}
	bags := map[string]Bag{}

	for scanner.Scan() {
		t := scanner.Text()

		d := strings.Split(t, " bags contain ")
		m := re.FindAllStringSubmatch(d[1], -1)

		b := Bag{Name: d[0], Listing: map[string]int{}}
		for _, x := range m {
			n, _ := strconv.Atoi(x[1])
			b.Listing[x[2]] = n

			within[x[2]] = struct{}{}
		}

		bags[d[0]] = b
	}

	// determine if each bag can possibly have the gold bag
	total := 0
	for name, _ := range bags {
		if can(bags, name) {
			total++
		}
	}

	fmt.Println(total)
}

func can(b map[string]Bag, n string) bool {
	if b[n].Listing["shiny gold"] > 0 {
		return true
	}

	for cur, _ := range b[n].Listing {
		if can(b, cur) {
			return true
		}
	}

	return false
}

func second() {
	f, _ := os.Open("./in")
	scanner := bufio.NewScanner(f)
	bags := map[string]Bag{}

	for scanner.Scan() {
		t := scanner.Text()

		d := strings.Split(t, " bags contain ")
		m := re.FindAllStringSubmatch(d[1], -1)

		b := Bag{Name: d[0], Listing: map[string]int{}}
		for _, x := range m {
			n, _ := strconv.Atoi(x[1])
			b.Listing[x[2]] = n
		}

		bags[d[0]] = b
	}

	fmt.Println(count(bags, "shiny gold"))
}

func count(b map[string]Bag, n string) int {
	total := 0
	for name, c := range b[n].Listing {
		total = total + c + (c * count(b, name))
	}

	return total
}
