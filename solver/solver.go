package solver

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"golang.org/x/exp/slices"
)

type Sudoku struct {
	matrix     [9][9]int // the (working) table
	complexity int       // how difficult is this sudoku?
	callback   func(s *Sudoku, step int, r int, c int, val int, strat string)
	//digitsFound []int     // how many of thee are in the matrix already
}

type dimension int

const (
	Row   dimension = 0
	Col   dimension = 1
	Block dimension = 2
)

// @return: is it solved?
func (s *Sudoku) Solve() bool {
	step := 0
	for !s.isDone() {
		step++

		if success, r, c, val, strat := s.findLevel1(); success {
			s.setComplexity(1)
			if s.callback != nil {
				s.callback(s, step, r, c, val, strat)
			}
			continue
		}

		if success, r, c, val, strat := s.findLevel2(); success {
			s.setComplexity(2)
			if s.callback != nil {
				s.callback(s, step, r, c, val, strat)
			}
			continue
		}
		return false
	}
	return true
}

func (s *Sudoku) String() string {
	str := ""
	for i := 0; i < 9; i++ {
		if i%3 == 0 {
			str += "+---+---+---+\n"
		}
		for j := 0; j < 9; j++ {
			if j%3 == 0 {
				str += "|"
			}
			if s.matrix[i][j] == 0 {
				str += " "
			} else {
				str += fmt.Sprint(s.matrix[i][j])
			}
		}
		str += "|\n"
	}
	str += "+---+---+---+\n"
	return str
}

/*
Load a sudoku table from textual input
Input can be:
  - pure digits where unknowns are 0s
  - comma separated digits where unknowns are X or 0 or empty
  - a pretty-printed table
*/
func (s *Sudoku) LoadString(input string) (err error) {
	i := 0
	lineno := 0
	lines := strings.Split(input, "\n")
	for i < 9 {
		line := lines[lineno]
		line = strings.Replace(line, ",", "", -1)
		line = strings.Replace(line, "-", "", -1)
		line = strings.Replace(line, "+", "", -1)
		line = strings.Replace(line, "|", "", -1)
		line = strings.Replace(strings.ToLower(line), "x", "0", -1)
		line = strings.Replace(line, " ", "0", -1)

		if len(line) == 0 || line[0] == '#' || len(line) < 9 {
			lineno++
			continue
		}

		for j := 0; j < 9; j++ {
			v := int(line[j]) - int('0')
			if v < 0 || v > 9 {
				return fmt.Errorf("invalid input at %d %d", i, j)
			}
			s.matrix[i][j] = v
		}

		lineno++
		i++
	}
	return nil
}

func (s *Sudoku) LoadFile(readFile *os.File) (err error) {
	scanner := bufio.NewScanner(readFile)
	var in string
	max := 0
	for scanner.Scan() && max < 100 {
		in += scanner.Text() + "\n"
		max++
	}
	readFile.Close()

	s.LoadString(in)

	return nil
}

func (s *Sudoku) missingTotal() int {
	n := 0
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if s.matrix[i][j] == 0 {
				n++
			}
		}
	}
	return n
}

func (s *Sudoku) isDone() bool {
	return s.missingTotal() == 0
}

// 8 in a row, col or block => fill the missing one
func (s *Sudoku) findLevel1() (success bool, r int, c int, val int, strat string) {
	fillMissingItem := func(dim dimension, where int) (success bool, r int, c int, val int) {
		n := 0
		loc := 0
		possiblevals := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		for i := 0; i < 9; i++ {
			r, c := translateIndextoRC(dim, where, i)
			if s.matrix[r][c] == 0 {
				n++
				loc = i
			} else {
				possiblevals = removeFromSlice(possiblevals, s.matrix[r][c])
			}
		}
		if n == 1 {
			r, c := translateIndextoRC(dim, where, loc)
			s.matrix[r][c] = possiblevals[0]
			return true, r, c, possiblevals[0]
		}
		return false, 9, 9, 9
	}

	for i := 0; i < 9; i++ {
		if success, r, c, val := fillMissingItem(Row, i); success {
			return true, r, c, val, "1"
		}
		if success, r, c, val := fillMissingItem(Col, i); success {
			return true, r, c, val, "1"
		}
		if success, r, c, val := fillMissingItem(Block, i); success {
			return true, r, c, val, "1"
		}
	}
	return false, 9, 9, 9, ""
}

