package main

import "testing"

func TestGetSumOfPartNumbersFull(t *testing.T) {
	expected := 4361

	testInput := []string{
		"467..114..",
		"...*......",
		"..35..633.",
		"......#...",
		"617*......",
		".....+.58.",
		"..592.....",
		"......755.",
		"...$.*....",
		".664.598..",
	}

	if result := GetSumOfPartNumbers(testInput); result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestGetSumOfPartNumbersFull2(t *testing.T) {
	expected := 413

	testInput := []string{
		"12.......*..",
		"+.........34",
		".......-12..",
		"..78........",
		"..*....60...",
		"78..........",
		".......23...",
		"....90*12...",
		"............",
		"2.2......12.",
		".*.........*",
		"1.1.......56",
	}

	if result := GetSumOfPartNumbers(testInput); result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestGetSumOfPartNumbersFull3(t *testing.T) {
	expected := 925

	testInput := []string{
		"12.......*..",
		"+.........34",
		".......-12..",
		"..78........",
		"..*....60...",
		"78.........9",
		".5.....23..$",
		"8...90*12...",
		"............",
		"2.2......12.",
		".*.........*",
		"1.1..503+.56",
	}

	if result := GetSumOfPartNumbers(testInput); result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestGetSumOfPartNumbersNextToInLine(t *testing.T) {
	expected := 617

	testInput := []string{
		"617*......",
		".....+.58.",
	}

	if result := GetSumOfPartNumbers(testInput); result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestGetSumOfPartNumbersEmpty(t *testing.T) {
	expected := 0

	testInput := []string{
		"100",
		"200",
	}

	if result := GetSumOfPartNumbers(testInput); result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestGetSumOfPartNumbersOneLine(t *testing.T) {
	expected := 503

	testInput := []string{
		"503+",
	}

	if result := GetSumOfPartNumbers(testInput); result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestGetSumOfPartNumbersThreeValid(t *testing.T) {
	expected := 156

	testInput := []string{
		"....................",
		"..-52..52-..52..52..",
		"..................-.",
	}

	if result := GetSumOfPartNumbers(testInput); result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestGetSumOfPartNumbersMusifter(t *testing.T) {
	expected := 40

	testInput := []string{
		".......5......",
		"..7*..*.......",
		"...*13*.......",
		".......15.....",
	}

	if result := GetSumOfPartNumbers(testInput); result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestGetSumOfPartNumbersIsatisCrucifer(t *testing.T) {
	expected := 4

	testInput := []string{
		"........",
		".24..4..",
		"......*..",
	}

	if result := GetSumOfPartNumbers(testInput); result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestGetSumOfPartNumbersYazirro(t *testing.T) {
	expected := 2

	testInput := []string{
		".2.",
		"*.2",
	}

	if result := GetSumOfPartNumbers(testInput); result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

// Part 2

func TestGetTotalGearRatiosFull(y *testing.T) {
	expected := 467835

	testInput := []string{
		"467..114..",
		"...*......",
		"..35..633.",
		"......#...",
		"617*......",
		".....+.58.",
		"..592.....",
		"......755.",
		"...$.*....",
		".664.598..",
	}

	if result := GetTotalGearRatios(testInput); result != expected {
		y.Errorf("Expected %d, got %d", expected, result)
	}
}
