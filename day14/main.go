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
	maskRE = regexp.MustCompile(`mask = ([01X]+)$`)
	memRE  = regexp.MustCompile(`mem\[(\d+)\] = (\d+)$`)
)

func main() {
	first()
	second()
}

func first() {
	f, _ := os.Open("./in")
	scanner := bufio.NewScanner(f)

	memory := make([]int64, 128000)
	mask := ""

	for scanner.Scan() {
		t := scanner.Text()

		if maskRE.MatchString(t) {
			x := maskRE.FindStringSubmatch(t)
			mask = x[1]

		} else if memRE.MatchString(t) {
			x := memRE.FindStringSubmatch(t)

			v, _ := strconv.Atoi(x[2])
			n := string(strconv.FormatInt(int64(v), 2))

			// now have the binary rep for the value to be put in memory

			max := len(mask) - len(n)
			for a := 0; a < max; a++ {
				n = "0" + n
			}

			val := ""
			for a, b := range mask {
				if b == 'X' {
					val += string(n[a])
				} else {
					val += string(b)
				}
			}

			pos, _ := strconv.Atoi(x[1])
			dec, _ := strconv.ParseInt(val, 2, 64)
			memory[pos] = dec
		}

	}

	out := int64(0)
	for _, x := range memory {
		out += x
	}

	fmt.Println("#1:", out)
}

func second() {
	f, _ := os.Open("./in")
	scanner := bufio.NewScanner(f)

	memory := make(map[uint64]uint64)
	mask := ""

	for scanner.Scan() {
		t := scanner.Text()

		if maskRE.MatchString(t) {
			x := maskRE.FindStringSubmatch(t)
			mask = x[1]

		} else if memRE.MatchString(t) {
			x := memRE.FindStringSubmatch(t)

			v, _ := strconv.Atoi(x[2])
			i, _ := strconv.Atoi(x[1])
			n := string(strconv.FormatInt(int64(i), 2))
			max := len(mask) - len(n)
			for a := 0; a < max; a++ {
				n = "0" + n
			}

			addr := getListOfAddresses(n, mask)
			for _, x := range addr {
				newAdd, _ := strconv.ParseInt(x, 2, 64)
				memory[uint64(newAdd)] = uint64(v)
			}
		}
	}

	out := uint64(0)
	for i, x := range memory {
		fmt.Println(i, "\t:", x)
		out += x
	}

	fmt.Println(out)
}

func getListOfAddresses(over, mask string) []string {
	o := []string{""}

	str := strings.Split(mask, "")
	for pos, x := range str {
		if x != "X" {
			add := x
			if over[pos] == '1' {
				add = "1"
			}

			n := o
			for a, b := range o {
				n[a] = b + add
			}
			o = n
		} else {
			n := []string{}
			for _, y := range o {
				n = append(n, y+"0", y+"1")
			}
			o = n
		}
	}

	return o
}
