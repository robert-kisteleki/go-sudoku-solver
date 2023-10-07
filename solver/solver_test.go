package solver

import (
	"testing"
)

// Test complexity 1(a)
func TestSolverLevel1(t *testing.T) {
	in := new(Sudoku)

	err := in.LoadString(`
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
	_ = in.Solve(1)

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
		t.Fatalf("incorrect solution for case 1a:\n%s", s)
	}

	err = in.LoadString(`
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
	_ = in.Solve(1)

	s = "\n" + in.String()
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
		t.Fatalf("incorrect solution for case 1b:\n%s", s)
	}

	err = in.LoadString(`
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
	_ = in.Solve(1)

	s = "\n" + in.String()
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
		t.Fatalf("incorrect solution for case 1c:\n%s", s)
	}
}

func TestSolverLevel2(t *testing.T) {
	in := new(Sudoku)

	err := in.LoadString(`
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
	_ = in.Solve(1)

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
		t.Fatalf("incorrect solution for case 2a:\n%s", s)
	}

	err = in.LoadString(`
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
	_ = in.Solve(1)

	s = "\n" + in.String()
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
		t.Fatalf("incorrect solution for case 2b:\n%s", s)
	}

	err = in.LoadString(`
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
	_ = in.Solve(1)

	s = "\n" + in.String()
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
		t.Fatalf("incorrect solution for case 2c1:\n%s", s)
	}

	err = in.LoadString(`
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
	_ = in.Solve(1)

	s = "\n" + in.String()
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
		t.Fatalf("incorrect solution for case 2c2:\n%s", s)
	}

	err = in.LoadString(`
+---+---+---+
|   |   |   |
|   |   |   |
|   |   |   |
+---+---+---+
|   |   |   |
|   |   |   |
|   |   |   |
+---+---+---+
|   | 1 |   |
|   |   |   |
|X23|   |456|
+---+---+---+
`)
	if err != nil {
		panic(err)
	}
	_ = in.Solve(1)

	s = "\n" + in.String()
	if s != `
+---+---+---+
|   |   |   |
|   |   |   |
|   |   |   |
+---+---+---+
|   |   |   |
|   |   |   |
|   |   |   |
+---+---+---+
|   | 1 |   |
|   |   |   |
|123|   |456|
+---+---+---+
` {
		t.Fatalf("incorrect solution for case 2d1:\n%s", s)
	}

	err = in.LoadString(`
+---+---+---+
|   |   |   |
|   |   |   |
|   |   |   |
+---+---+---+
|   |   |   |
|   |   |   |
|  1|   |   |
+---+---+---+
|   | 1 |   |
|   |   |   |
|X2 |   |456|
+---+---+---+
`)
	if err != nil {
		panic(err)
	}
	_ = in.Solve(1)

	s = "\n" + in.String()
	if s != `
+---+---+---+
|   |   |   |
|   |   |   |
|   |   |   |
+---+---+---+
|   |   |   |
|   |   |   |
|  1|   |   |
+---+---+---+
|   | 1 |   |
|   |   |   |
|12 |   |456|
+---+---+---+
` {
		t.Fatalf("incorrect solution for case 2d2:\n%s", s)
	}

	err = in.LoadString(`
+---+---+---+
|   |   |   |
|   |   |   |
|   |   |  1|
+---+---+---+
|   |   |   |
|   |   |   |
|  1|   |   |
+---+---+---+
|   | 1 |   |
|   |   |   |
|X2 |   |45 |
+---+---+---+
`)
	if err != nil {
		panic(err)
	}
	_ = in.Solve(1)

	s = "\n" + in.String()
	if s != `
+---+---+---+
|   |   |   |
|   |   |   |
|   |   |  1|
+---+---+---+
|   |   |   |
|   |   |   |
|  1|   |   |
+---+---+---+
|   | 1 |   |
|   |   |   |
|12 |   |45 |
+---+---+---+
` {
		t.Fatalf("incorrect solution for case 2d3:\n%s", s)
	}
}

func TestSolverLevel3(t *testing.T) {
	in := new(Sudoku)

	err := in.LoadString(`
+---+---+---+
| 26| 1 | 8 |
|   |624|   |
|  3| 5 |2 6|
+---+---+---+
|345|197|862|
|961|   |4  |
|872|465|913|
+---+---+---+
|23 |5 1|69 |
|6  |   |  8|
| 5 | 76| 2 |
+---+---+---+
`)
	if err != nil {
		panic(err)
	}
	_ = in.Solve(1)

	s := "\n" + in.String()
	if s != `
+---+---+---+
|426|713|589|
|589|624|731|
|713|859|246|
+---+---+---+
|345|197|862|
|961|238|475|
|872|465|913|
+---+---+---+
|234|581|697|
|697|342|158|
|158|976|324|
+---+---+---+
` {
		t.Fatalf("incorrect solution for case 3:\n%s", s)
	}
}

func TestInputSanity(t *testing.T) {
	in := new(Sudoku)

	err := in.LoadString(`
+---+---+---+
|1  |   |  1|
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
	if err == nil {
		t.Fatalf("duplicate values are accepted on input (row)")
	}

	err = in.LoadString(`
+---+---+---+
|1  |   |   |
|   |   |   |
|   |   |   |
+---+---+---+
|   |   |   |
|   |   |   |
|   |   |   |
+---+---+---+
|   |   |   |
|   |   |   |
|1  |   |   |
+---+---+---+
`)
	if err == nil {
		t.Fatalf("duplicate values are accepted on input (col)")
	}

	err = in.LoadString(`
+---+---+---+
|1  |   |   |
|   |   |   |
|  1|   |   |
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
	if err == nil {
		t.Fatalf("duplicate values are accepted on input (block)")
	}
}

func TestUnsolvable(t *testing.T) {
	in := new(Sudoku)

	err := in.LoadString(`
+---+---+---+
|1  |   |   |
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
		t.Fatalf("solved and unsolvable puzzle?")
	}
}
