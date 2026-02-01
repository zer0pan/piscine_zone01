package main

import (
	"fmt"
	"os"
)

const N = 9

// solves sudoku (call by reference)
func solve(grid *[N][N]rune) bool {
	//to *grid einai h timh tou pointer grid , diladi o pinakas grid[N][N]
	row, col := -1, -1
	empty := false

	// finds the next empty cell (empty cell is '.')
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if grid[i][j] == '.' {
				row, col = i, j
				empty = true
				break
			}
		}
		if empty {
			break
		}
	}

	// if there are any empty cells then done!
	if !empty {
		return true
	}

	for num := '1'; num <= '9'; num++ {
		if ValidSquare((*grid), row, col, num) {
			grid[row][col] = num

			if solve(grid) {
				return true
			}
			grid[row][col] = '.'
		}
	}
	return false
}

func ValidSquare(grid [N][N]rune, row, col int, num rune) bool {
	// checking by row
	for x := 0; x < N; x++ {
		if grid[row][x] == num {
			return false
		}
	}
	// checking by column
	for x := 0; x < N; x++ {
		if grid[x][col] == num {
			return false
		}
	}
	// checking by 3x3 squares
	srow := (row / 3) * 3
	scol := (col / 3) * 3
	for i := srow; i < srow+3; i++ {
		for j := scol; j < scol+3; j++ {
			if grid[i][j] == num {
				return false
			}
		}
	}
	return true
}

// it prints the sudoku grid
func printGrid(grid [N][N]rune) {
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if j == N-1 {
				fmt.Printf("%c", grid[i][j])
				break
			} else {
				fmt.Printf("%c ", grid[i][j])
			}
		}
		fmt.Println()
	}
}

// checking if the grid is valid for processing
func ValidSudoku(grid [N][N]rune) bool {
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if grid[i][j] != '.' {
				num := grid[i][j]
				grid[i][j] = '.'
				if !ValidSquare(grid, i, j, num) {
					return false // Invalid placement found
				}
				grid[i][j] = num
			}
		}
	}
	// If all cells are valid, return true
	return true
}

func main() {
	args := os.Args
	// input validity check
	if len(args) != 10 {
		fmt.Println("Error")
		return
	} else {
		for i := 1; i < len(args); i++ {
			if len(args[i]) != 9 {
				fmt.Println("Error")
				return
			}
			for j := 0; j < len(args[i]); j++ {
				if !((args[i][j] <= '9' && args[i][j] >= '1') || (args[i][j] == '.')) {
					fmt.Println("Error")
					return
				}
			}
		}
	}
	// constructing sudoku
	grid := [N][N]rune{}
	for i := 1; i < 10; i++ {
		for j := 0; j < 9; j++ {
			grid[i-1][j] = rune(args[i][j])
		}
	}
	// checking the validity of the grid
	if !ValidSudoku(grid) {
		fmt.Println("Error")
		return
	}

	if solve(&grid) {
		printGrid(grid)
	} else {
		fmt.Println("Error")
		return
	}
}
