package main

import (
	"testing"
)

// TestGetPowerOfBag tests the getPowerOfBag function.
func TestGetPowerOfBag(t *testing.T) {
	expected := 24

	testBag := Bag{
		Red:   2,
		Green: 3,
		Blue:  4,
	}

	if power := testBag.getPowerOfBag(); power != expected {
		t.Errorf("Expected %d, got %d", expected, power)
	}
}

func TestGetSumOfPossibleGameIndeces(t *testing.T) {
	expected := 8

	testBag := Bag{
		Red:   12,
		Green: 13,
		Blue:  14,
	}

	testInput := `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`

	if sum := GetSumOfPossibleGameIndeces(testInput, &testBag); sum != expected {
		t.Errorf("Expected %d, got %d", expected, sum)
	}
}

func TestGetSumOfPowerOfGames(t *testing.T) {
	expected := 2286

	testInput := `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`

	if sum := GetSumOfPowerOfGames(testInput); sum != expected {
		t.Errorf("Expected %d, got %d", expected, sum)
	}

}

func TestGetGameIndex(t *testing.T) {
	expected := 3
	testInput := "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red"

	if index, err := GetGameIndex(testInput); index != expected || err != nil {
		t.Errorf("Expected %d, got %d", expected, index)
	}
}
