package cgol

import "math/rand"

type CgolManager struct {
	xSize              int
	ySize              int
	currentGeneration  [][]int
	tempNextGeneration [][]int
}

func NewManager(xSize int, ySize int) CgolManager {
	m := CgolManager{
		xSize:              xSize,
		ySize:              ySize,
		currentGeneration:  makeMatrix(xSize, ySize),
		tempNextGeneration: makeMatrix(xSize, ySize),
	}
	m.initSeed()
	return m
}

func makeMatrix(rows, cols int) [][]int {
	m := make([][]int, rows)
	for i := range m {
		m[i] = make([]int, cols)
	}
	return m
}

func (m *CgolManager) initSeed() {
	for x := 0; x < m.xSize; x++ {
		for y := 0; y < m.ySize; y++ {
			rn := rand.Intn(101) + 1
			if rn < 70 {
				m.currentGeneration[x][y] = 0
			} else {
				m.currentGeneration[x][y] = 1
			}
		}
	}
}

func (m *CgolManager) countLiveNeighbours(x int, y int) int {
	liveNeighbours := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			outOfBounds := x+i < 0 || x+i >= m.xSize || y+j < 0 || y+j >= m.ySize
			sameCell := x+i == m.xSize && y+j == m.ySize
			if outOfBounds || sameCell {
				continue
			}
			liveNeighbours += m.currentGeneration[x+i][y+j]
		}
	}
	return liveNeighbours
}

func (m *CgolManager) ComputeNextGeneration() {
	for x := 0; x < m.xSize; x++ {
		for y := 0; y < m.ySize; y++ {
			liveNeighbours := m.countLiveNeighbours(x, y)
			if m.currentGeneration[x][y] == 1 && liveNeighbours < 2 {
				m.tempNextGeneration[x][y] = 0
			} else if m.currentGeneration[x][y] == 1 && liveNeighbours > 3 {
				m.tempNextGeneration[x][y] = 0
			} else if m.currentGeneration[x][y] == 0 && liveNeighbours == 3 {
				m.tempNextGeneration[x][y] = 1
			} else {
				m.tempNextGeneration[x][y] = m.currentGeneration[x][y]
			}
		}
	}
	for x := 0; x < m.xSize; x++ {
		for y := 0; y < m.ySize; y++ {
			m.currentGeneration[x][y] = m.tempNextGeneration[x][y]
		}
	}
}
