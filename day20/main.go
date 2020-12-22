package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Image struct {
	TopSide    string
	BottomSide string
	LeftSide   string
	RightSide  string

	InnerImage []string

	ID      int
	Matches int
	Applied bool

	X int
	Y int

	OnBottom int
	OnLeft   int
	OnTop    int
	OnRight  int
}

func (i Image) AlignTo(t Image, side string) (n Image) {
	n = i

	for loop := 0; loop < 4; loop++ {
		if !t.CanMatch(i.GetSide(side)) {
			n = n.Rotate()
		}
	}

	if !t.CanMatch(i.GetSide(side)) {
		n = n.Flip()
	}

	for loop := 0; loop < 4; loop++ {
		if !t.CanMatch(i.GetSide(side)) {
			n = n.Rotate()
		}
	}

	return
}

// CanMatch the current image against a provided side?
func (i Image) CanMatch(side string) bool {
	r := rev(side)
	if i.TopSide == side ||
		i.TopSide == r ||
		i.RightSide == side ||
		i.RightSide == r ||
		i.BottomSide == side ||
		i.BottomSide == r ||
		i.LeftSide == side ||
		i.LeftSide == r {
		return true
	}

	return false
}

func (i Image) GetSide(s string) string {
	switch s {
	case "top":
		return i.TopSide
	case "right":
		return i.RightSide
	case "bottom":
		return i.BottomSide
	default:
		return i.LeftSide
	}
}

// Rotate an image clockwise 90 degrees.
func (i Image) Rotate() (n Image) {
	n = i
	n.RightSide = i.TopSide
	n.OnRight = i.OnTop

	n.BottomSide = i.RightSide
	n.OnBottom = i.OnRight

	n.LeftSide = i.BottomSide
	n.OnLeft = i.OnBottom

	n.TopSide = i.LeftSide
	n.OnTop = i.OnLeft

	n.InnerImage = make([]string, len(i.InnerImage))
	for _, line := range i.InnerImage {
		for a, b := range line {
			n.InnerImage[a] = string(b) + n.InnerImage[a]
		}
	}

	return
}

// Flip an image horizontally.
func (i Image) Flip() (n Image) {
	n = i
	n.RightSide = i.LeftSide
	n.OnRight = i.OnLeft

	n.LeftSide = i.RightSide
	n.OnLeft = i.OnRight

	n.TopSide = rev(i.TopSide)
	n.BottomSide = rev(i.BottomSide)

	r := []string{}
	for _, x := range i.InnerImage {
		r = append(r, rev(x))
	}
	n.InnerImage = r

	return
}

func (i Image) FlipVert() (n Image) {
	n = i
	n.BottomSide = i.TopSide
	n.OnBottom = i.OnTop

	n.TopSide = i.BottomSide
	n.OnTop = i.OnBottom

	n.LeftSide = rev(i.LeftSide)
	n.RightSide = rev(i.RightSide)

	r := []string{}
	for _, x := range i.InnerImage {
		a := []string{x}
		a = append(a, r...)
		r = a
	}
	n.InnerImage = r

	return
}

func (i Image) Print() {
	fmt.Println(" **", i.ID, "**")
	for _, x := range i.InnerImage {
		fmt.Println(x)
	}
	fmt.Println("")
}

var (
	corners      = []Image{}
	grid         = make([][]Image, 0)
	imageLength  = 1
	tileMap      = make(map[int]Image)
	tiles        = []Image{}
	totalHash    = 0
	totalMonster = 0
	used         = make(map[int]Image)
)

