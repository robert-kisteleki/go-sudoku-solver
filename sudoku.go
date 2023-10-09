package main

import (
	"fmt"
	"os"
	"sudoku/solver"
)

func main() {
	in := new(solver.Sudoku)

	err := in.LoadFile(os.Stdin)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println("Input:")
	fmt.Print(in)

	if err := in.Solve(0); err != nil {
		fmt.Println(err)
	}

	for i, s := range *in.Steps() {
		fmt.Printf("Step %d, R:%d, C:%d, Value:%d, Strategy:%s\n", i+1, s.R+1, s.C+1, s.Value, s.Strategy)
		fmt.Print(s.BoardAfter.String())
	}
}
