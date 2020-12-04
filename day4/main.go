package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.Open("./in")
	scanner := bufio.NewScanner(f)
	total := 0

	l := ""
	for scanner.Scan() {
		t := scanner.Text()
		if len(t) < 1 {
			l = strings.TrimSpace(l)

			// process the current lines as a single passport
			if process(l) {
				total++
			}

			l = ""
		} else {
			// add the line to the current passport
			l = l + " " + t
		}
	}

	fmt.Printf("Total valid, %d\n", total)
}

func process(l string) bool {
	i := strings.Split(l, " ")
	c := 0

	for _, x := range i {
		if strings.HasPrefix(x, "byr:") {
			s := strings.TrimPrefix(x, "byr:")
			n, _ := strconv.Atoi(s)
			if n < 1920 || n > 2002 {
				return false
			}
			c++
		} else if strings.HasPrefix(x, "iyr:") {
			s := strings.TrimPrefix(x, "iyr:")
			n, _ := strconv.Atoi(s)
			if n < 2010 || n > 2020 {
				return false
			}
			c++
		} else if strings.HasPrefix(x, "eyr:") {
			s := strings.TrimPrefix(x, "eyr:")
			n, _ := strconv.Atoi(s)
			if n < 2020 || n > 2030 {
				return false
			}
			c++
		} else if strings.HasPrefix(x, "hgt:") {
			s := strings.TrimPrefix(x, "hgt:")
			if strings.HasSuffix(s, "cm") {
				a := strings.TrimSuffix(s, "cm")
				n, _ := strconv.Atoi(a)
				if n < 150 || n > 193 {
					return false
				}
				c++
			} else if strings.HasSuffix(s, "in") {
				a := strings.TrimSuffix(s, "in")
				n, _ := strconv.Atoi(a)
				if n < 59 || n > 76 {
					return false
				}
				c++
			}
		} else if strings.HasPrefix(x, "hcl:") {
			re := regexp.MustCompile(`hcl:#[a-z0-9]{6}`)
			if re.MatchString(x) {
				c++
			} else {
				return false
			}
		} else if strings.HasPrefix(x, "ecl:") {
			switch x {
			case "ecl:amb":
				c++
			case "ecl:blu":
				c++
			case "ecl:brn":
				c++
			case "ecl:gry":
				c++
			case "ecl:grn":
				c++
			case "ecl:hzl":
				c++
			case "ecl:oth":
				c++
			default:
				return false
			}
		} else if strings.HasPrefix(x, "pid:") {
			if len(x) != 13 {
				return false
			}
			c++
		}
	}

	return c >= 7
}
