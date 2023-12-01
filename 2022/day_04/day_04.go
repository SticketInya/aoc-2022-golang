package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("AOC 2022 - Day 4")

	task_1()
	task_2()
}

func task_1() {
	file, err := os.Open("day_04/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	counter := 0
	for scanner.Scan() {
		var startFirst, endFirst, startSecond, endSecond int
		fmt.Sscanf(scanner.Text(), "%d-%d,%d-%d", &startFirst, &endFirst, &startSecond, &endSecond)

		if startFirst >= startSecond && endSecond >= endFirst || startSecond >= startFirst && endFirst >= endSecond {
			counter++
		}
	}

	println("Task 1 Total: ", counter)
}

func task_2() {
	file, err := os.Open("day_04/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	counter := 0
	for scanner.Scan() {
		var startFirst, endFirst, startSecond, endSecond int
		fmt.Sscanf(scanner.Text(), "%d-%d,%d-%d", &startFirst, &endFirst, &startSecond, &endSecond)

		if startFirst <= startSecond && endFirst >= startSecond || startSecond <= startFirst && endSecond >= startFirst {
			counter++
		}
	}

	println("Task 2 Total: ", counter)
}
