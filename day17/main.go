package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	cycles = 6
)

func first(l []string) {
	start := len(l)
	max := len(l) + 12
	buffer := max * 2
	cube := make([][][]int, buffer)

	// init
	for x := 0; x < buffer; x++ {
		cube[x] = make([][]int, buffer)
		for y := 0; y < buffer; y++ {
			cube[x][y] = make([]int, buffer)
		}
	}

	// place the input data
	for x, line := range l {
		for y := 0; y < len(line); y++ {
			a := 0
			if line[y] == '#' {
				a = 1
			}
			cube[x+max][y+max][max] = a
		}
	}

	for c := 1; c <= cycles; c++ {
		sets := map[string]int{}
		for x := max - c; x < (max + start + c); x++ {
			for y := max - c; y < (max + start + c); y++ {
				for z := max - c; z < (max + start + c); z++ {
					neighbours := neigh(cube, x, y, z)

					if cube[x][y][z] == 1 && (neighbours == 2 || neighbours == 3) {
						sets[fmt.Sprintf("%d,%d,%d", x, y, z)] = 1
					} else if cube[x][y][z] == 0 && neighbours == 3 {
						sets[fmt.Sprintf("%d,%d,%d", x, y, z)] = 1
					} else if cube[x][y][z] == 1 {
						sets[fmt.Sprintf("%d,%d,%d", x, y, z)] = 0
					}
				}
			}
		}

		for x := max - c; x < (max + start + c); x++ {
			for y := max - c; y < (max + start + c); y++ {
				for z := max - c; z < (max + start + c); z++ {
					cube[x][y][z] = sets[fmt.Sprintf("%d,%d,%d", x, y, z)]
				}
			}
		}
	}

	total := 0
	for x := 0; x < buffer; x++ {
		for y := 0; y < buffer; y++ {
			for z := 0; z < buffer; z++ {
				if cube[x][y][z] == 1 {
					total++
				}
			}
		}
	}

	fmt.Println("#1:", total)
}

func second(l []string) {
	start := len(l)
	max := len(l) + 12
	buffer := max * 2
	cube := make([][][][]int, buffer)

	for a := 0; a < buffer; a++ {
		cube[a] = make([][][]int, buffer)
		for b := 0; b < buffer; b++ {
			cube[a][b] = make([][]int, buffer)
			for c := 0; c < buffer; c++ {
				cube[a][b][c] = make([]int, buffer)
			}
		}
	}

	for x, line := range l {
		for y := 0; y < len(line); y++ {
			a := 0
			if line[y] == '#' {
				a = 1
			}
			cube[x+max][y+max][max][max] = a
		}
	}

	for cycle := 1; cycle <= cycles; cycle++ {
		sets := map[string]int{}
		for a := max - cycle; a < (max + start + cycle); a++ {
			for b := max - cycle; b < (max + start + cycle); b++ {
				for c := max - cycle; c < (max + start + cycle); c++ {
					for d := max - cycle; d < (max + start + cycle); d++ {
						neighbours := neightoo(cube, a, b, c, d)

						if cube[a][b][c][d] == 1 && (neighbours == 2 || neighbours == 3) {
							sets[fmt.Sprintf("%d,%d,%d,%d", a, b, c, d)] = 1
						} else if cube[a][b][c][d] == 0 && neighbours == 3 {
							sets[fmt.Sprintf("%d,%d,%d,%d", a, b, c, d)] = 1
						} else if cube[a][b][c][d] == 1 {
							sets[fmt.Sprintf("%d,%d,%d,%d", a, b, c, d)] = 0
						}
					}
				}
			}
		}

		for x := max - cycle; x < (max + start + cycle); x++ {
			for y := max - cycle; y < (max + start + cycle); y++ {
				for z := max - cycle; z < (max + start + cycle); z++ {
					for w := max - cycle; w < (max + start + cycle); w++ {
						cube[x][y][z][w] = sets[fmt.Sprintf("%d,%d,%d,%d", x, y, z, w)]
					}
				}
			}
		}
	}

	total := 0
	for x := 0; x < buffer; x++ {
		for y := 0; y < buffer; y++ {
			for z := 0; z < buffer; z++ {
				for w := 0; w < buffer; w++ {
					if cube[x][y][z][w] == 1 {
						total++
					}
				}
			}
		}
	}

	fmt.Println("#2:", total)
}

func neigh(c [][][]int, x, y, z int) int {
	count := 0

	for a := x - 1; a < x+2; a++ {
		for b := y - 1; b < y+2; b++ {
			for d := z - 1; d < z+2; d++ {
				if a == x && b == y && d == z {
					continue
				}

				if c[a][b][d] == 1 {
					count++
				}
			}
		}
	}

	return count
}

func neightoo(cube [][][][]int, x, y, z, w int) int {
	count := 0

	for a := x - 1; a < x+2; a++ {
		for b := y - 1; b < y+2; b++ {
			for c := z - 1; c < z+2; c++ {
				for d := w - 1; d < w+2; d++ {
					if a == x && b == y && c == z && d == w {
						continue
					}

					if cube[a][b][c][d] == 1 {
						count++
					}
				}
			}
		}
	}

	return count
}

func dump(c [][][]int) {
	for _, x := range c {
		for _, y := range x {
			for _, z := range y {
				fmt.Printf("%d", z)
			}
			fmt.Println("")
		}
	}

	fmt.Println("")
}

func main() {
	f, _ := os.Open("./in")
	scanner := bufio.NewScanner(f)

	input := []string{}
	for scanner.Scan() {
		t := scanner.Text()
		input = append(input, t)
	}

	first(input)
	second(input)
}
