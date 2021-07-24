package sodokuSolver

type Cell struct {
	val      int
	possible []int
}

type Square struct {
	grid      [3][3]Cell
	available []int
}

type Row struct {
	row       [9]Cell
	available []int
}

type Column struct {
	col       [9]Cell
	available []int
}

type Board struct {
	squares [3][3]Square
	rows    [9]Row
	cols    [9]Column
}
