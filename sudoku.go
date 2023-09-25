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
		panic(err)
	}

	fmt.Println("Input:")
	fmt.Print(in)

	in.SetCallback(progress)
	if !in.Solve() {
		fmt.Println("unsolvable!")
	}
	fmt.Println(in)
}

func progress(s *solver.Sudoku, step int, r int, c int, val int, strat string) {
	fmt.Printf("Step %d, R:%d, C:%d, Val:%d, Strategy:%s\n", step, r, c, val, strat)
	fmt.Print(s)
}
