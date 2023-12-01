package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Println("AOC 2022 - Day 5")

	task_1()
	task_2()
}

type move struct {
	Amount int
	From   int
	To     int
}

func task_1() {
	file, err := os.Open("day_05/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	crates, moves := getCratesAndMoves(scanner)

	for _, move := range moves {
		makeMove(&crates, move)
	}

	println("Result: ", getTopCrates(&crates))
}

func task_2() {
	file, err := os.Open("day_05/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	crates, moves := getCratesAndMoves(scanner)

	for _, move := range moves {
		makeMoveByGroup(&crates, move)
	}

	println("Result 2: ", getTopCrates(&crates))
}

func getCratesAndMoves(scanner *bufio.Scanner) (crates [9]string, moves []move) {
	readingMoves := false

	for scanner.Scan() {

		line := scanner.Text()

		if line == "" {
			readingMoves = true
			continue
		}

		// Ignore crate numbering row
		if !readingMoves && strings.Contains(line, "[") {
			for i, c := range line {
				char := string(c)
				if char == " " || char == "[" || char == "]" {
					continue
				}
				crates[i/4] = fmt.Sprintf("%s%s", crates[i/4], char)

			}

		}

		if readingMoves {
			var amount, from, to int
			fmt.Sscanf(line, "move %d from %d to %d", &amount, &from, &to)

			moves = append(moves, move{
				Amount: amount,
				From:   from - 1,
				To:     to - 1,
			})
		}
	}

	return crates, moves
}

func makeMove(crates *[9]string, move move) {

	for i := 0; i < move.Amount; i++ {
		crateToMove := crates[move.From][0:1]
		crates[move.To] = fmt.Sprintf("%s%s", crateToMove, crates[move.To])
		crates[move.From] = crates[move.From][1:]
	}

}

func makeMoveByGroup(crates *[9]string, move move) {
	fromLeft := crates[move.From][:move.Amount]
	fromRemainder := crates[move.From][move.Amount:]
	crates[move.To] = fmt.Sprintf("%s%s", fromLeft, crates[move.To])
	crates[move.From] = fromRemainder
}

func getTopCrates(crates *[9]string) string {
	endString := ""
	for _, crate := range crates {
		if crate == "" {
			continue
		}
		endString = fmt.Sprintf("%s%s", endString, crate[0:1])
	}
	return endString
}
