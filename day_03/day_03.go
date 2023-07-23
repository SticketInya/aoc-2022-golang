package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Println("AOC 2022 - Day 3")
	println("a", int('a'))
	println("A", int('A'))

	task1()
	task2()
}

func task1() {

	file, err := os.Open("day_03/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	total := 0

	for scanner.Scan() {
		line := scanner.Text()
		chars := strings.Split(line, "")
		halfPoint := len(chars) / 2
		secondCompartment := strings.Join(chars[halfPoint:], "")
		var duplicateItem int
		for i := 0; i < halfPoint; i++ {
			if strings.Contains(secondCompartment, chars[i]) {
				letter := int(rune(chars[i][0]))
				// println("Found duplicate char: ", chars[i])
				duplicateItem = getLetterValue(letter)
			}
		}
		total += duplicateItem
	}

	println("Task 1 total: ", total)
}

func task2() {

	file, err := os.Open("day_03/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	total := 0
	groupItems := [3]string{}
	currentGroupCounter := 0

	for scanner.Scan() {
		line := scanner.Text()
		groupItems[currentGroupCounter] = line
		currentGroupCounter++

		if currentGroupCounter != 3 {
			continue
		}

		badge := 0
		chars := strings.Split(groupItems[0], "")

		for _, char := range chars {
			isInOtherBags := strings.Contains(groupItems[1], char) && strings.Contains(groupItems[2], char)

			if isInOtherBags {
				badge = int(char[0])
			}
		}
		// println("Found badge", string(badge))
		total += getLetterValue(badge)
		currentGroupCounter = 0
	}

	println("Task 2 total: ", total)
}

func getLetterValue(letter int) int {
	const lowercaseOffset = 96
	const uppercaseOffset = 38

	if letter < 97 {
		return letter - uppercaseOffset
	} else {
		return letter - lowercaseOffset
	}
}
