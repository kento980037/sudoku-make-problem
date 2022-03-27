package solve

import (
	"work/sudoku-make-problem/check"
	"work/sudoku-make-problem/table"
)

func Solve(b *table.Board) bool {
	if check.IsSolved(*b) {
		return true
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if b[i][j] == 0 {
				for c := 9; c > 0; c-- {
					b[i][j] = c
					if check.Check(*b) {
						if Solve(b) {
							return true
						}
					}
					b[i][j] = 0
				}
				return false
			}
		}
	}
	return false
}
