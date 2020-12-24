package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	input := in()
	totalSize := 220 // finger in the air?
	totalDays := 100
	grid := make([][]bool, totalSize)
	for x := 0; x < totalSize; x++ {
		grid[x] = make([]bool, totalSize)
	}

	for _, line := range input {
		carry := ""
		coords := []int{totalSize / 2, totalSize / 2} // pick the center

		for _, x := range strings.Split(line, "") {
			if x == "s" || x == "n" {
				carry = x
				continue
			}

			if carry != "" {
				x = carry + x
				carry = ""
			}

			switch x {
			case "nw":
				coords[0]--

			case "ne":
				coords[0]--
				coords[1]++

			case "e":
				coords[1]++

			case "se":
				coords[0]++

			case "sw":
				coords[0]++
				coords[1]--

			case "w":
				coords[1]--

			}
		}

		grid[coords[0]][coords[1]] = grid[coords[0]][coords[1]] == false
	}

	fmt.Println("#1:", totalCount(grid))

	// Part 2, simulate some days worth of the game of life, hex edition
	for day := 0; day < totalDays; day++ {
		newCount := totalSize + totalDays + day
		gridTwo := make([][]bool, newCount) // make sure it's big enough
		for x, row := range grid {
			gridTwo[x] = make([]bool, newCount)
			for y, v := range row {
				gridTwo[x][y] = v
			}
		}

		for x, row := range grid {
			for y, v := range row {
				n := countNeigh(grid, x, y)
				if v && (n == 0 || n > 2) {
					gridTwo[x][y] = false
				}
				if !v && n == 2 {
					gridTwo[x][y] = gridTwo[x][y] == false
				}
			}
		}

		grid = make([][]bool, newCount)
		for x, row := range gridTwo {
			grid[x] = make([]bool, newCount)
			for y, v := range row {
				grid[x][y] = v
			}
		}
	}

	fmt.Println("#2:", totalCount(grid))
}

func countNeigh(grid [][]bool, x, y int) (count int) {
	count = 0
	if x == 0 || x >= len(grid)-1 || y == 0 || y >= len(grid)-1 {
		return
	}

	if grid[x-1][y] {
		count++
	}
	if grid[x-1][y+1] {
		count++
	}
	if grid[x][y+1] {
		count++
	}
	if grid[x+1][y] {
		count++
	}
	if grid[x+1][y-1] {
		count++
	}
	if grid[x][y-1] {
		count++
	}

	return
}

func totalCount(grid [][]bool) (count int) {
	count = 0
	for _, x := range grid {
		for _, y := range x {
			if y {
				count++
			}
		}
	}
	return
}

func in() (in []string) {
	f, _ := os.Open("./in")
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		t := scanner.Text()

		in = append(in, t)
	}

	return
}
