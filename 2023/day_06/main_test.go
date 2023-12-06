package main

import "testing"

func TestGetProductOfPossibleWaysToWinFul(t *testing.T) {
	expected := 288
	input := []string{
		"Time:      7  15   30",
		"Distance:  9  40  200",
	}

	if result := getProductOfPossibleWaysToWin(input); result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestGetProductOfPossibleWaysToWinKerningFul(t *testing.T) {
	expected := 71503
	input := []string{
		"Time:      7  15   30",
		"Distance:  9  40  200",
	}

	if result := getProductOfPossibleWaysToWinWithKerning(input); result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}
