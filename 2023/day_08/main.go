package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	fmt.Println("Advent of Code 2023 - Day 8")

	file, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	contents := strings.Split(string(file), "\n")

	steps := getStepsToReachDestination(contents)
	fmt.Printf("Steps to reach destination: %d\n", steps)

	ghostSteps := getGhostStepsToReachDestination(contents)
	fmt.Printf("Steps to reach ghost destination: %d\n", ghostSteps)
}

type Node struct {
	Value string
	Left  *Node
	Right *Node
}

func CreateNodes(input []string) map[string]*Node {
	nodeMap := map[string]*Node{}

	for _, line := range input {
		parts := strings.Split(line, " ")
		value := parts[0]

		nodeMap[value] = &Node{
			Value: value,
		}
	}

	return nodeMap
}

func getStepsToReachDestination(input []string) int {
	instructions := input[0]
	nodeMap := CreateNodes(input[2:])
	var start *Node

	re := regexp.MustCompile(`[A-Z]{3}`)

	for _, line := range input[2:] {
		matches := re.FindAllString(line, -1)
		key, left, right := matches[0], matches[1], matches[2]
		nodeMap[key].Left = nodeMap[left]
		nodeMap[key].Right = nodeMap[right]

		if key == "AAA" {
			start = nodeMap[key]
		}
	}

	if start == nil {
		log.Fatal("Start node not found")
	}

	// fmt.Printf("Start node: %s\n", start.Value)
	// fmt.Printf("End node: %s\n", nodeMap["ZZZ"].Value)
	// fmt.Printf("Instruction: %s\n", instructions)
	// fmt.Println("Nodes:")
	// for _, node := range nodeMap {
	// 	fmt.Printf("%s -> %s, %s\n", node.Value, node.Left.Value, node.Right.Value)
	// }

	steps := 0
	cn := start
	for cn.Value != "ZZZ" {
		ci := instructions[steps%(len(instructions))]
		switch ci {
		case 'L':
			cn = cn.Left
		case 'R':
			cn = cn.Right
		default:
			log.Fatal("Invalid instruction")
		}
		steps++
	}

	return steps
}

func findGCD(a, b int) int {
	if b == 0 {
		return a
	}
	return findGCD(b, a%b)
}

func findLCM(a, b int, integers ...int) int {
	multiple := a * b / findGCD(a, b)

	for _, integer := range integers {
		multiple = findLCM(multiple, integer)
	}

	return multiple
}

func getGhostStepsToReachDestination(input []string) int {
	instructions := input[0]
	nodeMap := CreateNodes(input[2:])
	starts := []*Node{}

	re := regexp.MustCompile(`[0-9A-Z]{3}`)

	for _, line := range input[2:] {
		matches := re.FindAllString(line, -1)
		key, left, right := matches[0], matches[1], matches[2]
		nodeMap[key].Left = nodeMap[left]
		nodeMap[key].Right = nodeMap[right]

		if len(key) > 2 && key[2] == 'A' {
			starts = append(starts, nodeMap[key])
		}
	}

	if len(starts) == 0 {
		log.Fatal("Start node not found")
	}

	steps := 0
	endSteps := []int{}
	var reachedEnd bool
	for !reachedEnd {

		ci := instructions[steps%(len(instructions))]

		for i, node := range starts {
			if node == nil {
				continue
			}

			switch ci {
			case 'L':
				starts[i] = starts[i].Left
			case 'R':
				starts[i] = starts[i].Right
			default:
				log.Fatal("Invalid instruction")
			}
		}
		steps++

		for i, node := range starts {
			if node == nil {
				continue
			}

			if node.Value[2] == 'Z' {
				endSteps = append(endSteps, steps)

				starts = append(starts[:i], starts[i+1:]...)
			}
		}

		if len(starts) == 0 {
			reachedEnd = true
		}

	}
	// fmt.Printf("Steps: %v\n", endSteps)

	return findLCM(endSteps[0], endSteps[1], endSteps[2:]...)
}
