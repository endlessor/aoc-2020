package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var (
	re = regexp.MustCompile(`(\w)(\d+)$`)
)

func main() {
	first()
	second()
}

func first() {
	f, _ := os.Open("./in")
	scanner := bufio.NewScanner(f)

	east := 0
	north := 0
	dir := "E"

	for scanner.Scan() {
		t := scanner.Text()

		m := re.FindStringSubmatch(t)
		n, _ := strconv.Atoi(m[2])
		switch m[1] {
		case "F":
			// move forward
			switch dir {
			case "E":
				east += n
			case "N":
				north += n
			case "W":
				east -= n
			case "S":
				north -= n

			}

		case "L":
			d := n / 90
			for x := 0; x < d; x++ {
				switch dir {
				case "E":
					dir = "N"
				case "N":
					dir = "W"
				case "W":
					dir = "S"
				case "S":
					dir = "E"
				}
			}

		case "R":
			d := n / 90
			for x := 0; x < d; x++ {
				switch dir {
				case "E":
					dir = "S"
				case "N":
					dir = "E"
				case "W":
					dir = "N"
				case "S":
					dir = "W"
				}
			}

		case "S":
			north -= n

		case "N":
			north += n

		case "W":
			east -= n

		case "E":
			east += n

		}
	}

	if east < 0 {
		east *= -1
	}
	if north < 0 {
		north *= -1
	}

	fmt.Println("#1:", north+east)
}

func second() {
	f, _ := os.Open("./in")
	scanner := bufio.NewScanner(f)

	startEast := 0
	startNorth := 0

	east := 10
	north := 1

	for scanner.Scan() {
		t := scanner.Text()

		m := re.FindStringSubmatch(t)
		n, _ := strconv.Atoi(m[2])

		switch m[1] {
		case "N":
			north += n

		case "E":
			east += n

		case "S":
			north -= n

		case "W":
			east -= n

		case "L":
			d := n / 90
			for x := 0; x < d; x++ {
				nn := east
				ne := -1 * north

				east = ne
				north = nn
			}

		case "R":
			d := n / 90
			for x := 0; x < d; x++ {
				nn := -1 * east
				ne := north

				east = ne
				north = nn
			}

		case "F":
			startEast += (east * n)
			startNorth += (north * n)

		}
	}

	if startEast < 0 {
		startEast *= -1
	}
	if startNorth < 0 {
		startNorth *= -1
	}

	fmt.Println("#2:", startEast+startNorth)
}
