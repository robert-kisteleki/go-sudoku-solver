package solver

import (
	"testing"
)

// Test complexity 1(a)
func TestSolverLevel1a(t *testing.T) {

	in := new(Sudoku)
	err := in.Load(`
+---+---+---+
|123|456|78X|
|   |   |   |
|   |   |   |
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
	_ = in.Solve()

	s := "\n" + in.String()
	if s != `
+---+---+---+
|123|456|789|
|   |   |   |
|   |   |   |
+---+---+---+
|   |   |   |
|   |   |   |
|   |   |   |
+---+---+---+
|   |   |   |
|   |   |   |
|   |   |   |
+---+---+---+
` {
		t.Fatalf("incorrect soution for case 1a:\n%s", s)
	}
}

func TestSolverLevel1b(t *testing.T) {

	in := new(Sudoku)
	err := in.Load(`
+---+---+---+
|1  |   |   |
|2  |   |   |
|3  |   |   |
+---+---+---+
|4  |   |   |
|5  |   |   |
|6  |   |   |
+---+---+---+
|7  |   |   |
|8  |   |   |
|X  |   |   |
+---+---+---+
`)
	if err != nil {
		panic(err)
	}
	_ = in.Solve()

	s := "\n" + in.String()
	if s != `
+---+---+---+
|1  |   |   |
|2  |   |   |
|3  |   |   |
+---+---+---+
|4  |   |   |
|5  |   |   |
|6  |   |   |
+---+---+---+
|7  |   |   |
|8  |   |   |
|9  |   |   |
+---+---+---+
` {
		t.Fatalf("incorrect soution for case 1b:\n%s", s)
	}
}

func TestSolverLevel1c(t *testing.T) {

	in := new(Sudoku)
	err := in.Load(`
+---+---+---+
|123|   |   |
|456|   |   |
|78X|   |   |
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
	_ = in.Solve()

	s := "\n" + in.String()
	if s != `
+---+---+---+
|123|   |   |
|456|   |   |
|789|   |   |
+---+---+---+
|   |   |   |
|   |   |   |
|   |   |   |
+---+---+---+
|   |   |   |
|   |   |   |
|   |   |   |
+---+---+---+
` {
		t.Fatalf("incorrect soution for case 1c:\n%s", s)
	}
}

func TestSolverLevel2a(t *testing.T) {

	in := new(Sudoku)
	err := in.Load(`
+---+---+---+
|1  |   |   |
|   |1  |   |
|   |   |  X|
+---+---+---+
|   |   |1  |
|   |   |   |
|   |   |   |
+---+---+---+
|   |   | 1 |
|   |   |   |
|   |   |   |
+---+---+---+
`)
	if err != nil {
		panic(err)
	}
	_ = in.Solve()

	s := "\n" + in.String()
	if s != `
+---+---+---+
|1  |   |   |
|   |1  |   |
|   |   |  1|
+---+---+---+
|   |   |1  |
|   |   |   |
|   |   |   |
+---+---+---+
|   |   | 1 |
|   |   |   |
|   |   |   |
+---+---+---+
` {
		t.Fatalf("incorrect soution for case 2a:\n%s", s)
	}
}

func TestSolverLevel2b(t *testing.T) {

	in := new(Sudoku)
	err := in.Load(`
+---+---+---+
|1  |   |   |
|   |1  |   |
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
	_ = in.Solve()

	s := "\n" + in.String()
	if s != `
+---+---+---+
|1  |   |   |
|   |1  |   |
|   |   |231|
+---+---+---+
|   |   |   |
|   |   |   |
|   |   |   |
+---+---+---+
|   |   |   |
|   |   |   |
|   |   |   |
+---+---+---+
` {
		t.Fatalf("incorrect soution for case 2b:\n%s", s)
	}
}

func TestSolverLevel2c1(t *testing.T) {

	in := new(Sudoku)
	err := in.Load(`
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
	_ = in.Solve()

	s := "\n" + in.String()
	if s != `
+---+---+---+
|1  |   |   |
|   |   |456|
|   |   |231|
+---+---+---+
|   |   |   |
|   |   |   |
|   |   |   |
+---+---+---+
|   |   |   |
|   |   |   |
|   |   |   |
+---+---+---+
` {
		t.Fatalf("incorrect soution for case 2c1:\n%s", s)
	}
}

func TestSolverLevel2c2(t *testing.T) {

	in := new(Sudoku)
	err := in.Load(`
+---+---+---+
|1  |   |   |
|   |   |   |
|   |   |   |
+---+---+---+
|   |   |   |
|   |   |   |
|   |   |   |
+---+---+---+
| 42|   |   |
| 53|   |   |
| 6X|   |   |
+---+---+---+
`)
	if err != nil {
		panic(err)
	}
	_ = in.Solve()

	s := "\n" + in.String()
	if s != `
+---+---+---+
|1  |   |   |
|   |   |   |
|   |   |   |
+---+---+---+
|   |   |   |
|   |   |   |
|   |   |   |
+---+---+---+
| 42|   |   |
| 53|   |   |
| 61|   |   |
+---+---+---+
` {
		t.Fatalf("incorrect soution for case 2c2:\n%s", s)
	}
}
