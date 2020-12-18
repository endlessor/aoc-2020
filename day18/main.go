package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	first()
	second()
}

func first() {
	f, _ := os.Open("./in")
	scanner := bufio.NewScanner(f)

	sum := uint64(0)
	for scanner.Scan() {
		t := scanner.Text()

		sum += expr(t, false)
	}

	fmt.Println("#1:", sum)
}

func second() {
	f, _ := os.Open("./in")
	scanner := bufio.NewScanner(f)

	sum := uint64(0)
	for scanner.Scan() {
		t := scanner.Text()

		x := expr(t, true)
		sum += x
	}

	fmt.Println("#2:", sum)
}

func p(n []uint64) {
	for _, x := range n {
		fmt.Printf("%d ", x)
	}
	fmt.Println("")
}

func computeLayer(n []uint64, o []string) uint64 {
	total := uint64(1)

	newNum := []uint64{}
	newOp := []string{}
	offset := 0
	opOffset := 0

	for a, b := range n {
		if a == 0 {
			newNum = append(newNum, b)
			continue
		}

		if o[opOffset] == "+" {
			newNum[offset] += b
		} else {
			newNum = append(newNum, b)
			newOp = append(newOp, o[opOffset])
			offset++
		}
		opOffset++
	}

	for _, x := range newNum {
		total *= x
	}

	return total
}

func expr(e string, two bool) uint64 {
	// the layer, in terms of parenthesis
	depth := 0
	// the running total in that layer
	layerTotal := make(map[int]uint64, 1)
	// the current operation for this layer
	layerOp := make(map[int]string, 1)

	// the list of operations for this equation
	layerOps := make(map[int][]string, 1)
	// the list of numbers for this equation
	layerNums := make(map[int][]uint64, 1)

	for _, b := range strings.Split(e, "") {
		switch b {
		case " ":
			continue
		case "(":
			depth++
			if two {
				layerOps[depth] = []string{}
				layerNums[depth] = []uint64{}

				continue
			}

			// reset the layer
			layerOp[depth] = ""
			layerTotal[depth] = 0

		case ")":
			depth--
			if two {
				val := computeLayer(layerNums[depth+1], layerOps[depth+1])
				layerNums[depth] = append(layerNums[depth], val)

				continue
			}

			if layerOp[depth] == "" {
				layerTotal[depth] = layerTotal[depth+1]
			} else {
				if layerOp[depth] == "*" {
					layerTotal[depth] *= layerTotal[depth+1]
				} else {
					layerTotal[depth] += layerTotal[depth+1]
				}
			}

		case "*":
			if two {
				layerOps[depth] = append(layerOps[depth], b)
				continue
			}
			layerOp[depth] = b

		case "+":
			if two {
				layerOps[depth] = append(layerOps[depth], b)
				continue
			}
			layerOp[depth] = b

		default:
			n, _ := strconv.Atoi(b)

			if two {
				layerNums[depth] = append(layerNums[depth], uint64(n))
				continue
			}

			if layerOp[depth] == "" {
				layerTotal[depth] = uint64(n)
			} else {
				if layerOp[depth] == "*" {
					layerTotal[depth] *= uint64(n)
				} else {
					layerTotal[depth] += uint64(n)
				}
			}

		}
	}

	if two {
		return computeLayer(layerNums[0], layerOps[0])
	}
	return layerTotal[0]
}
