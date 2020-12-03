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
	re = regexp.MustCompile(`(\d+)-(\d+) ([a-z]): (.+)$`)
)

func Parse(l string) (bool, bool) {
	a, b := false, false
	m := re.FindStringSubmatch(l)

	min, _ := strconv.Atoi(m[1])
	max, _ := strconv.Atoi(m[2])
	count := strings.Count(m[4], m[3])

	if count >= min && count <= max {
		a = true
	}

	// then the second calculation
	x := string(m[4][min-1]) == m[3]
	y := string(m[4][max-1]) == m[3]
	if (x || y) && !(x && y) {
		fmt.Println("True for", l)
		b = true
	}

	return a, b
}

func main() {
	valid := 0
	also := 0
	f, err := os.Open("./in")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		t := scanner.Text()
		a, b := Parse(t)
		if a {
			valid++
		}

		if b {
			also++
		}
	}

	fmt.Println(valid)
	fmt.Println(also)
}
