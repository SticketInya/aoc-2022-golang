package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	fmt.Println("Advent of code 2023 - Day 1")

	part1()
	part2()

}

func part1() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	total := 0

	for scanner.Scan() {
		line := scanner.Text()

		re := regexp.MustCompile(`\d`)
		numbers := re.FindAllString(line, -1)

		calibration, err := strconv.Atoi(numbers[0] + numbers[len(numbers)-1])

		if err != nil {
			log.Fatal(err)
		}

		total += calibration

	}

	fmt.Println(total)
}

func part2() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	total := 0

	for scanner.Scan() {
		line := scanner.Text()

		re := regexp.MustCompile(`(\d|one|two|three|four|five|six|seven|eight|nine)`)
		first := re.FindAllString(line, -1)[0]
		last := regexp.MustCompile(`.*` + re.String()).FindStringSubmatch(line)[1]

		calibration, err := strconv.Atoi(fmt.Sprintf("%d%d", matchNumberStrings(first), matchNumberStrings(last)))
		if err != nil {
			log.Fatal(err)
		}
		total += calibration

	}

	fmt.Println(total)
}

func matchNumberStrings(numberLike string) int {
	if number, err := strconv.Atoi(numberLike); err == nil {
		return number
	}

	switch numberLike {
	case "one":
		return 1
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	default:
		log.Fatal("Unknown number")
	}
	return 0
}
