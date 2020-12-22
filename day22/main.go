package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	PlayerA = []int{}
	PlayerB = []int{}

	PlayerASec = []int{}
	PlayerBSec = []int{}
)

func main() {
	f, _ := os.Open("./in")
	scanner := bufio.NewScanner(f)
	scanner.Scan()

	alt := false
	for scanner.Scan() {
		t := scanner.Text()

		if len(t) < 1 {
			scanner.Scan()
			scanner.Scan()
			t = scanner.Text()

			alt = true
		}

		n, _ := strconv.Atoi(t)
		if !alt {
			PlayerA = append(PlayerA, n)
		} else {
			PlayerB = append(PlayerB, n)
		}
	}

	// create a copy ready for part 2
	for _, x := range PlayerA {
		PlayerASec = append(PlayerASec, x)
	}
	for _, x := range PlayerB {
		PlayerBSec = append(PlayerBSec, x)
	}

	for round := 0; round >= 0; round++ {
		cardA := PlayerA[0]
		cardB := PlayerB[0]

		if cardA > cardB {
			PlayerA = append(PlayerA, cardA, cardB)

			PlayerA = PlayerA[1:]
			PlayerB = PlayerB[1:]
		} else {
			PlayerB = append(PlayerB, cardB, cardA)

			PlayerA = PlayerA[1:]
			PlayerB = PlayerB[1:]
		}

		if len(PlayerA) < 1 || len(PlayerB) < 1 {
			fmt.Println("Total rounds taken of", round)
			break
		}
	}

	if len(PlayerA) > 0 {
		fmt.Println("#1: Player 1 wins, total of", CalcTotal(PlayerA))
	} else if len(PlayerB) > 0 {
		fmt.Println("#1: Player 2 wins, total of", CalcTotal(PlayerB))
	}

	// Play the second game now, recursive
	donePlayerA, donePlayerB := Play(PlayerASec, PlayerBSec)
	if len(donePlayerA) > 0 {
		fmt.Println("#2: Player 1 wins, total of", CalcTotal(donePlayerA))
	} else {
		fmt.Println("#2: Player 2 wins, total of", CalcTotal(donePlayerB))
	}
}

func CalcTotal(l []int) int {
	total := 0
	length := len(l)

	for x := 0; x < length; x++ {
		total += (l[length-x-1] * (x + 1))
	}

	return total
}

func Play(playerA, playerB []int) ([]int, []int) {
	freshA := make([]int, 0)
	for _, y := range playerA {
		freshA = append(freshA, y)
	}
	freshB := make([]int, 0)
	for _, y := range playerB {
		freshB = append(freshB, y)
	}

	state := map[string]struct{}{}

	for round := 0; round >= 0; round++ {
		cardA := freshA[0]
		cardB := freshB[0]

		if cardA < len(freshA) && cardB < len(freshB) {
			subA := make([]int, 0)
			for _, x := range freshA[1 : cardA+1] {
				subA = append(subA, x)
			}
			subB := make([]int, 0)
			for _, x := range freshB[1 : cardB+1] {
				subB = append(subB, x)
			}

			newPlayerA, _ := Play(subA, subB)

			// A was the winner
			if len(newPlayerA) > 0 {
				freshA = append(freshA[1:], cardA, cardB)
				freshB = freshB[1:]
			} else {
				freshB = append(freshB[1:], cardB, cardA)
				freshA = freshA[1:]
			}

		} else if cardA > cardB {
			freshA = append(freshA[1:], cardA, cardB)
			freshB = freshB[1:]
		} else {
			freshB = append(freshB[1:], cardB, cardA)
			freshA = freshA[1:]
		}

		if _, ok := state[CreateState(freshA)+"::"+CreateState(freshB)]; ok {
			return freshA, freshB
		}
		state[CreateState(freshA)+"::"+CreateState(freshB)] = struct{}{}

		if len(freshA) < 1 || len(freshB) < 1 {
			break
		}
	}

	return freshA, freshB
}

func CreateState(a []int) string {
	out := ""
	for _, x := range a {
		out += fmt.Sprintf("%d,", x)
	}
	return out
}

func hasDupe(l []int) bool {
	k := make(map[int]struct{}, len(l))

	for _, x := range l {
		if _, ok := k[x]; ok {
			return true
		}
		k[x] = struct{}{}
	}
	return false
}
