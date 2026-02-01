# Sudoku Solver (Go)

This project is a **Sudoku solver written in Go**, developed during the  
**Zone01 Piscine – Selection Bootcamp**.

## 📌 Description

The program:
- Reads a 9×9 Sudoku grid from command-line arguments
- Validates the initial grid according to Sudoku rules
- Solves the Sudoku using **backtracking**
- Prints the solved grid, or `Error` if:
  - the input is invalid
  - the Sudoku has no solution

Empty cells are represented by `.`.

## ⚙️ Usage

### Run with Go

```sh
go run . ".96.4...1" "1...6...4" "5.481.39." "..795..43" ".3..8...." "4.5.23.18" ".1.63..59" ".59.7.83." "..35....."
