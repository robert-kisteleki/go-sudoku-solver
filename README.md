# Sudoku Solver in Go

A command line tool to solve sudokus.

The package part (`solver/`) may be usable separately as well.

# Use

The compiled program expects a 9x9 table in input, like:

```
760000015
412030689
305000704
600308001
150060093
200701006
804000102
976020438
520000067
```

or

```
+---+---+---+
|   |791|   |
| 17|4 8|35 |
| 2 |5 6|781|
+---+---+---+
| 98|   |52 |
|5  |   |  7|
| 76|   |89 |
+---+---+---+
| 4 |6 2| 3 |
| 81|9 5|24 |
|   |314|   |
+---+---+---+
```

Solve by using:

```
sudoku < solvethis.txt
```

The input can contain pretty-printed sudokus, or ones where missing cells
are marked with X or 0 or spaces. There should be 9 lines of useful stuff.