func (s *Sudoku) findLevel2() (success bool, r int, c int, val int, strat string) {
	// try to put v in all the free positions
	for v := 1; v <= 9; v++ {
		// OPT: if v already has 9 instances, skip it
		for r := 0; r < 9; r++ {
			for c := 0; c < 9; c++ {
				if s.matrix[r][c] == 0 &&
					!s.hasAlready(Row, r, v) &&
					!s.hasAlready(Col, c, v) &&
					!s.hasAlready(Block, translateRCToBlock(r, c), v) {

					r1 := 3*(r/3) + (r+1)%3 // first other row
					r2 := 3*(r/3) + (r+2)%3 // second other row
					c1 := 3*(c/3) + (c+1)%3 // first other column
					c2 := 3*(c/3) + (c+2)%3 // second other row

					r1has := s.hasAlready(Row, r1, v)
					r2has := s.hasAlready(Row, r2, v)
					c1has := s.hasAlready(Col, c1, v)
					c2has := s.hasAlready(Col, c2, v)

					// check if *other* rows/cols in the same row/col group have this v
					// if yes for all => fill it in!
					if r1has && r2has && c1has && c2has {
						s.matrix[r][c] = v
						return true, r, c, v, "2a"
					}

					// check if *other* rows/cols in the same row/col group have this v
					// also, if *other* rows/cols cannot have it because in this block they are filled
					// if yes for all => fill it in!
					if (r1has || (s.matrix[r1][c] != 0 && s.matrix[r2][c] != 0)) &&
						(r2has || (s.matrix[r1][c] != 0 && s.matrix[r2][c] != 0)) &&
						(c1has || (s.matrix[r][c1] != 0 && s.matrix[r][c2] != 0)) &&
						(c2has || (s.matrix[r][c1] != 0 && s.matrix[r][c2] != 0)) {
						s.matrix[r][c] = v
						return true, r, c, v, "2b"
					}

					// check if *other* rows/cols in the same row/col group have this v
					// also, if *other* rows/cols cannot have it because in this block they are filled
					// also, ignore already filled rows/cols
					// if yes for all => fill it in!
					if (r1has || (s.matrix[r1][c] != 0 && s.matrix[r1][c1] != 0 && s.matrix[r1][c2] != 0)) &&
						(r2has || (s.matrix[r2][c] != 0 && s.matrix[r2][c1] != 0 && s.matrix[r2][c2] != 0)) &&
						s.matrix[r][c1] != 0 && s.matrix[r][c2] != 0 {
						s.matrix[r][c] = v
						return true, r, c, v, "2c1"
					}
					if (c1has || (s.matrix[r][c1] != 0 && s.matrix[r1][c1] != 0 && s.matrix[r2][c1] != 0)) &&
						(c2has || (s.matrix[r][c2] != 0 && s.matrix[r1][c2] != 0 && s.matrix[r2][c2] != 0)) &&
						s.matrix[r1][c] != 0 && s.matrix[r2][c] != 0 {
						s.matrix[r][c] = v
						return true, r, c, v, "2c2"
					}
				}
			}
		}
	}
	return false, 9, 9, 9, ""
}

func (s *Sudoku) hasAlready(dim dimension, where int, val int) bool {
	for i := 0; i < 9; i++ {
		r, c := translateIndextoRC(dim, where, i)
		if s.matrix[r][c] == val {
			return true
		}
	}
	return false
}

func (s *Sudoku) setComplexity(complexity int) {
	if s.complexity < complexity {
		s.complexity = complexity
	}
}

func (s *Sudoku) SetCallback(callback func(s *Sudoku, step int, r int, c int, val int, strategy string)) {
	s.callback = callback
}

func (s *Sudoku) Complexity() int {
	return s.complexity
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

func translateRCToBlock(r int, c int) int {
	return (r/3)*3 + c/3
}
