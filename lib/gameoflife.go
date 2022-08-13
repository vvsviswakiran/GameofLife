package gameoflife

type Map struct {
	rows      int
	columns   int
	liveCells int
	Grid      [][]bool
}

func CreateMap(rows, columns int) Map {
	if rows < 0 || columns < 0 {
		panic("rows and columns must be positive")
	}
	tempMap := make([][]bool, rows)
	for i := range tempMap {
		tempMap[i] = make([]bool, columns)
	}
	return Map{rows, columns, 0, tempMap}
}
