package main

import "testing"

func TestLongestIncreasingPath(t *testing.T) {
	cases := []struct {
		in   [][]int
		want int
	}{
		{[][]int{{9, 9, 4}, {6, 6, 8}, {2, 1, 1}}, 4}, // 1-2-6-9
		{[][]int{{3, 4, 5}, {3, 2, 6}, {2, 2, 1}}, 4}, // 3-4-5-6
		{[][]int{{1}}, 1},
		{[][]int{{7, 7}, {7, 7}}, 1}, // strict increase: plateaus don't chain
		{nil, 0},
		{[][]int{{1, 2, 3, 4, 5}}, 5},
	}
	for _, c := range cases {
		if got := LongestIncreasingPath(c.in); got != c.want {
			t.Errorf("LongestIncreasingPath = %d, want %d", got, c.want)
		}
	}
}
