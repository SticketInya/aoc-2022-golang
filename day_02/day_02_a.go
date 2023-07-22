package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("AOC 2022 - Day 2 A")

	task_1()
	task_2()

}

func task_1() {
	scores := map[string]int{
		"A X": 4,
		"A Y": 8,
		"A Z": 3,
		"B X": 1,
		"B Y": 5,
		"B Z": 9,
		"C X": 7,
		"C Y": 2,
		"C Z": 6,
	}

	fmt.Printf("T1 Total score: %v\n", calculateScores(scores))
}

func task_2() {
	scores := map[string]int{
		"A X": 3, // Rock loss -> Scissor 3 + loss 0 = 3
		"A Y": 4, // Rock draw -> Rock 1 + draw 3 = 4
		"A Z": 8, // Rock win -> Paper 2 + win 6 = 8
		"B X": 1, // Paper loss -> Rock 1 + loss 0 = 1
		"B Y": 5, // Paper draw -> Paper 2 + draw 3 = 5
		"B Z": 9, // Paper win -> Scissor 3 + win 6 = 9
		"C X": 2, // Scissor loss -> Paper 2 + loss 0 = 2
		"C Y": 6, // Scissor draw -> Scissor 3 + draw 3 = 6
		"C Z": 7, // Scissor win -> Rock 1 + win 6 = 7
	}

	fmt.Printf("T2 Total score: %v\n", calculateScores(scores))
}

func calculateScores(scores map[string]int) int {
	file, err := os.Open("day_02/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	total := 0
	for scanner.Scan() {
		total += scores[scanner.Text()]
	}
	return total
}
