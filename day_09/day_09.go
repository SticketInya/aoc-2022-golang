package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("AOC 2022 - Day 9")
	task_1()
	task_2()
}

type position struct {
	x, y int
}

type move struct {
	Direction string
	Steps     int
}

func task_1() {
	file, err := os.Open("day_09/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	head := position{0, 0}
	tail := position{0, 0}

	visitedMap := map[position]bool{}

	for scanner.Scan() {
		parts := strings.Fields(scanner.Text())
		steps, err := strconv.Atoi(parts[1])
		if err != nil {
			fmt.Printf("Failed to parse charachter: %v\n", parts[1])
			continue
		}

		// fmt.Printf("direction: %s, steps:%v\n", parts[0], parts[1])
		makeMove(&head, &tail, visitedMap, move{
			Direction: parts[0],
			Steps:     steps,
		})
	}

	counter := 0

	for _, visited := range visitedMap {
		if visited {
			counter++
		}
	}
	println("Visited positions count: ", counter)
}

func task_2() {
	file, err := os.Open("day_09/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	rope := make([]position, 10)

	visitedMap := map[position]bool{}

	for scanner.Scan() {
		parts := strings.Fields(scanner.Text())
		steps, err := strconv.Atoi(parts[1])
		if err != nil {
			fmt.Printf("Failed to parse charachter: %v\n", parts[1])
			continue
		}

		for i := 0; i < steps; i++ {
			movePosition(&rope[0], parts[0], 1)

			for i := range rope[:len(rope)-1] {
				adjustTail(&rope[i], &rope[i+1])
				visitedMap[rope[9]] = true
			}
		}

	}

	counter := 0

	for _, visited := range visitedMap {
		if visited {
			counter++
		}
	}
	println("Visited positions count task 2: ", counter)
}

func printCurrentState(head *position, tail *position) {
	for i := 0; i < 6; i++ {
		for j := 0; j < 6; j++ {
			if head.x == i && head.y+4 == j {
				fmt.Printf("H ")
				continue
			}

			if tail.x == i && tail.y+4 == j {
				fmt.Printf("T ")
				continue
			}

			fmt.Printf(". ")

		}
		fmt.Printf("\n")
	}
}

func makeMove(head *position, tail *position, visitedMap map[position]bool, move move) {

	for i := 0; i < move.Steps; i++ {
		// printCurrentState(head, tail)
		visitedMap[*tail] = true
		movePosition(head, move.Direction, 1)
		// fmt.Printf("Tail: %v\n", tail)
		// fmt.Printf("Head: %v\n", head)
		adjustTail(head, tail)

	}
}

func adjustTail(head *position, tail *position) {
	if !arePositionsTouching(head, tail) {

		if head.x == tail.x {
			if head.y > tail.y {
				tail.y++
				return
			}

			tail.y--
			return
		}

		if head.y == tail.y {
			if head.x > tail.x {
				tail.x++
				return
			}

			tail.x--

			return
		}

		if head.x > tail.x && head.y > tail.y {
			tail.x++
			tail.y++
			return
		}

		if head.x > tail.x && head.y < tail.y {
			tail.x++
			tail.y--
			return

		}

		if head.x < tail.x && head.y < tail.y {
			tail.x--
			tail.y--
			return

		}

		tail.x--
		tail.y++

	}
}

func movePosition(item *position, direction string, step int) {
	switch direction {
	case "R":
		{
			item.x += step
		}
	case "L":
		{
			item.x -= step
		}
	case "U":
		{
			item.y -= step
		}
	case "D":
		{
			item.y += step
		}
	default:
		{
			fmt.Println("Invalid direction:", direction)
		}
	}
}

func arePositionsTouching(head *position, tail *position) bool {
	areTouching := math.Abs(float64(head.x)-float64(tail.x)) > 1 || math.Abs(float64(head.y)-float64(tail.y)) > 1

	// println("Are touching? ", !areTouching, head.x-tail.x, head.y-tail.y)

	return !areTouching
}
