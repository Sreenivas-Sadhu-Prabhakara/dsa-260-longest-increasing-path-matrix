package main

import "fmt"

// LongestIncreasingPath returns the length of the longest strictly
// increasing path through 4-adjacent cells. DFS with memoization: strict
// increase makes the move graph a DAG, so each cell's answer is computed once.
func LongestIncreasingPath(matrix [][]int) int {
	rows := len(matrix)
	if rows == 0 || len(matrix[0]) == 0 {
		return 0
	}
	cols := len(matrix[0])
	memo := make([][]int, rows)
	for r := range memo {
		memo[r] = make([]int, cols)
	}
	var longest func(r, c int) int
	longest = func(r, c int) int {
		if memo[r][c] != 0 {
			return memo[r][c]
		}
		best := 1
		for _, d := range [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
			nr, nc := r+d[0], c+d[1]
			if nr < 0 || nr >= rows || nc < 0 || nc >= cols || matrix[nr][nc] <= matrix[r][c] {
				continue
			}
			if l := 1 + longest(nr, nc); l > best {
				best = l
			}
		}
		memo[r][c] = best
		return best
	}
	answer := 0
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if l := longest(r, c); l > answer {
				answer = l
			}
		}
	}
	return answer
}

func main() {
	fmt.Println(LongestIncreasingPath([][]int{{9, 9, 4}, {6, 6, 8}, {2, 1, 1}}))
}
