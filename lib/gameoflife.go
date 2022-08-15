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
		if liveCells[i][0] >= 0 && liveCells[i][0] < tempMap.rows && liveCells[i][1] >= 0 && liveCells[i][1] < tempMap.columns {
			tempMap.Grid[liveCells[i][0]][liveCells[i][1]] = true
		} else {
			panic("Index out of range")
		}
	}
	tempMap.liveCells = liveCellCount
}

func (tempMap Map) GetNumberOfLiveNeighbours(row, column int) int {
	if row < 0 || row >= tempMap.rows || column < 0 || column >= tempMap.columns {
		panic("Index out of range")
	}
	neighbours := 0
	for i := row - 1; i <= row+1; i++ {
		for j := column - 1; j <= column+1; j++ {
			if i >= 0 && j >= 0 && i < tempMap.rows && j < tempMap.columns {
				if tempMap.Grid[i][j] == true {
					neighbours++
				}
			}
		}
	}
	if tempMap.Grid[row][column] == true {
		return neighbours - 1
	}
	return neighbours
}

func (tempMap Map) GetNextGeneration() Map {
	if tempMap.liveCells == 0 {
		return tempMap
	}
	nextGenMap := CreateMap(tempMap.rows, tempMap.columns)
	nextGenMap.liveCells = tempMap.liveCells
	for i := 0; i < tempMap.rows; i++ {
		for j := 0; j < tempMap.columns; j++ {
			neighbours := tempMap.GetNumberOfLiveNeighbours(i, j)
			if tempMap.Grid[i][j] == false {
				if neighbours == 3 {
					nextGenMap.Grid[i][j] = true
					nextGenMap.liveCells += 1
				}
			} else {
				if neighbours < 2 || neighbours > 3 {
					nextGenMap.Grid[i][j] = false
					nextGenMap.liveCells -= 1
				} else {
					nextGenMap.Grid[i][j] = true
				}
			}
		}
	}
	return nextGenMap
}
