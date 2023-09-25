package main

import (
	"fmt"
	"sudoku/solver"
)

func main() {
	in := new(solver.Sudoku)
	err := in.Load(`
# complexity 2?
+---+---+---+
|1  |   |   |
|   |   |456|
|   |   |23X|
+---+---+---+
|   |   |   |
|   |   |   |
|   |   |   |
+---+---+---+
|   |   |   |
|   |   |   |
|   |   |   |
+---+---+---+
`)

	if err != nil {
		panic(err)
	}

	fmt.Print(in)
	in.SetVerbose(true)
	if !in.Solve() {
		fmt.Print("unsolveable\n", in.String(), "\n")
	}
}
