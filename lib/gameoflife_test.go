package gameoflife

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateMap(t *testing.T) {
	t.Run("check if map can be created", func(t *testing.T) {
		assert.IsType(t, Map{}, CreateMap(3, 3))
	})

	t.Run("check if it raises exception when rows or column passed is negative", func(t *testing.T) {
		assert.Panics(t, func() {
			CreateMap(2, -2)
		})
	})

	t.Run("check if it does not raise exception when rows or columns passed is non-negative", func(t *testing.T) {
		assert.NotPanics(t, func() {
			CreateMap(3, 4)
		})
		assert.NotPanics(t, func() {
			CreateMap(0, 9)
		})
	})

	rows := 2
	columns := 3
	tempMap := CreateMap(rows, columns)

	t.Run("check if rows are initialised correctly", func(t *testing.T) {
		assert.Equal(t, 2, tempMap.rows)
	})

	t.Run("check if columns are initialised coreectly", func(t *testing.T) {
		assert.Equal(t, 3, tempMap.columns)
	})

	t.Run("check if live cells are initialized to zero", func(t *testing.T) {
		assert.Equal(t, 0, tempMap.liveCells)
	})

	t.Run("check if grid is created and all cells are initialised as dead:false", func(t *testing.T) {
		for i := 0; i < rows; i++ {
			for j := 0; j < columns; j++ {
				assert.Equal(t, false, tempMap.Grid[i][j])
			}
		}
	})
}

func TestInitialiseGrid(t *testing.T) {
	rows := 3
	columns := 3
	tempMap := CreateMap(rows, columns)
	NumberOfLiveCells := 2
	liveCells := make([][]int, NumberOfLiveCells)
	for i := range liveCells {
		liveCells[i] = make([]int, 2)
	}
	liveCells[0][0] = 1
	liveCells[0][1] = 1
	liveCells[1][0] = 2
	liveCells[1][1] = 2

	t.Run("check if we can initialise Grid", func(t *testing.T) {
		assert.NotPanics(t, func() {
			tempMap.InitialiseGrid(NumberOfLiveCells, liveCells)
		})
	})

	t.Run("check if Grid is properly initialised with passed live cells", func(t *testing.T) {
		for i := 0; i < NumberOfLiveCells; i++ {
			assert.Equal(t, true, tempMap.Grid[liveCells[i][0]][liveCells[i][1]])
		}
	})

	liveCells[1][0] = 3
	liveCells[1][1] = 3

	t.Run("check if it panics when index of live cells go out of range", func(t *testing.T) {
		assert.Panics(t, func() {
			tempMap.InitialiseGrid(NumberOfLiveCells, liveCells)
		})
	})

	t.Run("check if number of live cells is initialised coorectly", func(t *testing.T) {
		assert.Equal(t, NumberOfLiveCells, tempMap.liveCells)
	})
}
