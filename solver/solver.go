package solver

import (
	"fmt"
	"strings"

	"golang.org/x/exp/slices"
)

type Sudoku struct {
	matrix     [9][9]int // the (working) table
	complexity int       // how difficult is this sudoku?
	//digitsFound []int     // how many of thee are in the matrix already
}

type dimension int

const (
	Row   dimension = 0
	Col             = 1
	Block           = 2
)

func (sudoku *Sudoku) Solve() (err error) {
	for sudoku.missingTotal() > 0 {
		if sudoku.findTrivial() {
			fmt.Print(sudoku)
			continue
		}
		return fmt.Errorf("Unsolveable, complexity=%d", sudoku.complexity)
	}
	return nil
}

func (sudoku *Sudoku) String() string {
	s := ""
	for i := 0; i < 9; i++ {
		if i%3 == 0 {
			s += "+---+---+---+\n"
		}
		for j := 0; j < 9; j++ {
			if j%3 == 0 {
				s += "|"
			}
			s += fmt.Sprint(sudoku.matrix[i][j])
		}
		s += "|\n"
	}
	s += "+---+---+---+\n"
	return s
}

/*
Load a sudoku table from textual input
Input can be:
  - pure digits where unknowns are 0s
  - comma separated digits where unknowns are X or 0 or empty
  - //a pretty-printed table
*/
func (sudoku *Sudoku) Load(input []string) (err error) {
	i := 0
	for i < 9 {
		line := input[i]
		line = strings.Replace(line, ",", "", 9)
		line = strings.Replace(line, "X", "0", 9)
		for j := 0; j < 9; j++ {
			v := int(line[j]) - int('0')
			if v < 0 || v > 9 {
				return fmt.Errorf("Invalid input at %d %d", i, j)
			}
			sudoku.matrix[i][j] = v
		}
		i++
	}
	return nil
}

func (sudoku *Sudoku) missingTotal() int {
	n := 0
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if sudoku.matrix[i][j] == 0 {
				n++
			}
		}
	}
	return n
}

func (sudoku *Sudoku) isDone() bool {
	return sudoku.missingTotal() == 0
}

// 8 in a row, col or block => fill the missing one
func (sudoku *Sudoku) findTrivial() bool {
	fillMissingItem := func(dim dimension, where int) bool {
		n := 0
		loc := 0
		possiblevals := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		for i := 0; i < 9; i++ {
			r, c := translateIndextoRC(dim, where, i)
			if sudoku.matrix[r][c] == 0 {
				n++
				loc = i
			} else {
				possiblevals = removeFromSlice(possiblevals, sudoku.matrix[r][c])
			}
		}
		if n == 1 {
			r, c := translateIndextoRC(dim, where, loc)
			sudoku.matrix[r][c] = possiblevals[0]
			fmt.Println("Filled", dim, "for index", where, "R:", r, "C:", c, "with ", possiblevals[0])
			return true
		}
		return false
	}

	for i := 0; i < 9; i++ {
		if fillMissingItem(Row, i) ||
			fillMissingItem(Col, i) ||
			fillMissingItem(Block, i) {
			sudoku.setComplexity(1)
			return true
		}
	}
	return false
}

func (sudoku *Sudoku) setComplexity(complexity int) {
	if sudoku.complexity < complexity {
		sudoku.complexity = complexity
	}
}

// ////////////////////// utils
func removeFromSlice(slice []int, remove int) []int {
	for i := 0; i < len(slice); i++ {
		if slice[i] == remove {
			return slices.Delete(slice, i, i+1)
		}
	}
	return slice
}

// Given a dimension (row, col, block), translate the index into the row/col coordinates of that R/C/B
func translateIndextoRC(dim dimension, where int, index int) (r, c int) {
	switch dim {
	case Row:
		return where, index
	case Col:
		return index, where
	case Block:
		return index/3 + 3*(where/3), index%3 + 3*(where%3)
	}
	return 9, 9 // this will blow up
}
