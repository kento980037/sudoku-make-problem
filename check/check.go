package check

import "work/sudoku-make-problem/table"

// Check for compliance with Sudoku rules.
func Check(b table.Board) bool {
	for i := 0; i < 9; i++ {
		var c [10]int
		for j := 0; j < 9; j++ {
			c[b[i][j]]++
		}
		if duplicated(c) {
			return false
		}
	}
	for j := 0; j < 9; j++ {
		var c [10]int
		for i := 0; i < 9; i++ {
			c[b[i][j]]++
		}
		if duplicated(c) {
			return false
		}
	}
	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {
			var c [10]int
			for row := i; row < i+3; row++ {
				for col := j; col < j+3; col++ {
					c[b[row][col]]++
				}
			}
			if duplicated(c) {
				return false
			}
		}
	}
	return true
}

// Returns false if the element of array c with indices 1~9 is greater than or equal to 2
func duplicated(c [10]int) bool {
	for k, v := range c {
		if k == 0 {
			continue
		}
		if v >= 2 {
			return true
		}
	}
	return false
}

func IsSolved(b table.Board) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if b[i][j] == 0 {
				return false
			}
		}
	}
	return true
}
