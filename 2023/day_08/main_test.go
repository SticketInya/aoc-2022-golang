package main

import "testing"

func TestGetStepsToReachDestination(t *testing.T) {
	type test struct {
		input    []string
		expected int
	}

	cases := []test{
		{
			input: []string{
				"RL",
				"",
				"AAA = (BBB, CCC)",
				"BBB = (DDD, EEE)",
				"CCC = (ZZZ, GGG)",
				"DDD = (DDD, DDD)",
				"EEE = (EEE, EEE)",
				"GGG = (GGG, GGG)",
				"ZZZ = (ZZZ, ZZZ)",
			},
			expected: 2,
		},
		{
			input: []string{
				"LLR",
				"",
				"AAA = (BBB, BBB)",
				"BBB = (AAA, ZZZ)",
				"ZZZ = (ZZZ, ZZZ)",
			},
			expected: 6,
		},
	}

	for _, c := range cases {
		if result := getStepsToReachDestination(c.input); result != c.expected {
			t.Errorf("Expected %d, got %d", c.expected, result)
		}
	}
}

func TestGetGhostStepsToReachDestination(t *testing.T) {
	type test struct {
		input    []string
		expected int
	}

	cases := []test{
		{
			input: []string{
				"LR",
				"",
				"11A = (11B, XXX)",
				"11B = (XXX, 11Z)",
				"11Z = (11B, XXX)",
				"22A = (22B, XXX)",
				"22B = (22C, 22C)",
				"22C = (22Z, 22Z)",
				"22Z = (22B, 22B)",
				"XXX = (XXX, XXX)",
			},
			expected: 6,
		},
	}

	for _, c := range cases {
		if result := getGhostStepsToReachDestination(c.input); result != c.expected {
			t.Errorf("Expected %d, got %d", c.expected, result)
		}
	}
}
