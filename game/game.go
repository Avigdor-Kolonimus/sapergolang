package game

import (
	"math/rand"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	winWidth  = 300
	winHeight = 450
)

func NewGame() *GameState {
	g := &GameState{
		rows:          9,
		cols:          9,
		mines:         10,
		selectedLevel: levelBeginner,
	}

	g.reset()

	return g
}

func (g *GameState) Start() {
	rl.InitWindow(winWidth, winHeight, "minesweeper")
	rl.SetTargetFPS(60)
	defer rl.CloseWindow()

	for !rl.WindowShouldClose() {
		if g.menu || g.gameWon {
			centerWindow(winWidth, winHeight)
		} else {
			centerWindow(g.getWidth(), g.getHeight())
		}

		rl.BeginDrawing()
		rl.ClearBackground(rl.DarkGray)

		if g.gameWon {
			g.drawCongratulations()
		} else if g.menu {
			g.drawMenu()
		} else {
			g.drawField()
		}

		rl.EndDrawing()
	}
}

func (g *GameState) reset() {
	g.menu = true
	g.gameOver = false
	g.gameWon = false
}

func (g *GameState) start() {
	// Build grid
	g.field = make([][]point, g.rows)
	for x := range g.rows {
		g.field[x] = make([]point, g.cols)
		for y := range g.cols {
			g.field[x][y] = point{}
		}
	}

	// Plant mines
	m := g.mines
	for m > 0 {
		x, y := rand.Intn(g.rows), rand.Intn(g.cols)

		// make sure placements are unique
		if g.field[x][y].hasMine {
			continue
		}

		g.field[x][y].hasMine = true
		// mark neighbours
		g.doForNeighbours(x, y, func(x, y int) {
			g.field[x][y].minesAround++
		})
		m--
	}

	g.menu = false
	g.startedAt = time.Now()
}
