package game

import (
	"time"
)

func (g *GameState) revealTile(x, y int) {
	if g.gameOver || g.gameWon {
		return
	}

	cell := &g.field[x][y]

	if cell.opened || cell.marked {
		return
	}

	cell.opened = true

	if cell.hasMine {
		g.gameOver = true
		g.finishedAt = time.Now()
		return
	}

	g.revealEmptyNeighbors(x, y)

	g.gameWon = g.isGameWon()

	if g.gameWon {
		g.finishedAt = time.Now()
	}
}

func (g *GameState) revealEmptyNeighbors(x, y int) {
	if g.field[x][y].minesAround != 0 {
		return
	}

	g.doForNeighbours(x, y, func(nx, ny int) {
		cell := &g.field[nx][ny]

		if cell.hasMine || cell.opened || cell.marked {
			return
		}

		g.revealTile(nx, ny)
	})
}

func (g *GameState) chordCell(x, y int) {
	cell := g.field[x][y]
	if !cell.opened {
		return
	}

	marked := 0
	g.doForNeighbours(x, y, func(nx, ny int) {
		if g.field[nx][ny].marked {
			marked++
		}
	})

	// The number of flags does not match the number of mines
	if marked != cell.minesAround {
		return
	}

	// Open neighboring cells
	g.doForNeighbours(x, y, func(nx, ny int) {
		if !g.field[nx][ny].marked && !g.field[nx][ny].opened {
			g.revealTile(nx, ny)
		}
	})
}

func (g *GameState) doForNeighbours(x, y int, do func(x, y int)) {
	// with diagonals
	dx := []int{-1, 0, 1, -1, 1, -1, 0, 1}
	dy := []int{-1, -1, -1, 0, 0, 1, 1, 1}

	for i := range len(dx) {
		nx := x + dx[i]
		ny := y + dy[i]

		if nx >= 0 && nx < g.rows && ny >= 0 && ny < g.cols {
			do(nx, ny)
		}
	}
}
