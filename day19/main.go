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
	rules        = map[int]string{}
	instructions = []string{}
)

func main() {
	f, _ := os.Open("./in")
	scanner := bufio.NewScanner(f)

	doRules := true
	for scanner.Scan() {
		t := scanner.Text()

		if len(t) < 1 {
			doRules = false
			continue
		}

		if doRules {
			s := strings.Split(t, ": ")
			n, _ := strconv.Atoi(s[0])
			rules[n] = s[1]
		} else {
			instructions = append(instructions, t)
		}
	}

	// turn the rules into a proper matching tree regexp?
	count := 0
	out := "^" + flatten(0, false, true) + "$"
	re := regexp.MustCompile(out)

	fmt.Println("Rule is", out)

	for _, b := range instructions {
		if re.MatchString(b) {
			count++
		}
	}

	fmt.Println("Count:", count)
}

func flatten(idx int, exponential, second bool) string {
	r := ""
	p := strings.Split(rules[idx], " | ")
	if len(p) > 1 {
		// have sub-sections
		r += "("
	}

	for a, b := range p {
		i := strings.Split(b, " ")
		for _, x := range i {
			r += "("
			if x[0] == '"' { // when the letter is defined
				r += string(x[1])
			} else { // standard next rule recursion
				n, _ := strconv.Atoi(x)

				if idx == 8 && second {
					// it's an exponential
					exponential = true
					r += flatten(n, false, second)
				} else if n != idx {
					r += flatten(n, false, second)
				}
			}

			r += ")"
			if exponential {
				r += "+"
			}
		}

		if a+1 < len(p) {
			r += "|"
		}
	}

	if len(p) > 1 {
		r += ")"
		if exponential {
			r += "+"
		}
	}

	return r
}
