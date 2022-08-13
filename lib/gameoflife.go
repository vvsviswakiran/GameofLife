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

func (tempMap *Map) InitialiseGrid(liveCellCount int, liveCells [][]int) {
	for i := 0; i < liveCellCount; i++ {
		if liveCells[i][0] < tempMap.rows && liveCells[i][1] < tempMap.columns {
			tempMap.Grid[liveCells[i][0]][liveCells[i][1]] = true
		} else {
			panic("Index out of range")
		}
	}
	tempMap.liveCells = liveCellCount
}
