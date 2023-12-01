package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("AOC 2022 - Day 6")

	task_1()
	task_2()
}

func task_1() {
	file, err := os.Open("day_06/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("Packet marker at: %v\n", getMarker(line, 4))
	}
}

func task_2() {
	file, err := os.Open("day_06/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("Message marker at: %v\n", getMarker(line, 14))
	}
}

func getMarker(line string, treshold int) int {

	for i := treshold; i < len(line); i++ {
		charMap := map[rune]bool{}
		substring := line[i-treshold : i]
		allUnique := true

		for _, char := range substring {
			if charMap[char] {
				allUnique = false
			}
			charMap[char] = true
		}

		if allUnique {
			return i
		}

	}

	return -1
}
