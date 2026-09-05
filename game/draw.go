package game

import (
	"fmt"
	"time"

	gui "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	size = 30

	biginnerRows  = 9
	biginnerCols  = 9
	biginnerMines = 10

	intermediateRows  = 16
	intermediateCols  = 16
	intermediateMines = 40

	expertRows  = 30
	expertCols  = 30
	expertMines = 99
)

func (g *GameState) getWidth() int {
	return size * g.cols
}

func (g *GameState) getHeight() int {
	return size*g.rows + size
}

func (g *GameState) revealTile(x, y int) {
	if g.field[x][y].opened {
		return
	}

	g.field[x][y].opened = true

	if g.field[x][y].hasMine {
		g.gameOver = true
		g.finishedAt = time.Now()

		return
	}

	g.gameWon = g.isGameWon()

	// No neighbors, reveal all adjacent tiles recursively
	if g.field[x][y].minesAround == 0 {
		g.doForNeighbours(x, y, func(nx, ny int) {
			g.revealTile(nx, ny)
		})
	}
}

func (g *GameState) drawMenu() {
	var (
		rowSpacing float32 = 50
		baseY      float32 = 50
	)

	if clicked := gui.Button(rl.NewRectangle(0, baseY, winWidth, size), "BEGINNER"); clicked {
		g.rows = biginnerRows
		g.cols = biginnerCols
		g.mines = biginnerMines
	}
	baseY += rowSpacing

	if clicked := gui.Button(rl.NewRectangle(0, baseY, winWidth, size), "INTERMEDIATE"); clicked {
		g.rows = intermediateRows
		g.cols = intermediateCols
		g.mines = intermediateMines
	}
	baseY += rowSpacing

	if clicked := gui.Button(rl.NewRectangle(0, baseY, winWidth, size), "EXPERT"); clicked {
		g.rows = expertRows
		g.cols = expertCols
		g.mines = expertMines
	}
	baseY += rowSpacing * 2

	if clicked := gui.Button(rl.NewRectangle(0, baseY, winWidth, size), "START"); clicked {
		g.start()
	}
}

func (g *GameState) drawCongratulations() {
	var lineHeight int32 = 50
	w := winWidth

	if g.gameWon {
		rl.DrawText("WELL DONE !", 0, lineHeight, size, rl.White)
	}

	clicked := gui.Button(rl.NewRectangle(0, float32(2*lineHeight), float32(w), size), "PLAY AGAIN")
	if clicked {
		g.reset()
	}
}

func (g *GameState) drawField() {
	w := float32(g.getWidth())
	h := float32(g.getHeight())

	gui.StatusBar(rl.NewRectangle(0, h-size, w, size), g.getStatus())
	if restart := gui.Button(rl.NewRectangle(w-65, h-size+5, 60, size-10), "RESTART"); restart {
		g.reset()

		return
	}

	for x := range g.field {
		for y := range g.field[x] {
			if g.gameOver {
				var (
					text  string
					color rl.Color
				)

				if g.field[x][y].hasMine {
					text = "*"
					color = rl.Red
				} else if g.field[x][y].minesAround > 0 {
					color = getTextColor(g.field[x][y].minesAround)
					text = fmt.Sprintf("%d", g.field[x][y].minesAround)
				}

				rl.DrawText(text, 5+int32(x)*size, 5+int32(y)*size, 20, color)
				continue
			}

			rect := rl.NewRectangle(float32(x*size), float32(y*size), size, size)

			// Mark on right mouse button
			if rl.IsMouseButtonPressed(rl.MouseButtonRight) {
				if rl.CheckCollisionPointRec(rl.GetMousePosition(), rect) {
					if !g.field[x][y].opened {
						g.field[x][y].marked = !g.field[x][y].marked
					}
				}
			}

			if g.field[x][y].marked {
				rl.DrawText("M", 5+int32(x)*size, 5+int32(y)*size, 20, rl.Violet)
			} else if g.field[x][y].opened {
				text := ""
				if g.field[x][y].minesAround > 0 {
					text = fmt.Sprintf("%d", g.field[x][y].minesAround)
				}

				rl.DrawText(text, 5+int32(x)*size, 5+int32(y)*size, 20, getTextColor(g.field[x][y].minesAround))
			} else {
				if open := gui.Button(rect, ""); open {
					g.revealTile(x, y)
				}
			}
		}
	}
}
