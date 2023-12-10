package main

import "testing"

func TestGetFurthestPosition(t *testing.T) {
	type test struct {
		input []string
		want  int
	}

	cases := []test{
		{
			input: []string{
				"..F7.",
				".FJ|.",
				"SJ.L7",
				"|F--J",
				"LJ...",
			},
			want: 8,
		},
		{
			input: []string{
				".....",
				".S-7.",
				".|.|.",
				".L-J.",
				".....",
			},
			want: 4,
		},
	}

	for _, c := range cases {
		if result := getFurthestPosition(c.input); result != c.want {
			t.Errorf("getFurthestPosition(%v) == %d, want %d", c.input, result, c.want)
		}
	}
}
