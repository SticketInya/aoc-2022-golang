package main

import "testing"

func TestGetTotalWinnings(t *testing.T) {
	type test struct {
		input  []string
		output int
	}

	var testGetWinningsCases = []test{
		{
			input: []string{
				"32T3K 765",
				"T55J5 684",
				"KK677 28",
				"KTJJT 220",
				"QQQJA 483",
			},
			output: 6440,
		},
	}

	for _, c := range testGetWinningsCases {
		if result := getTotalWinnings(c.input); result != c.output {
			t.Errorf("GetWinnings(%v) == %d, expected %d", c.input, result, c.output)
		}
	}

}

func TestGetTotalWinningsWithHouseRules(t *testing.T) {
	type test struct {
		input  []string
		output int
	}

	var testGetWinningsCases = []test{
		{
			input: []string{
				"32T3K 765",
				"T55J5 684",
				"KK677 28",
				"KTJJT 220",
				"QQQJA 483",
			},
			output: 5905,
		},
	}

	for _, c := range testGetWinningsCases {
		if result := getTotalWinningsWithHouseRules(c.input); result != c.output {
			t.Errorf("GetWinnings(%v) == %d, expected %d", c.input, result, c.output)
		}
	}

}

func TestGetHandStrength(t *testing.T) {

	type test struct {
		input  string
		output int
	}

	var testGetHandStrengthCases = []test{
		{
			input:  "AAAAA",
			output: 7,
		},
		{
			input:  "AA8AA",
			output: 6,
		},
		{
			input:  "23332",
			output: 5,
		},
		{
			input:  "TTT98",
			output: 4,
		},
		{
			input:  "23432",
			output: 3,
		},
		{
			input:  "A23A4",
			output: 2,
		},
		{
			input:  "23456",
			output: 1,
		},
	}

	for _, c := range testGetHandStrengthCases {
		hand := CCHand{
			Cards: c.input,
			Bid:   0,
		}
		if result := hand.getStrength(false); result != c.output {
			t.Errorf("GetHandStrength(%v) == %d, expected %d", c.input, result, c.output)
		}
	}

}

func TestGetHandStrengthWithHouseRule(t *testing.T) {

	type test struct {
		input  string
		output int
	}

	var testGetHandStrengthCases = []test{
		{
			input:  "QJJQ2",
			output: 6,
		},
		{
			input:  "AAAAA",
			output: 7,
		},
		{
			input:  "AKT9J",
			output: 2,
		},
	}

	for _, c := range testGetHandStrengthCases {
		hand := CCHand{
			Cards: c.input,
			Bid:   0,
		}
		if result := hand.getStrength(true); result != c.output {
			t.Errorf("GetHandStrength(%v) == %d, expected %d", c.input, result, c.output)
		}
	}

}

func TestHandCompare(t *testing.T) {
	type test struct {
		input1 CCHand
		input2 CCHand
		output int
	}

	var testHandCompareCases = []test{
		{
			input1: CCHand{
				Cards: "AAAAA",
				Bid:   0,
			},
			input2: CCHand{
				Cards: "AAAAA",
				Bid:   0,
			},
			output: 0,
		},
		{
			input1: CCHand{
				Cards: "33332",
				Bid:   0,
			},
			input2: CCHand{
				Cards: "2AAAA",
				Bid:   0,
			},
			output: 1,
		},
		{
			input1: CCHand{
				Cards: "77888",
				Bid:   0,
			},
			input2: CCHand{
				Cards: "77788",
				Bid:   0,
			},
			output: 1,
		},
		{
			input1: CCHand{
				Cards: "KK677",
				Bid:   0,
			},
			input2: CCHand{
				Cards: "KTJJT",
				Bid:   0,
			},
			output: 1,
		},
	}

	for _, c := range testHandCompareCases {
		if result := c.input1.compare(&c.input2, false); result != c.output {
			t.Errorf("HandCompare(%v, %v) == %d, expected %d", c.input1, c.input2, result, c.output)
		}
	}
}
