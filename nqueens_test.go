// Author: Spiros Nikoloudakis
// Email: sp.nikoloudakis@gmail.com
// GitHub: @grsprs
// Year: 2026

package main

import (
	"fmt"
	"testing"
)

// TestNQueensSolutionCount validates the correct number of solutions for known N values.
// Requirement: Verify backtracking finds all valid distinct arrangements.
func TestNQueensSolutionCount(t *testing.T) {
	tests := []struct {
		n        int
		expected int
	}{
		{1, 1},
		{2, 0},
		{3, 0},
		{4, 2},
		{5, 10},
		{6, 4},
		{7, 40},
		{8, 92},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("N=%d", tt.n), func(t *testing.T) {
			board := make([]int, tt.n)
			cols := make(map[int]bool)
			diag1 := make(map[int]bool)
			diag2 := make(map[int]bool)
			count := 0

			var backtrack func(int)
			backtrack = func(row int) {
				if row == tt.n {
					count++
					return
				}
				for col := 0; col < tt.n; col++ {
					if cols[col] || diag1[row-col] || diag2[row+col] {
						continue
					}
					board[row] = col
					cols[col] = true
					diag1[row-col] = true
					diag2[row+col] = true
					backtrack(row + 1)
					cols[col] = false
					diag1[row-col] = false
					diag2[row+col] = false
				}
			}

			backtrack(0)

			if count != tt.expected {
				t.Errorf("N=%d: got %d solutions, want %d", tt.n, count, tt.expected)
			}
		})
	}
}

// TestNoQueensAttack validates that no two queens attack each other in solutions.
// Requirement: Queens must not share row, column, or diagonal.
func TestNoQueensAttack(t *testing.T) {
	n := 4
	board := make([]int, n)
	cols := make(map[int]bool)
	diag1 := make(map[int]bool)
	diag2 := make(map[int]bool)
	solutions := [][]int{}

	var backtrack func(int)
	backtrack = func(row int) {
		if row == n {
			sol := make([]int, n)
			copy(sol, board)
			solutions = append(solutions, sol)
			return
		}
		for col := 0; col < n; col++ {
			if cols[col] || diag1[row-col] || diag2[row+col] {
				continue
			}
			board[row] = col
			cols[col] = true
			diag1[row-col] = true
			diag2[row+col] = true
			backtrack(row + 1)
			cols[col] = false
			diag1[row-col] = false
			diag2[row+col] = false
		}
	}

	backtrack(0)

	for _, sol := range solutions {
		for i := 0; i < n; i++ {
			for j := i + 1; j < n; j++ {
				if sol[i] == sol[j] {
					t.Errorf("Queens at rows %d and %d share column %d", i, j, sol[i])
				}
				if abs(i-j) == abs(sol[i]-sol[j]) {
					t.Errorf("Queens at (%d,%d) and (%d,%d) share diagonal", i, sol[i], j, sol[j])
				}
			}
		}
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
