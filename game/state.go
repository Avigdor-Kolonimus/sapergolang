package game

import (
	"fmt"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type point struct {
	hasMine     bool
	opened      bool
	marked      bool
	minesAround int
}

type GameState struct {
	menu       bool
	gameOver   bool
	gameWon    bool
	rows       int
	cols       int
	mines      int
	field      [][]point
	startedAt  time.Time
	finishedAt time.Time
}

func (g *GameState) isGameWon() bool {
	open := 0
	total := int(g.rows * g.cols)

	for x := range g.rows {
		for y := range g.cols {
			if g.field[x][y].opened {
				open++
			}
		}
	}

	return open == total-int(g.mines)
}

func (g *GameState) getStatus() string {
	fps := rl.GetFPS()
	var elapsed time.Duration
	if g.gameOver || g.gameWon {
		elapsed = g.finishedAt.Sub(g.startedAt)
	} else {
		elapsed = time.Since(g.startedAt)
	}

	return fmt.Sprintf("FPS: %d, TIME: %.2f", fps, elapsed.Seconds())
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
