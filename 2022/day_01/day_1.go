package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main(){
	fmt.Println("AOC 2022 - Day 1")

	file, err := os.Open("day_01/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	totalFood := []int{}
	currentTotal := 0

	for scanner.Scan(){
		line := scanner.Text()
		if line == ""{
			totalFood = append(totalFood, currentTotal)
			currentTotal =0
			continue
		}
		
		parsed, err := strconv.Atoi(line)
		if err != nil {
			fmt.Printf("failed to parse line %s with err %v",line,err)
			continue
		}

		currentTotal+=parsed
	}
	
	sort.Ints(totalFood)
	fmt.Printf("The top total: %v\n",totalFood[len(totalFood)-1])

	topThree := totalFood[len(totalFood)-3:]
	total := 0

	for _, val := range topThree{
		total+=val
	}

	fmt.Printf("Total of top three: %v\n",total)
}
	