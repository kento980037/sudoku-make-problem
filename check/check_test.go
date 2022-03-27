package check

import (
	"testing"
	"work/sudoku-make-problem/table"
)

func TestCheck(t *testing.T) {
	type args struct {
		b table.Board
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "normal_1",
			args: args{
				b: table.Board{
					{0, 0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0, 0},
				},
			},
			want: true,
		},
		{
			name: "normal_2",
			args: args{
				b: table.Board{
					{1, 4, 0, 6, 0, 0, 0, 8, 0},
					{7, 0, 2, 1, 0, 0, 5, 0, 0},
					{0, 0, 0, 0, 0, 0, 7, 0, 6},
					{0, 0, 0, 0, 0, 4, 0, 0, 3},
					{4, 5, 6, 0, 0, 0, 0, 0, 0},
					{9, 0, 3, 0, 0, 6, 0, 7, 1},
					{0, 1, 0, 0, 7, 0, 0, 6, 0},
					{0, 0, 0, 0, 0, 0, 0, 3, 5},
					{0, 0, 5, 0, 0, 1, 9, 0, 0},
				},
			},
			want: true,
		},
		{
			name: "normal_3",
			args: args{
				b: table.Board{
					{1, 4, 9, 6, 5, 7, 3, 8, 2},
					{7, 6, 2, 1, 3, 8, 5, 9, 4},
					{5, 3, 8, 4, 2, 9, 7, 1, 6},
					{8, 7, 1, 2, 9, 4, 6, 5, 3},
					{4, 5, 6, 7, 1, 3, 8, 2, 9},
					{9, 2, 3, 5, 8, 6, 4, 7, 1},
					{3, 1, 4, 9, 7, 5, 2, 6, 8},
					{6, 9, 7, 8, 4, 2, 1, 3, 5},
					{2, 8, 5, 3, 6, 1, 9, 4, 7},
				},
			},
			want: true,
		},
		{
			name: "abnormal_1: Error in the rows",
			args: args{
				b: table.Board{
					{1, 0, 0, 1, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0, 0},
				},
			},
			want: false,
		},
		{
			name: "abnormal_2: Error in the cols",
			args: args{
				b: table.Board{
					{1, 0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0, 0},
					{1, 0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0, 0},
				},
			},
			want: false,
		},
		{
			name: "abnormal_3: Error in the 3*3 box",
			args: args{
				b: table.Board{
					{1, 0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0, 0},
					{0, 1, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0, 0},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Check(tt.args.b); got != tt.want {
				t.Errorf("Check() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsSolved(t *testing.T) {
	type args struct {
		b table.Board
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "normal",
			args: args{
				b: table.Board{
					{1, 4, 9, 6, 5, 7, 3, 8, 2},
					{7, 6, 2, 1, 3, 8, 5, 9, 4},
					{5, 3, 8, 4, 2, 9, 7, 1, 6},
					{8, 7, 1, 2, 9, 4, 6, 5, 3},
					{4, 5, 6, 7, 1, 3, 8, 2, 9},
					{9, 2, 3, 5, 8, 6, 4, 7, 1},
					{3, 1, 4, 9, 7, 5, 2, 6, 8},
					{6, 9, 7, 8, 4, 2, 1, 3, 5},
					{2, 8, 5, 3, 6, 1, 9, 4, 7},
				},
			},
			want: true,
		},
		{
			name: "abnormal",
			args: args{
				b: table.Board{
					{0, 0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0, 0},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsSolved(tt.args.b); got != tt.want {
				t.Errorf("IsSolved() = %v, want %v", got, tt.want)
			}
		})
	}
}
