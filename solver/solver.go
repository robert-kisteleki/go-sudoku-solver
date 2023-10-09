package solver

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

type Sudoku struct {
	board      Board   // the (working) board
	complexity int     // how difficult is this sudoku?
	occurences [10]int // how many of these are in the board already (never mind 0)
	steps      []Step
}

type dimension int

const (
	row   dimension = 0
	col   dimension = 1
	block dimension = 2
)

type Step struct {
	R, C       int
	Value      int
	Strategy   string
	BoardAfter Board
}

type Board [9][9]int

// Solve the sudoku
// maxsteps limits how many steps should be done (0==all)
// @return: is it solved? err==nil means yes
func (s *Sudoku) Solve(maxteps int) (err error) {
	solveLevelFunc := func(level int) func() (success bool) {
		switch level {
		case 1:
			return s.findLevel1
		case 2:
			return s.findLevel2
		case 3:
			return s.findLevel3
		}
		return nil
	}

	step := 0
	for !s.isDone() && (maxteps <= 0 || step < maxteps) {
		step++

		found := false
		for level := 1; level <= 3; level++ {
			if success := solveLevelFunc(level)(); success {
				s.setComplexity(level)
				if err := s.isSane(); err != nil {
					return fmt.Errorf("oops, after level %d step the solution is not sane: %v", level, err)
				}
				found = true
				break
			}
		}

		if !found {
			return fmt.Errorf("not solvable")
		}
	}

	return nil
}

func (board *Board) String() string {
	str := ""
	for i := 0; i < 9; i++ {
		if i%3 == 0 {
			str += "+---+---+---+\n"
		}
		for j := 0; j < 9; j++ {
			if j%3 == 0 {
				str += "|"
			}
			if board[i][j] == 0 {
				str += " "
			} else {
				str += fmt.Sprint(board[i][j])
			}
		}
		str += "|\n"
	}
	str += "+---+---+---+\n"
	return str
}

func (s *Sudoku) String() string {
	return s.board.String()
}

/*
Load a sudoku table from textual input
Input can be:
  - pure digits where unknowns are 0s
  - comma separated digits where unknowns are X or 0 or empty
  - a pretty-printed table
*/
func (s *Sudoku) LoadString(input string) (err error) {
	var board [9][9]int
	r := 0
	lineno := 0
	lines := strings.Split(input, "\n")
	for r < 9 {
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

		for c := 0; c < 9; c++ {
			v := int(line[c]) - int('0')
			if v < 0 || v > 9 {
				return fmt.Errorf("invalid input at %d %d", r, c)
			}
			board[r][c] = v
		}

		lineno++
		r++
	}

	if err := s.LoadArray(board); err != nil {
		return err
	}
	return s.isSane()
}

// LoadArray initialises from another [][]int matrix
// input should be 9x9
func (s *Sudoku) LoadArray(input [9][9]int) (err error) {
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			s.setValue(r, c, input[r][c], "")
		}
	}
	return s.isSane()
}

// LoadFile loads from a file (!)
func (s *Sudoku) LoadFile(readFile *os.File) (err error) {
	scanner := bufio.NewScanner(readFile)
	var in string
	max := 0
	for scanner.Scan() && max < 100 {
		in += scanner.Text() + "\n"
		max++
	}
	readFile.Close()

	return s.LoadString(in)
}

// set a cell value, log and update statistics
func (s *Sudoku) setValue(r int, c int, v int, strat string) {
	s.board[r][c] = v
	s.occurences[v]++
	if strat != "" {
		s.steps = append(s.steps, Step{r, c, v, strat, s.board})
	}
}

// chek if current values are possible or not
func (s *Sudoku) isSane() (err error) {
	// check if values are [0..9]
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if s.board[r][c] < 0 || s.board[r][c] > 9 {
				return fmt.Errorf("invalid input at R:%d C:%d V:%d", r, c, s.board[r][c])
			}
		}
	}

	// check how many [1..9] we have per R/C/B
	for r := 0; r < 9; r++ {
		for v := 1; v <= 9; v++ {
			n := s.count(row, r, v)
			if n > 1 {
				return fmt.Errorf("row %d contains %d appearances of %d", r+1, n, v)
			}
		}
	}
	for c := 0; c < 9; c++ {
		for v := 1; v <= 9; v++ {
			n := s.count(col, c, v)
			if n > 1 {
				return fmt.Errorf("col %d contains %d appearances of %d", c+1, n, v)
			}
		}
	}
	for b := 0; b < 9; b++ {
		for v := 1; v <= 9; v++ {
			n := s.count(block, b, v)
			if n > 1 {
				return fmt.Errorf("block %d contains %d appearances of %d", b+1, n, v)
			}
		}
	}

	return nil
}

