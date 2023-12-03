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
	fmt.Println("Advent of Code 2023 - Day 3")

	file, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	contents := strings.Split(string(file), "\n")

	totalPartNumbers := GetSumOfPartNumbers(contents)
	fmt.Println("Task 1:", totalPartNumbers)

	totalGearRatios := GetTotalGearRatios(contents)
	fmt.Println("Task 2:", totalGearRatios)
}

type Coordinate struct {
	LineIndex int
	Start     int
	End       int
	Value     string
	Used      bool
}

func GetSumOfPartNumbers(input []string) int {
	digitRe := regexp.MustCompile(`\d`)

	numbers := []Coordinate{}
	specials := []Coordinate{}

	for i, line := range input {
		nbrBuffer := ""
		for j, char := range line {
			c := string(char)
			isDigit := digitRe.MatchString(c)
			if isDigit {
				nbrBuffer += c
			}

			if j == len(line)-1 && nbrBuffer != "" {
				if isDigit {
					numbers = append(numbers, Coordinate{
						LineIndex: i,
						Start:     j + 1 - len(nbrBuffer),
						End:       j,
						Value:     nbrBuffer})
				} else {
					numbers = append(numbers, Coordinate{
						LineIndex: i,
						Start:     j - len(nbrBuffer),
						End:       j - 1,
						Value:     nbrBuffer})
				}

				nbrBuffer = ""
			}

			if !isDigit && nbrBuffer != "" {
				numbers = append(numbers, Coordinate{
					LineIndex: i,
					Start:     j - len(nbrBuffer),
					End:       j - 1,
					Value:     nbrBuffer})
				nbrBuffer = ""
			}

			if !isDigit && char != '.' {
				specials = append(specials, Coordinate{
					LineIndex: i,
					Start:     j,
					End:       j,
					Value:     c})
			}

		}
	}

	total := 0
	for _, special := range specials {
		for j, number := range numbers {
			if number.Used {
				continue
			}

			// Check if a number is in the same line and next to a special character
			if special.LineIndex == number.LineIndex && (special.Start == number.End+1 || special.Start == number.Start-1) {
				numbers[j].Used = true
				parsed, _ := strconv.Atoi(number.Value)
				total += parsed

				continue
			}

			// Check if a number is above the line and next to a special character
			if special.LineIndex == number.LineIndex-1 {
				if special.Start >= number.Start-1 && special.Start <= number.End+1 {
					numbers[j].Used = true
					parsed, _ := strconv.Atoi(number.Value)
					total += parsed

					continue
				}
			}

			// Check if a number is below the line and next to a special character
			if special.LineIndex == number.LineIndex+1 {
				if special.Start >= number.Start-1 && special.Start <= number.End+1 {
					numbers[j].Used = true
					parsed, _ := strconv.Atoi(number.Value)
					total += parsed

					continue
				}
			}

		}

	}

	return total
}

func GetTotalGearRatios(input []string) int {
	digitRe := regexp.MustCompile(`\d`)

	numbers := []Coordinate{}
	specials := []Coordinate{}

	for i, line := range input {
		nbrBuffer := ""
		for j, char := range line {
			c := string(char)
			isDigit := digitRe.MatchString(c)
			if isDigit {
				nbrBuffer += c
			}

			if j == len(line)-1 && nbrBuffer != "" {
				if isDigit {
					numbers = append(numbers, Coordinate{
						LineIndex: i,
						Start:     j + 1 - len(nbrBuffer),
						End:       j,
						Value:     nbrBuffer})
				} else {
					numbers = append(numbers, Coordinate{
						LineIndex: i,
						Start:     j - len(nbrBuffer),
						End:       j - 1,
						Value:     nbrBuffer})
				}

				nbrBuffer = ""
			}

			if !isDigit && nbrBuffer != "" {
				numbers = append(numbers, Coordinate{
					LineIndex: i,
					Start:     j - len(nbrBuffer),
					End:       j - 1,
					Value:     nbrBuffer})
				nbrBuffer = ""
			}

			if !isDigit && char != '.' {
				specials = append(specials, Coordinate{
					LineIndex: i,
					Start:     j,
					End:       j,
					Value:     c})
			}

		}
	}

	total := 0
	for _, special := range specials {
		closeByNumbers := []int{}

		for _, number := range numbers {

			// Check if a number is in the same line and next to a special character
			if special.LineIndex == number.LineIndex && (special.Start == number.End+1 || special.Start == number.Start-1) {
				parsed, _ := strconv.Atoi(number.Value)
				closeByNumbers = append(closeByNumbers, parsed)

				continue
			}

			// Check if a number is above the line and next to a special character
			if special.LineIndex == number.LineIndex-1 {
				if special.Start >= number.Start-1 && special.Start <= number.End+1 {
					parsed, _ := strconv.Atoi(number.Value)
					closeByNumbers = append(closeByNumbers, parsed)

					continue
				}
			}

			// Check if a number is below the line and next to a special character
			if special.LineIndex == number.LineIndex+1 {
				if special.Start >= number.Start-1 && special.Start <= number.End+1 {
					parsed, _ := strconv.Atoi(number.Value)
					closeByNumbers = append(closeByNumbers, parsed)

					continue
				}
			}

		}

		if len(closeByNumbers) == 2 {
			total += closeByNumbers[0] * closeByNumbers[1]
		}

	}

	return total
}
