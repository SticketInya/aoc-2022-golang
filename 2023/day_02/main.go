package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Bag struct {
	Red   int
	Green int
	Blue  int
}

type Game struct {
	Index  int
	Rounds []Bag
}

func main() {
	fmt.Println("Advent of Code 2023 - Day 2")

	file, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	parsed := string(file)

	bagOfCubes := Bag{
		Red:   12,
		Green: 13,
		Blue:  14,
	}

	total := GetSumOfPossibleGameIndeces(parsed, &bagOfCubes)
	fmt.Println("Total possible game indeces: ", total)

	power := GetSumOfPowerOfGames(parsed)
	fmt.Println("Total power of games: ", power)
}

// Part 1 of the advent of code puzzle
// Sums the indeces of the games that are possible to play
func GetSumOfPossibleGameIndeces(input string, sourceBag *Bag) int {
	lines := strings.Split(input, "\n")

	totalPossibleGameIndeces := 0
	for _, line := range lines {
		game := NewGame(line)

		if game.isPossibleForBag(sourceBag) {
			totalPossibleGameIndeces += game.Index
		}
	}

	return totalPossibleGameIndeces
}

// Part 2 of the advent of code puzzle
// Sums the power of the games based on the minimum number
// of each cubed required for the game to be possible.
func GetSumOfPowerOfGames(input string) int {
	lines := strings.Split(input, "\n")

	totalPower := 0
	for _, line := range lines {
		game := NewGame(line)
		minBag := game.getMinimumPossibleBag()
		totalPower += minBag.getPowerOfBag()
	}

	return totalPower
}

func (b *Bag) isSubsetOfBag(other *Bag) bool {
	return b.Red <= other.Red && b.Green <= other.Green && b.Blue <= other.Blue
}

func BagFromStringArray(arr []string) Bag {
	var red, green, blue int

	for _, s := range arr {
		chars := strings.Split(s, " ")
		if len(chars) != 2 {
			log.Fatal("Invalid string array")
		}
		switch chars[1] {
		case "red":
			red, _ = strconv.Atoi(chars[0])
		case "green":
			green, _ = strconv.Atoi(chars[0])
		case "blue":
			blue, _ = strconv.Atoi(chars[0])
		}
	}

	return Bag{
		Red:   red,
		Green: green,
		Blue:  blue,
	}
}

func (b *Bag) getPowerOfBag() int {
	return b.Red * b.Green * b.Blue
}

func NewGame(s string) Game {

	index, err := GetGameIndex(s)
	if err != nil {
		log.Fatal(err)
	}

	return Game{
		Index:  index,
		Rounds: GetGameRounds(s),
	}
}

func GetGameIndex(s string) (int, error) {
	re := regexp.MustCompile(`(Game\s\d*:)`)
	match := regexp.MustCompile(`(\d+)`).FindString(re.FindString(s))

	return strconv.Atoi(match)
}

func GetGameRound(s string) Bag {
	re := regexp.MustCompile(`(\d*\s(blue|green|red))`)
	matches := re.FindAllString(s, -1)

	return BagFromStringArray(matches)
}

func GetGameRounds(s string) []Bag {
	rounds := strings.Split(strings.Split(s, ":")[1], ";")

	var bags []Bag
	for _, round := range rounds {
		bags = append(bags, GetGameRound(round))
	}

	return bags
}

func (g *Game) isPossibleForBag(b *Bag) bool {
	for _, roundBag := range g.Rounds {
		if !roundBag.isSubsetOfBag(b) {
			return false
		}
	}
	return true
}

func (g *Game) getMinimumPossibleBag() Bag {
	var minBag Bag
	for _, roundBag := range g.Rounds {
		if roundBag.Red > minBag.Red {
			minBag.Red = roundBag.Red
		}
		if roundBag.Green > minBag.Green {
			minBag.Green = roundBag.Green
		}
		if roundBag.Blue > minBag.Blue {
			minBag.Blue = roundBag.Blue
		}
	}
	return minBag
}