func main() {
	tile := Image{}
	previousLine := ""

	f, _ := os.Open("./t")
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		t := scanner.Text()

		if strings.HasPrefix(t, "Tile ") {
			t = strings.TrimPrefix(strings.TrimSuffix(t, ":"), "Tile ")
			n, _ := strconv.Atoi(t)
			tile.ID = n

			scanner.Scan()
			tile.TopSide = scanner.Text()
			tile.LeftSide = string(tile.TopSide[0])
			tile.RightSide = string(tile.TopSide[len(tile.TopSide)-1])

		} else if t == "" {
			tile.BottomSide = previousLine
			tile.InnerImage = tile.InnerImage[:len(tile.InnerImage)-1]
			tiles = append(tiles, tile)
			tile = Image{}

		} else {
			tile.InnerImage = append(tile.InnerImage, string(t[1:len(t)-1]))
			tile.LeftSide += string(t[0])
			tile.RightSide += string(t[len(t)-1])
			previousLine = t

		}
	}

	startingIdentifier := 0
	corners := 1
	for a, x := range tiles {
		tiles[a] = tileWithMatch(x)
		tileMap[x.ID] = tiles[a]

		if tiles[a].Matches == 2 {
			corners *= tiles[a].ID
			if startingIdentifier == 0 || (tiles[a].OnLeft == 0 && tiles[a].OnTop == 0) {
				startingIdentifier = tiles[a].ID
			}
		}
	}
	fmt.Println("#1:", corners)

	// Now part 2

	// 1. Run through the images, building up the full image
	// 2. Collapse the full image into a single reference
	// 3. Count the `#`
	// 4. Scan for sea monsters
	// 5. If none, rotate the image and try again
	// 6. If all rotated, flip and begin again
	// 7. Answer is (total `#` - (15 * number of monsters))

	for a := 0; a < 4; a++ {
		if tileMap[startingIdentifier].OnTop != 0 || tileMap[startingIdentifier].OnLeft != 0 {
			tileMap[startingIdentifier] = tileMap[startingIdentifier].Rotate()
		}
	}

	imageLength = int(math.Sqrt(float64(len(tiles))))
	grid = make([][]Image, imageLength)
	for idx := 0; idx < imageLength; idx++ {
		grid[idx] = make([]Image, imageLength)
	}

	for y := 0; y < imageLength; y++ {
		for x := 0; x < imageLength; x++ {
			if y == 0 && x == 0 {
				grid[0][0] = tileMap[startingIdentifier]
				continue
			}

			if y == 0 {
				previousID := grid[0][x-1].ID
				nt := tileMap[grid[0][x-1].OnRight]

				for a := 0; a < 4; a++ {
					if nt.OnLeft != previousID {
						nt = nt.Rotate()
					}
				}
				if nt.OnTop != 0 {
					nt = nt.FlipVert()
				}

				grid[0][x] = nt
			} else if x == 0 {
				previousID := grid[y-1][0].ID
				nt := tileMap[grid[y-1][0].OnBottom]

				for a := 0; a < 4; a++ {
					if nt.OnTop != previousID {
						nt = nt.Rotate()
					}
				}
				if nt.OnLeft != 0 {
					nt = nt.Flip()
				}

				grid[y][0] = nt
			} else if y == imageLength-1 {
				previousID := grid[y][x-1].ID
				nt := tileMap[grid[y][x-1].OnRight]

				for a := 0; a < 8; a++ {
					if nt.OnLeft != previousID && nt.OnTop != grid[y-1][x].ID {
						//fmt.Println("Rotate for the inner along the bottom")
						nt = nt.Rotate()
					} else {
						//fmt.Println("Good for the bottom row")
					}
				}
				if nt.OnBottom != 0 {
					nt = nt.FlipVert()
				}

				grid[y][x] = nt
				//fmt.Println(nt.OnLeft, "should be", previousID)
				//fmt.Println(nt.OnTop, "should be", grid[y-1][x].ID)
			} else if x == imageLength-1 {
				previousID := grid[y-1][x].ID
				nt := tileMap[grid[y-1][x].OnBottom]

				for a := 0; a < 4; a++ {
					if nt.OnTop != previousID {
						nt = nt.Rotate()
					}
				}
				if nt.OnRight != 0 {
					nt = nt.Flip()
				}

				grid[y][x] = nt
				//fmt.Println(nt.OnTop, "should be", previousID)
				//fmt.Println(nt.OnRight, "should be 0")
				//fmt.Println(nt.OnLeft, "should be", grid[y][x-1].ID)
			} else {
				previousID := grid[y][x-1].ID
				nt := tileMap[grid[y][x-1].OnRight]

				for a := 0; a < 4; a++ {
					if nt.OnLeft != previousID {
						nt = nt.Rotate()
					}
				}
				if nt.OnTop != grid[y-1][x].ID {
					nt = nt.FlipVert()
				}

				grid[y][x] = nt
				//fmt.Println(nt.OnLeft, "should be", previousID)
			}
		}
	}

	output := make([]string, 8*imageLength)
	for rowID, lineOfImage := range grid {
		for x := 0; x < 8; x++ {
			for _, img := range lineOfImage {
				output[rowID*8+x] += img.InnerImage[x]
			}
		}
	}

	for x := 0; x < 8; x++ {
		totalMonster = countMonster(output)
		if totalMonster < 1 {
			output = rotateOutput(output)
		} else {
			break
		}

		if x != 0 && x%4 == 0 {
			output = flipOutput(output)
		}
	}

	for _, y := range output {
		for _, x := range y {
			if x == '#' {
				totalHash++
			}
		}
	}

	fmt.Println("#2:", totalHash-(totalMonster*15))
}

func countMonster(i []string) int {
	count := 0
	beginY := 1
	stopY := len(i) - 1
	beginX := 0
	stopX := len(i) - 20

	for y := beginY; y < stopY; y++ {
		for x := beginX; x <= stopX; x++ {

			if i[y][x] == '#' &&
				i[y+1][x+1] == '#' &&
				i[y+1][x+4] == '#' &&
				i[y][x+5] == '#' &&
				i[y][x+6] == '#' &&
				i[y+1][x+7] == '#' &&
				i[y+1][x+10] == '#' &&
				i[y][x+11] == '#' &&
				i[y][x+12] == '#' &&
				i[y+1][x+13] == '#' &&
				i[y+1][x+16] == '#' &&
				i[y][x+17] == '#' &&
				i[y-1][x+18] == '#' &&
				i[y][x+18] == '#' &&
				i[y][x+19] == '#' {

				count++

			}
		}
	}

	return count
}

func flipOutput(in []string) (o []string) {
	o = make([]string, len(in))
	for x, line := range in {
		o[len(in)-x-1] = line
	}

	return
}

func rotateOutput(in []string) (o []string) {
	o = make([]string, len(in))
	for _, line := range in {
		for a, b := range line {
			o[a] = string(b) + o[a]
		}
	}

	return
}

func tileWithMatch(t Image) (newImage Image) {
	newImage = t

	for _, x := range tiles {
		if x.ID == t.ID {
			continue
		}

		if x.CanMatch(t.TopSide) {
			newImage.Matches++
			newImage.OnTop = x.ID
		} else if x.CanMatch(t.RightSide) {
			newImage.Matches++
			newImage.OnRight = x.ID
		} else if x.CanMatch(t.BottomSide) {
			newImage.Matches++
			newImage.OnBottom = x.ID
		} else if x.CanMatch(t.LeftSide) {
			newImage.Matches++
			newImage.OnLeft = x.ID
		}
	}

	return
}

func rev(m string) string {
	rns := []rune(m)
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}
	return string(rns)
}