func (s *Sudoku) isDone() bool {
	n := 0
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if !s.isFilled(i, j) {
				n++
			}
		}
	}
	return n == 0 && s.isSane() == nil
}

// check for complexity 1
// 8 in a row, col or block => fill the missing one
func (s *Sudoku) findLevel1() (success bool) {
	fillMissingItem := func(dim dimension, where int) (success bool) {
		n := 0
		loc := 0
		possiblevals := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		for i := 0; i < 9; i++ {
			r, c := translateIndextoRC(dim, where, i)
			if s.board[r][c] == 0 {
				n++
				loc = i
			} else {
				possiblevals = removeFromSlice(possiblevals, s.board[r][c])
			}
		}
		if n == 1 {
			r, c := translateIndextoRC(dim, where, loc)
			s.setValue(r, c, possiblevals[0], "1")
			return true
		}
		return false
	}

	for i := 0; i < 9; i++ {
		for check := range []dimension{row, col, block} {
			if success := fillMissingItem(dimension(check), i); success {
				return true
			}
		}
	}
	return false
}

// check for complexity 2
// check for cases where there can be only one value in a cell
func (s *Sudoku) findLevel2() (success bool) {
	// try to put v in all the free positions
	for _, v := range orderedCandidates(s.occurences[:]) {
		for r := 0; r < 9; r++ {
			for c := 0; c < 9; c++ {
				if !s.isFilled(r, c) &&
					!s.hasAlready(row, r, v) &&
					!s.hasAlready(col, c, v) &&
					!s.hasAlready(block, translateRCToBlock(r, c), v) {

					r1 := 3*(r/3) + (r+1)%3 // first other row
					r2 := 3*(r/3) + (r+2)%3 // second other row
					c1 := 3*(c/3) + (c+1)%3 // first other column
					c2 := 3*(c/3) + (c+2)%3 // second other column

					r1has := s.hasAlready(row, r1, v)
					r2has := s.hasAlready(row, r2, v)
					c1has := s.hasAlready(col, c1, v)
					c2has := s.hasAlready(col, c2, v)

					// check if *all other* rows/cols in the same row/col group have this v
					// if yes for all => fill it in!
					if r1has && r2has && c1has && c2has {
						s.setValue(r, c, v, "2a")
						return true
					}

					// check if *all other* rows/cols in the same row/col group have this v
					// also, if *other* rows/cols cannot have it because in this block they are filled
					// if yes for all => fill it in!
					if (r1has || (s.isFilled(r1, c) && s.isFilled(r1, c1) && s.isFilled(r1, c2))) &&
						(r2has || (s.isFilled(r2, c) && s.isFilled(r2, c1) && s.isFilled(r2, c2))) &&
						(c1has || (s.isFilled(r, c1) && s.isFilled(r1, c1) && s.isFilled(r2, c1))) &&
						(c2has || (s.isFilled(r, c2) && s.isFilled(r1, c2) && s.isFilled(r2, c2))) {
						s.setValue(r, c, v, "2b")
						return true
					}

					// check if *other* rows/cols in the same row/col group have this v
					// also, if *other* rows/cols cannot have it because in this block they are filled
					// also, ignore already filled rows/cols
					// if yes for all => fill it in!
					if (r1has || (s.isFilled(r1, c) && s.isFilled(r1, c1) && s.isFilled(r1, c2))) &&
						(r2has || (s.isFilled(r2, c) && s.isFilled(r2, c1) && s.isFilled(r2, c2))) &&
						s.isFilled(r, c1) && s.isFilled(r, c2) {
						s.setValue(r, c, v, "2c1")
						return true
					}
					if (c1has || (s.isFilled(r, c1) && s.isFilled(r1, c1) && s.isFilled(r2, c1))) &&
						(c2has || (s.isFilled(r, c2) && s.isFilled(r1, c2) && s.isFilled(r2, c2))) &&
						s.isFilled(r1, c) && s.isFilled(r2, c) {
						s.setValue(r, c, v, "2c2")
						return true
					}

					// are there any other cells in this row where v could be?
					otherCs := false
					for i := 0; i < 9 && !otherCs; i++ {
						if i == c || // same cell, skip
							s.isFilled(r, i) || // cell is filled already
							s.hasAlready(block, translateRCToBlock(r, i), v) || // if the block already has this v
							s.hasAlready(col, i, v) { // some other row in this col has this v already
							continue
						}
						otherCs = true
					}
					if !otherCs {
						// this v cannot be in any other columns
						s.setValue(r, c, v, "2dc")
						return true
					}

					// are there any other cells in this col where v could be?
					otherRs := false
					for i := 0; i < 9 && !otherRs; i++ {
						if i == r || //  same cell, skip
							s.isFilled(i, c) || // cell is filled already
							s.hasAlready(block, translateRCToBlock(i, c), v) || // if the block already has this v
							s.hasAlready(row, i, v) { // some other col in this row has this v already
							continue
						}
						otherRs = true
					}
					if !otherRs {
						// this v cannot be in any other rows
						s.setValue(r, c, v, "2dr")
						return true
					}
				}
			}
		}
	}
	return false
}

