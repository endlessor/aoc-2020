package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

type Food struct {
	Ingredients []string
	Allergens   []string
}

func (f Food) Contains(a string) bool {
	for _, x := range f.Allergens {
		if x == a {
			return true
		}
	}
	return false
}

func (f Food) MadeOf(i string) bool {
	for _, x := range f.Ingredients {
		if x == i {
			return true
		}
	}
	return false
}

var (
	actual      = map[string]string{}
	allergens   = map[string]struct{}{}
	ingredients = map[string]int{}

	definitelyMaybe = map[string]struct{}{}
	definitelyNot   = map[string]int{}

	inputFile = "./in"

	foods = []Food{}

	re = regexp.MustCompile(`^([a-z]+\s)+ \(contains (\w+)\)$`)
)

func main() {
	f, _ := os.Open(inputFile)
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		t := scanner.Text()

		f := Food{}
		bigger := strings.Split(t, " (")

		f.Ingredients = strings.Split(bigger[0], " ")
		f.Allergens = strings.Split(
			strings.TrimSuffix(
				strings.TrimPrefix(
					strings.Replace(
						bigger[1],
						",",
						"",
						-1,
					),
					"contains ",
				),
				")",
			),
			" ",
		)

		for _, x := range f.Allergens {
			allergens[x] = struct{}{}
		}
		for _, x := range f.Ingredients {
			ingredients[x]++
		}

		foods = append(foods, f)
	}

	potentials := map[string][]string{}
	for allergen := range allergens {
		count := 0
		ing := map[string]int{}

		for _, food := range foods {
			if food.Contains(allergen) {
				count++

				for _, i := range food.Ingredients {
					ing[i]++
				}
			}
		}

		for single, found := range ing {
			if found < count {
				if _, ok := definitelyMaybe[single]; !ok {
					definitelyNot[single]++
				}
			} else if found >= count {
				// could still be for the diff allergen
				delete(definitelyNot, single)
				definitelyMaybe[single] = struct{}{}

				potentials[allergen] = append(potentials[allergen], single)
			}
		}
	}

	totalCount := 0
	for _, food := range foods {
		for x := range definitelyNot {
			if food.MadeOf(x) {
				totalCount++
			}
		}
	}

	fmt.Println("#1:", totalCount)

	for {
		for allergen, list := range potentials {
			if len(list) == 1 {
				// do it meow
				containingIngredient := list[0]
				actual[allergen] = containingIngredient

				delete(potentials, allergen)

				for b, listing := range potentials {
					for x, o := range listing {
						if o == containingIngredient {
							newList := listing[:x]
							newList = append(newList, listing[x+1:]...)

							potentials[b] = newList
						}
					}
				}
			}
		}

		if len(potentials) < 2 {
			break
		}
	}

	keys := []string{}
	for k := range actual {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	ans := ""
	for _, x := range keys {
		ans += actual[x] + ","
	}
	ans = strings.TrimSuffix(ans, ",")

	fmt.Println("#2:", ans)
}
