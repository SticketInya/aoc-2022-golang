package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Advent of Code 2023 - Day 2")

	task1()
	task2()
}

type Bag struct {
	Red   int
	Green int
	Blue  int
}

func task1() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(file), "\n")

	bagOfCubes := Bag{
		Red:   12,
		Green: 13,
		Blue:  14,
	}

	totalPossibleGameIndeces := 0
	for i, line := range lines {
		game := strings.Split(line, ":")
		if len(game) != 2 {
			log.Fatal("Invalid game line at ", i+1)
		}

		plays := strings.Split(game[1], ";")

		hasInvalidGame := false
		for _, play := range plays {
			re := regexp.MustCompile(`(\d*\s(blue|green|red))`)
			matches := re.FindAllString(play, -1)

			gameBag := BagFromStringArray(matches)
			if !bagOfCubes.isGamePossible(gameBag) {
				hasInvalidGame = true
				break
			}

		}
		if !hasInvalidGame {
			totalPossibleGameIndeces += i + 1
		}

	}

	fmt.Println("Total possible game indeces: ", totalPossibleGameIndeces)
}

func task2() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(file), "\n")

	sumOfPowerofCubes := 0
	for _, line := range lines {
		game := strings.Split(line, ":")
		if len(game) != 2 {
			log.Fatal("Invalid game line")
		}

		plays := strings.Split(game[1], ";")

		minBag := Bag{
			Red:   0,
			Green: 0,
			Blue:  0,
		}
		for _, play := range plays {
			re := regexp.MustCompile(`(\d*\s(blue|green|red))`)
			matches := re.FindAllString(play, -1)

			gameBag := BagFromStringArray(matches)

			if gameBag.Red > minBag.Red {
				minBag.Red = gameBag.Red
			}
			if gameBag.Green > minBag.Green {
				minBag.Green = gameBag.Green
			}
			if gameBag.Blue > minBag.Blue {
				minBag.Blue = gameBag.Blue
			}
		}
		sumOfPowerofCubes += minBag.getPowerOfBag()
	}

	fmt.Println("Sum of power of cubes: ", sumOfPowerofCubes)

}

func (b *Bag) isGamePossible(bagLike Bag) bool {
	if b.Red >= bagLike.Red && b.Green >= bagLike.Green && b.Blue >= bagLike.Blue {
		return true
	}

	return false
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
