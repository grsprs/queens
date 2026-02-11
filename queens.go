// Author: Spiros Nikoloudakis
// Email: sp.nikoloudakis@gmail.com
// GitHub: @grsprs
// Year: 2026

package main

import "fmt"

// solve finds all N-Queens solutions using backtracking.
// Places one queen per row, validates columns and diagonals.
func solve(n int) {
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
	printSolutions(solutions, n)
}

// printSolutions outputs all solutions in required format.
func printSolutions(solutions [][]int, n int) {
	for i, sol := range solutions {
		fmt.Printf("Solution %d:\n", i+1)
		for row := 0; row < n; row++ {
			for col := 0; col < n; col++ {
				if sol[row] == col {
					fmt.Print("Q")
				} else {
					fmt.Print(".")
				}
			}
			fmt.Println()
		}
		fmt.Println()
	}
	fmt.Printf("Total solutions: %d\n", len(solutions))
}

func main() {
	var n int
	fmt.Scanln(&n)
	solve(n)
}
