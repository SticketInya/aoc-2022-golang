package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strings"
)

func main() {
	fmt.Println("Advent of Code 2023 - Day 10")
	file, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	contents := strings.Split(string(file), "\n")

	furthestPosition := getFurthestPosition(contents)
	fmt.Println("The furthest position from the start is at:", furthestPosition)
}

type Tile struct {
	Value   rune
	X       int
	Y       int
	Visited bool
	Inside  bool
}

func (t *Tile) getNextPosition(m [][]*Tile, direction rune, start *Tile) int {
	// fmt.Println("Current tile:", t)
	t.Visited = true

	if start != nil && start == t {
		return 0
	}

	startTile := start
	if startTile == nil {
		startTile = t
	}

	var nextDir rune
	for _, dir := range t.getDirections() {
		if dir != direction {
			nextDir = dir
			break
		}
	}

	var nextTile *Tile
	switch nextDir {
	case 'N':
		nextTile = m[t.Y-1][t.X]
	case 'S':
		nextTile = m[t.Y+1][t.X]
	case 'E':
		nextTile = m[t.Y][t.X+1]
	case 'W':
		nextTile = m[t.Y][t.X-1]
	}

	var oppDir rune
	switch nextDir {
	case 'N':
		oppDir = 'S'
	case 'S':
		oppDir = 'N'
	case 'E':
		oppDir = 'W'
	case 'W':
		oppDir = 'E'
	default:
		log.Fatal("Invalid direction")
	}

	return 1 + nextTile.getNextPosition(m, oppDir, startTile)
}

func (t *Tile) getDirections() []rune {
	switch t.Value {
	case '┌':
		return []rune{'S', 'E'}
	case '┐':
		return []rune{'S', 'W'}
	case '└':
		return []rune{'N', 'E'}
	case '┘':
		return []rune{'N', 'W'}
	case '│':
		return []rune{'N', 'S'}
	case '─':
		return []rune{'W', 'E'}
	default:
		return []rune{}
	}
}

func getPipeSymbol(c rune) rune {
	switch c {
	case 'S':
		return 'S'
	case 'F':
		return '┌'
	case 'J':
		return '┘'
	case 'L':
		return '└'
	case '-':
		return '─'
	case '7':
		return '┐'
	case '|':
		return '│'
	default:
		return c
	}
}

func (t *Tile) replaceSymbol(m [][]*Tile) {
	north := slices.Contains(m[t.Y-1][t.X].getDirections(), 'S')
	south := slices.Contains(m[t.Y+1][t.X].getDirections(), 'N')
	east := slices.Contains(m[t.Y][t.X+1].getDirections(), 'W')
	west := slices.Contains(m[t.Y][t.X-1].getDirections(), 'E')

	// s := m[t.X+1][t.Y].getDirections()
	// e := m[t.X][t.Y+1].getDirections()

	// fmt.Println("North:", north, "South:", south, "East:", east, "West:", west)
	// fmt.Println("Current tile:", string(t.Value))
	// fmt.Println("South:", string(s[0]), string(s[1]))
	// fmt.Println("East:", string(e[0]), string(e[1]))

	if north && south {
		t.Value = '│'
	} else if north && east {
		t.Value = '└'
	} else if north && west {
		t.Value = '┘'
	} else if south && east {
		t.Value = '┌'
	} else if south && west {
		t.Value = '┐'
	} else if east && west {
		t.Value = '─'
	}
}

func normalizeInput(input []string) []string {
	l := len(input[0])
	normalized := []string{
		strings.Repeat(".", l),
	}
	for _, line := range input {
		normalized = append(normalized, "."+line+".")
	}
	normalized = append(normalized, strings.Repeat(".", l))
	return normalized
}

func getIsInside(m map[rune]int) bool {
	isInside := false
	if m['│']%2 == 1 {
		isInside = !isInside
	}

	if (m['└']-m['┘'])%2 == 1 {
		isInside = !isInside
	}

	if (m['┌']-m['┐'])%2 == 1 {
		isInside = !isInside
	}

	if (m['└']-m['┐'])%2 == 0 {
		isInside = !isInside
	}

	if (m['┌']-m['┘'])%2 == 0 {
		isInside = !isInside
	}

	// fmt.Println("Inside:", isInside, "│", m['│'], "└", m['└'], "┘", m['┘'], "┌", m['┌'], "┐", m['┐'])

	return isInside
}

func printMap(m [][]*Tile, s *Tile) {
	colorReset := "\033[0m"
	colorRed := "\033[31m"
	colorBlue := "\033[34m"
	colorGreen := "\033[32m"

	for _, line := range m {
		for _, tile := range line {
			if tile.Visited {
				fmt.Print(colorRed)
			}
			if tile.Inside {
				fmt.Print(colorGreen)
			}
			if tile == s {
				fmt.Print(colorBlue)
				fmt.Print("S")
			} else {
				fmt.Print(string(tile.Value))
			}
			fmt.Print(colorReset)
		}
		fmt.Println()
	}
}

func getFurthestPosition(input []string) int {
	normalized := normalizeInput(input)
	areaMap := [][]*Tile{}
	var start *Tile

	for i, line := range normalized {
		subMap := []*Tile{}
		for j, char := range line {
			t := Tile{
				Value: getPipeSymbol(char),
				X:     j,
				Y:     i,
			}

			if char == 'S' {
				start = &t
			}
			subMap = append(subMap, &t)
		}

		areaMap = append(areaMap, subMap)
	}

	if start == nil {
		log.Fatal("Start not found")
	}
	// fmt.Println("Start at:", start.X, start.Y)

	start.replaceSymbol(areaMap)

	// fmt.Println("Start replaced to:", string(start.Value))
	dir := start.getDirections()
	// fmt.Println("Directions:", string(dir[0]), string(dir[1]))

	loopLength := start.getNextPosition(areaMap, dir[0], nil)

	count := 0
	for i, line := range areaMap {
		pipeMap := map[rune]int{}
		for j, tile := range line {
			if tile.Visited {
				pipeMap[tile.Value]++
			} else if getIsInside(pipeMap) {
				areaMap[i][j].Inside = true
				count++
			}
		}

	}

	printMap(areaMap, start)
	fmt.Println("Count:", count)
	return int(math.Ceil(float64(loopLength) / 2))
}
