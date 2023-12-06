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
	fmt.Println("Advent of Code 2023 - Day 6")

	file, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	contents := strings.Split(string(file), "\n")

	productOfWins := getProductOfPossibleWaysToWin(contents)
	fmt.Println("Product of possible ways to win:", productOfWins)

	productOfWins = getProductOfPossibleWaysToWinWithKerning(contents)
	fmt.Println("Product of possible ways to win with kerning:", productOfWins)
}

type Round struct {
	Time     int
	Distance int
}

type Race struct {
	Rounds []Round
}

func NewRace(input []string) *Race {
	digitRe := regexp.MustCompile(`\d+`)

	race := &Race{}
	times := []string{}
	distances := []string{}
	for i, line := range input {
		if i == 0 {
			times = append(times, digitRe.FindAllString(line, -1)...)
		} else {
			distances = append(distances, digitRe.FindAllString(line, -1)...)
		}
	}

	for i := 0; i < len(times); i++ {
		time, _ := strconv.Atoi(times[i])
		distance, _ := strconv.Atoi(distances[i])
		race.Rounds = append(race.Rounds, Round{Time: time, Distance: distance})
	}

	return race
}

func NewRaceWithKerning(input []string) *Race {
	digitRe := regexp.MustCompile(`\d+`)

	race := &Race{}
	round := Round{}
	for i, line := range input {
		if i == 0 {
			time, _ := strconv.Atoi(strings.Join(digitRe.FindAllString(line, -1), ""))
			round.Time = time
		} else {
			distances, _ := strconv.Atoi(strings.Join(digitRe.FindAllString(line, -1), ""))
			round.Distance = distances
		}
	}

	race.Rounds = append(race.Rounds, round)

	return race
}

func (r *Round) getPossibleWins() int {
	// TODO - Use the quadratic formula and calculate for ranges
	possibleWins := []int{}

	for i := 1; i < r.Time; i++ {
		calced := (r.Time - i) * i
		if calced > r.Distance {
			possibleWins = append(possibleWins, calced)
		}
	}
	return len(possibleWins)
}

func (r *Race) getProductOfWins() int {
	product := 1
	for _, round := range r.Rounds {
		product *= round.getPossibleWins()
	}
	return product
}

func getProductOfPossibleWaysToWin(input []string) int {
	race := NewRace(input)
	return race.getProductOfWins()
}

func getProductOfPossibleWaysToWinWithKerning(input []string) int {
	race := NewRaceWithKerning(input)
	return race.getProductOfWins()
}
