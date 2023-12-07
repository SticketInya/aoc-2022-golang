package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Advent of Code 2023 - Day 7")

	file, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	contents := strings.Split(string(file), "\n")

	totalWinnings := getTotalWinnings(contents)
	fmt.Printf("Total Winnings: %d\n", totalWinnings)

	totalWinningsWithHouseRules := getTotalWinningsWithHouseRules(contents)
	fmt.Printf("Total Winnings with House Rules: %d\n", totalWinningsWithHouseRules)

}

type CCHand struct {
	Cards string
	Bid   int
}

type CCSet struct {
	Hands []CCHand
}

func NewCCSet(input []string) *CCSet {
	hands := []CCHand{}
	for _, line := range input {
		parts := strings.Split(line, " ")
		if len(parts) != 2 {
			log.Fatal("Invalid input")
		}

		bid, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Fatal(err)
		}

		hands = append(hands, CCHand{
			Cards: parts[0],
			Bid:   bid,
		})
	}
	return &CCSet{
		Hands: hands,
	}
}

const (
	FiveOfAKind  = 7
	FourOfAKind  = 6
	FullHouse    = 5
	ThreeOfAKind = 4
	TwoPair      = 3
	OnePair      = 2
	HighCard     = 1
)

func (h *CCHand) getStrength(withHouseRule bool) int {
	charMap := map[string]int{}
	for _, char := range h.Cards {
		charMap[string(char)]++
	}

	if withHouseRule {
		if charMap["J"] > 0 {
			highestChar := ""
			for char, count := range charMap {
				if char == "J" {
					continue
				}

				if highestChar == "" {
					highestChar = char
					continue
				}

				if charMap[highestChar] < count {
					highestChar = char
				}
			}

			charMap[highestChar] += charMap["J"]
			delete(charMap, "J")
		}
	}

	switch len(charMap) {
	case 1:
		{
			return FiveOfAKind
		}
	case 2:
		{
			for _, count := range charMap {

				// Four of a kind
				if count == 4 {
					return FourOfAKind
				}

				// Full House
				if count == 3 {
					return FullHouse
				}
			}
		}
	case 3:
		{
			for _, count := range charMap {

				// Three of a kind
				if count == 3 {
					return ThreeOfAKind
				}

				// Two Pair
				if count == 2 {
					return TwoPair
				}
			}
		}
	case 4:
		{
			return OnePair
		}
	case 5:
		{
			return HighCard
		}
	default:
		{
			log.Fatal("Invalid hand")
		}
	}
	return 0
}

func (h *CCHand) getCardStrength(c byte, withHouseRule bool) int {
	cardOrder := []string{
		"2",
		"3",
		"4",
		"5",
		"6",
		"7",
		"8",
		"9",
		"T",
		"J",
		"Q",
		"K",
		"A",
	}

	if withHouseRule {
		cardOrder = []string{
			"J",
			"2",
			"3",
			"4",
			"5",
			"6",
			"7",
			"8",
			"9",
			"T",
			"Q",
			"K",
			"A",
		}
	}

	char := string(c)
	for i, card := range cardOrder {
		if card == char {
			return i
		}
	}

	log.Fatal("Invalid card")
	return 0
}

// Returns 1 if h is stronger than other, -1 if other is stronger, 0 if equal
func (h *CCHand) compare(other *CCHand, withHouseRule bool) int {
	if h.getStrength(withHouseRule) > other.getStrength(withHouseRule) {
		return 1
	}

	if h.getStrength(withHouseRule) < other.getStrength(withHouseRule) {
		return -1
	}

	for i := 0; i < len(h.Cards); i++ {
		if h.getCardStrength(h.Cards[i], withHouseRule) > h.getCardStrength(other.Cards[i], withHouseRule) {
			return 1
		}

		if h.getCardStrength(h.Cards[i], withHouseRule) < h.getCardStrength(other.Cards[i], withHouseRule) {
			return -1
		}
	}

	return 0
}

func (s *CCSet) getWinnings(withHouseRule bool) int {
	sort.Slice(s.Hands, func(i, j int) bool {
		return s.Hands[i].compare(&s.Hands[j], withHouseRule) != 1
	})

	total := 0
	for i, hand := range s.Hands {
		total += (i + 1) * hand.Bid
	}

	return total
}

func getTotalWinnings(input []string) int {
	game := NewCCSet(input)

	return game.getWinnings(false)
}

func getTotalWinningsWithHouseRules(input []string) int {
	game := NewCCSet(input)

	return game.getWinnings(true)
}
