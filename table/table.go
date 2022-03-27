package table

import (
	"bytes"
	"strconv"
)

type Board [9][9]int

func Organize(b Board) string {
	var buf bytes.Buffer
	buf.WriteString("\n")
	for i := 0; i < 9; i++ {
		if i%3 == 0 {
			buf.WriteString("+---+---+---+\n")
		}
		for j := 0; j < 9; j++ {
			if j%3 == 0 {
				buf.WriteString("|")
			}
			if b[i][j] == 0 {
				buf.WriteString(" ")
			} else {
				buf.WriteString(strconv.Itoa(b[i][j]))
			}

		}
		buf.WriteString("|\n")
	}
	buf.WriteString("+---+---+---+\n")
	return buf.String()
}