// check for complexity 3
// recursive: try to fill all values in all places, see if that is sane & solvable
func (s *Sudoku) findLevel3() (success bool) {
	for _, v := range orderedCandidates(s.occurences[:]) {
		for r := 0; r < 9; r++ {
			for c := 0; c < 9; c++ {
				if !s.isFilled(r, c) &&
					!s.hasAlready(row, r, v) &&
					!s.hasAlready(col, c, v) &&
					!s.hasAlready(block, translateRCToBlock(r, c), v) {

					// board[r][c]=v may be good, make a new board with that
					// and check if that can be solved
					clone := new(Sudoku)
					clone.LoadArray(s.board)
					clone.setValue(r, c, v, "3")
					if err := clone.Solve(0); err == nil {
						// yes it's done, "backfill" board and solution steps from there
						s.LoadArray(clone.board)
						s.steps = append(s.steps, clone.steps...)
						return true
					}
				}
			}
		}
	}
	return false
}

// count how many cells have this value in a row/col/block
func (s *Sudoku) count(dim dimension, where int, val int) (n int) {
	n = 0
	for i := 0; i < 9; i++ {
		r, c := translateIndextoRC(dim, where, i)
		if s.board[r][c] == val {
			n++
		}
	}
	return n
}

// does this row/col/block contain this value already?
func (s *Sudoku) hasAlready(dim dimension, where int, val int) bool {
	return s.count(dim, where, val) >= 1
}

// check if a cell has a value
func (s *Sudoku) isFilled(r, c int) bool {
	return s.board[r][c] != 0
}

func (s *Sudoku) setComplexity(complexity int) {
	if s.complexity < complexity {
		s.complexity = complexity
	}
}

func (s *Sudoku) Complexity() int {
	return s.complexity
}

func (s *Sudoku) Steps() *[]Step {
	return &s.steps
}

////////////////////// utils

// given a list of occurences, return an ordered list of candidates
// the more we have froma value, the erlier it should be on the list
// e.g. [X,9,9,9,0,6,5,4,3,2] -> [5,6,7,8,9,4]
func orderedCandidates(in []int) []int {
	res := make([]int, 0)
	for o := 8; o >= 0; o-- {
		for i := 1; i <= 9; i++ {
			if in[i] == o {
				res = append(res, i)
			}
		}
	}
	return res
}

// remove a particular element from a slice and return the rest
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
	case row:
		return where, index
	case col:
		return index, where
	case block:
		return index/3 + 3*(where/3), index%3 + 3*(where%3)
	}
	return 9, 9 // this will blow up
}

// given a row+col, return the correspoding block number
func translateRCToBlock(r int, c int) int {
	return (r/3)*3 + c/3
}
