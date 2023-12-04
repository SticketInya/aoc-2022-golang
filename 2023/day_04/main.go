package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Advent of Code 2023 - Day 4")

	file, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	contents := strings.Split(string(file), "\n")

	totalValue := getTotalValueOfCards(contents)
	fmt.Println("The total value of the scratchcards is:", totalValue)

	totalCopies := getTotalCopiesOfCards(contents)
	fmt.Println("The total copies of the scratchcards is:", totalCopies)
}

type ScratchCard struct {
	Winning []int
	Present []int
	Count   int
}

func NewScratchCard(vinput string) *ScratchCard {
	numbers := strings.Split(vinput, ":")
	split := strings.Split(numbers[1], "|")

	digitRe := regexp.MustCompile(`\d+`)

	winning := digitRe.FindAllString(split[0], -1)
	present := digitRe.FindAllString(split[1], -1)

	winningNums := []int{}

	for _, v := range winning {
		value, _ := strconv.Atoi(v)
		winningNums = append(winningNums, value)
	}

	presentNums := []int{}

	for _, v := range present {
		value, _ := strconv.Atoi(v)
		presentNums = append(presentNums, value)
	}

	return &ScratchCard{
		Winning: winningNums,
		Present: presentNums,
		Count:   1,
	}
}

func (s *ScratchCard) getMatches() int {
	total := 0
	for _, v := range s.Present {
		if slices.Contains(s.Winning, v) {
			total++
		}
	}
	return total
}

func (s *ScratchCard) getValue() int {
	matches := s.getMatches()

	if matches == 0 {
		return 0
	}

	return int(math.Pow(2, float64(matches-1)))
}

func getTotalValueOfCards(input []string) int {

	var total int
	for _, line := range input {
		card := NewScratchCard(line)
		total += card.getValue()
	}

	return total
}

func getCopiesOfCards(cards []*ScratchCard) {
	if len(cards) == 0 {
		return
	}

	matches := cards[0].getMatches()
	// fmt.Printf("Current card: %d, copies: %d, matches: %d\n", 7-len(cards), cards[0].Count, matches)
	if matches > 0 {
		for i := 0; i < cards[0].Count; i++ {
			for j := 1; j < matches+1; j++ {
				if j >= len(cards) {
					break
				}
				// fmt.Println("Adding copy to ", 7-len(cards)+j)
				cards[j].Count++
			}
		}
	}

	getCopiesOfCards(cards[1:])
}

func getTotalCopiesOfCards(input []string) int {

	total := 0
	cards := []*ScratchCard{}
	for _, line := range input {
		cards = append(cards, NewScratchCard(line))
	}

	getCopiesOfCards(cards)

	for _, v := range cards {
		total += v.Count
	}

	return total
}
