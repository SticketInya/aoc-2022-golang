package main

import "testing"

func TestGetSumOfExtrapolatedValues(t *testing.T) {
	type test struct {
		input    []string
		expected int
	}

	cases := []test{
		{
			input: []string{
				"0 3 6 9 12 15",
			},
			expected: 18,
		},
		{
			input: []string{
				"1 3 6 10 15 21",
			},
			expected: 28,
		},
		{
			input: []string{
				"10 13 16 21 30 45",
			},
			expected: 68,
		},
		{
			input: []string{
				"6 12 17 21 24 26 27 27 26 24 21 17 12 6 1 9 18 28 39 51 64",
			},
			expected: -261858,
		},
		{
			input: []string{
				"8 -2 -17 -37 -52 -23 148 661 1876 4388 9121 17443 31304 53399 87358 137965 211408 315562 460307 657883 923284",
			},
			expected: 1274693,
		},
		{
			input: []string{
				"0 3 6 9 12 15",
				"1 3 6 10 15 21",
				"10 13 16 21 30 45",
			},
			expected: 114,
		},
	}

	for _, c := range cases {
		if result := getSumOfExtrapolatedValues(c.input); result != c.expected {
			t.Errorf("Expected %d, got %d", c.expected, result)
		}
	}
}

func TestGetSumOfPrevExtrapolatedValues(t *testing.T) {
	type test struct {
		input    []string
		expected int
	}

	cases := []test{
		{
			input: []string{
				"0 3 6 9 12 15",
			},
			expected: -3,
		},
		{
			input: []string{
				"1 3 6 10 15 21",
			},
			expected: 0,
		},
		{
			input: []string{
				"10 13 16 21 30 45",
			},
			expected: 5,
		},
		{
			input: []string{
				"0 3 6 9 12 15",
				"1 3 6 10 15 21",
				"10 13 16 21 30 45",
			},
			expected: 2,
		},
	}

	for _, c := range cases {
		if result := getSumOfPrevExtrapolatedValues(c.input); result != c.expected {
			t.Errorf("Expected %d, got %d", c.expected, result)
		}
	}
}

func TestGetNextDiff(t *testing.T) {
	type test struct {
		input    []int
		expected int
	}

	cases := []test{
		{
			input:    []int{0, 3, 6},
			expected: 3,
		},
		{
			input:    []int{1, 3, 6, 10, 15, 21},
			expected: 7,
		},
	}

	for _, c := range cases {
		if result := getNextDiff(c.input); result != c.expected {
			t.Errorf("Expected %d, got %d", c.expected, result)
		}
	}
}
