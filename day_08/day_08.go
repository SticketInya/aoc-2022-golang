package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	fmt.Println("AOC 2022 - Day 8")

	task_1()
}

func task_1() {
	file, err := os.Open("day_08/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	forest := [][]int{}

	for scanner.Scan() {
		line := scanner.Text()

		subArray := []int{}

		for _, c := range line {
			parsed, err := strconv.Atoi(string(c))
			if err != nil {
				fmt.Println("Failed to parse charachter:", string(c))
				continue
			}

			subArray = append(subArray, parsed)
		}
		forest = append(forest, subArray)
	}

	fmt.Printf("total visible:%v\n", getTotalVisibleTreeCount(forest))
	fmt.Printf("max scenic score:%v\n", getMaxScenicScore(forest))
}

func getTotalVisibleTreeCount(forest [][]int) int {
	total := (len(forest) * 2) + (len(forest[0]) * 2) - 4

	for i := 1; i < len(forest)-1; i++ {
		for j := 1; j < len(forest[i])-1; j++ {
			if isTreeVisible(i, j, forest) {
				total++
			}
		}
	}

	return total
}

func isTreeVisible(x int, y int, matrix [][]int) bool {
	yAxis := []int{}
	xAxis := matrix[x]

	for _, sa := range matrix {
		yAxis = append(yAxis, sa[y])
	}

	// fmt.Printf("x:%v %v\n", xAxis, isTreeVisibleInRow(y, xAxis))
	// fmt.Printf("y:%v %v\n", yAxis, isTreeVisibleInRow(x, yAxis))

	return isTreeVisibleInRow(y, xAxis) || isTreeVisibleInRow(x, yAxis)
}

func isTreeVisibleInRow(position int, row []int) bool {
	height := row[position]
	var fromLeft, fromRight bool

	for i := 0; i < position; i++ {
		if row[i] >= height {
			fromLeft = true
			break
		}
	}

	for i := len(row) - 1; i > position; i-- {
		if row[i] >= height {
			fromRight = true
			break
		}
	}

	return !fromLeft || !fromRight
}

func getMaxScenicScore(forest [][]int) int {
	max := getScenicScoreForTree(0, 0, forest)

	for i := 0; i < len(forest); i++ {
		for j := 0; j < len(forest[i]); j++ {
			newScenicScore := getScenicScoreForTree(i, j, forest)
			// println(newScenicScore)
			if newScenicScore > max {
				max = newScenicScore
			}
		}
	}

	return max
}

func getScenicScoreForTree(x int, y int, matrix [][]int) int {
	yAxis := []int{}
	xAxis := matrix[x]

	for _, sa := range matrix {
		yAxis = append(yAxis, sa[y])
	}

	// fmt.Printf("x:%v %v\n", xAxis, isTreeVisibleInRow(y, xAxis))
	// fmt.Printf("y:%v %v\n", yAxis, isTreeVisibleInRow(x, yAxis))
	a, b := getScenicPairForRow(y, xAxis)
	c, d := getScenicPairForRow(x, yAxis)

	// println("Scenic scores for:", matrix[x][y], a, b, c, d)

	return a * b * c * d
}

func getScenicPairForRow(position int, row []int) (int, int) {
	height := row[position]
	fromLeft := 0
	fromRight := 0

	for i := position - 1; i >= 0; i-- {
		fromLeft++
		if row[i] >= height {
			break
		}
	}

	for i := position + 1; i < len(row); i++ {
		fromRight++

		if row[i] >= height {
			break
		}
	}

	return fromLeft, fromRight
}
