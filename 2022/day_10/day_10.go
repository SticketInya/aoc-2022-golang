package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("AOC 2022 - Day 10")
	task_1()
	task_2()
}

func task_1() {
	file, err := os.Open("day_10/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	counter := 0
	currentTotal := 1
	signalStrength := 0

	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())

		if fields[0] == "noop" {
			counter++
			signalStrength = checkCycle(counter, currentTotal, signalStrength)

		}

		if fields[0] == "addx" {
			value, err := strconv.Atoi(fields[1])
			if err != nil {
				fmt.Println("Failed to parse character ", fields[1])
				continue
			}

			counter += 1
			signalStrength = checkCycle(counter, currentTotal, signalStrength)
			counter += 1
			signalStrength = checkCycle(counter, currentTotal, signalStrength)
			currentTotal += value
		}
	}

	fmt.Printf("Total signal strength: %v\n", signalStrength)

}
func checkCycle(counter int, currentTotal int, signalStrength int) int {
	if (counter-20)%40 == 0 {
		println("Adding", counter*currentTotal, counter, currentTotal)
		return signalStrength + counter*currentTotal
	}
	return signalStrength
}

func task_2() {
	file, err := os.Open("day_10/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	counter := 0
	currentTotal := 1

	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())

		if fields[0] == "noop" {
			counter += 1
			drawCycle(counter, currentTotal)
		}

		if fields[0] == "addx" {
			value, err := strconv.Atoi(fields[1])
			if err != nil {
				fmt.Println("Failed to parse character ", fields[1])
				continue
			}

			counter += 1
			drawCycle(counter, currentTotal)
			counter += 1
			drawCycle(counter, currentTotal)
			currentTotal += value
		}
	}

}

func drawCycle(counter int, currentTotal int) {
	currentIndex := counter % 40

	if currentIndex == currentTotal || currentIndex == currentTotal+2 || currentIndex == currentTotal+1 {
		fmt.Printf("#")
	} else {
		fmt.Printf(".")
	}
	// println(currentIndex, currentTotal)

	if currentIndex == 0 {
		fmt.Printf("\n")
	}
}
