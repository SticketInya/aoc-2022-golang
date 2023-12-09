package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Advent of Code 2023 - Day 9")
	file, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	contents := strings.Split(string(file), "\n")

	sumOfExtrapolatedValues := getSumOfExtrapolatedValues(contents)
	fmt.Println("Sum of extrapolated values:", sumOfExtrapolatedValues)

	sumOfPrev := getSumOfPrevExtrapolatedValues(contents)
	fmt.Println("Sum of prev extrapolated values:", sumOfPrev)
}

type History struct {
	Values []int
}

type Reading struct {
	Reports []History
}

func NewReading(input []string) *Reading {
	reading := Reading{}
	for _, line := range input {
		matches := strings.Split(line, " ")
		History := History{
			Values: []int{},
		}
		for _, match := range matches {
			parsed, _ := strconv.Atoi(match)
			History.Values = append(History.Values, parsed)
		}
		reading.Reports = append(reading.Reports, History)
	}
	return &reading
}

func getNextDiff(nums []int) int {
	if len(nums) == 1 {
		return 0
	}

	diffs := []int{}
	isAllZeros := true
	for i := 1; i < len(nums); i++ {
		diff := nums[i] - nums[i-1]
		diffs = append(diffs, diff)
		if diff != 0 {
			isAllZeros = false
		}
	}
	if isAllZeros {
		return 0
	}

	return diffs[len(diffs)-1] + getNextDiff(diffs)
}

func getPrevDiff(nums []int) int {
	if len(nums) == 1 {
		return 0
	}

	diffs := []int{}
	isAllZeros := true
	for i := 1; i < len(nums); i++ {
		diff := nums[i] - nums[i-1]
		diffs = append(diffs, diff)
		if diff != 0 {
			isAllZeros = false
		}
	}
	if isAllZeros {
		return 0
	}

	return diffs[0] - getPrevDiff(diffs)
}

func (r *Reading) getSumOfNextReadingValues() int {
	sum := 0
	for _, history := range r.Reports {
		nv := history.Values[len(history.Values)-1] + getNextDiff(history.Values)
		sum += nv
	}
	return sum
}

func (r *Reading) getSumOfPrevReadingValues() int {
	sum := 0
	for _, history := range r.Reports {
		nv := history.Values[0] - getPrevDiff(history.Values)
		sum += nv
	}
	return sum
}

func getSumOfExtrapolatedValues(input []string) int {
	reading := NewReading(input)

	return reading.getSumOfNextReadingValues()
}

func getSumOfPrevExtrapolatedValues(input []string) int {
	reading := NewReading(input)

	return reading.getSumOfPrevReadingValues()
}
