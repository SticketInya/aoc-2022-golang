package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Advent of Code 2023 - Day 5")

	file, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	contents := strings.Split(string(file), "\n")

	firstSeed := getFirstSeedToPlant(contents)
	fmt.Println("First seed to plant:", firstSeed)

	firstSeedFromPair := getFirstSeedToPlantFromPairs(contents)
	fmt.Println("First seed to plant from pair:", firstSeedFromPair)
}

type RangeMap struct {
	Source      []int
	Destination []int
	Range       []int
}

// Returns the mapped value for the source and if it is within the map
func (rm *RangeMap) getDestinationForInstance(source, ci int) (int, bool) {
	if source > (rm.Source[ci]+rm.Range[ci]-1) || source < rm.Source[ci] {
		return source, false
	}

	return rm.Destination[ci] + (source - rm.Source[ci]), true
}

func (rm *RangeMap) getDestination(source int) int {
	for i := range rm.Source {
		dest, ok := rm.getDestinationForInstance(source, i)
		if ok {
			return dest
		}
	}

	return source
}

func NewRangeMap(input []string) *RangeMap {
	titleRe := regexp.MustCompile(`\w+:`)
	digitRe := regexp.MustCompile(`\d+`)
	rm := RangeMap{}

	for _, line := range input {
		if titleRe.MatchString(line) {
			continue
		}

		numbers := digitRe.FindAllString(line, -1)
		if len(numbers) != 3 {
			log.Fatal("Invalid input", line)
		}
		source, _ := strconv.Atoi(numbers[1])
		destination, _ := strconv.Atoi(numbers[0])
		r, _ := strconv.Atoi(numbers[2])

		rm.Source = append(rm.Source, source)
		rm.Destination = append(rm.Destination, destination)
		rm.Range = append(rm.Range, r)
	}

	return &rm
}

func parseSections(input []string) [][]string {
	sections := [][]string{
		{},
	}

	var currentIndex int
	for _, line := range input {
		if line == "" {
			currentIndex++
			sections = append(sections, []string{})
			continue
		}

		sections[currentIndex] = append(sections[currentIndex], line)
	}

	if len(sections) != 8 {
		log.Fatal("Invalid input. Failed to properly parse sections.")
	}

	return sections
}

func parseSeeds(input string) []int {
	digitRe := regexp.MustCompile(`\d+`)

	matches := digitRe.FindAllString(input, -1)
	seeds := []int{}

	for _, m := range matches {
		seed, _ := strconv.Atoi(m)
		seeds = append(seeds, seed)
	}
	return seeds
}

func parseSeedPairs(input string) []int {
	digitRe := regexp.MustCompile(`\d+`)

	matches := digitRe.FindAllString(input, -1)
	seeds := []int{}

	for _, m := range matches {
		seed, _ := strconv.Atoi(m)
		seeds = append(seeds, seed)
	}

	if len(seeds)%2 != 0 {
		log.Fatal("Invalid input. Seed count is not even.")
	}

	seedPairs := []int{}
	for i := 0; i < len(seeds); i += 2 {
		for j := 0; j < seeds[i+1]; j++ {
			seedPairs = append(seedPairs, seeds[i]+j)
		}
	}

	return seedPairs
}

func getFirstSeedToPlant(input []string) int {
	sections := parseSections(input)
	seeds := parseSeeds(sections[0][0])

	seedMaps := []*RangeMap{}
	for i := 1; i < len(sections); i++ {
		seedMaps = append(seedMaps, NewRangeMap(sections[i]))
	}

	for i := range seeds {
		// fmt.Println("Results for seed ", i+1)
		for _, seedMap := range seedMaps {
			seeds[i] = seedMap.getDestination(seeds[i])
			// fmt.Println(seeds[i])
		}
	}

	// fmt.Println("Planted seeds:", seeds)
	return slices.Min(seeds)
}

// Takes about 8m 32s for the given input
func getFirstSeedToPlantFromPairs(input []string) int {
	sections := parseSections(input)
	seeds := parseSeedPairs(sections[0][0])

	// log.Println("Seeds:", seeds)

	seedMaps := []*RangeMap{}
	for i := 1; i < len(sections); i++ {
		seedMaps = append(seedMaps, NewRangeMap(sections[i]))
	}

	for i := range seeds {
		// fmt.Println("Results for seed ", i+1)
		for _, seedMap := range seedMaps {
			seeds[i] = seedMap.getDestination(seeds[i])
			// fmt.Println(seeds[i])
		}
	}

	// fmt.Println("Planted seeds:", seeds)
	return slices.Min(seeds)
}
