package game

import (
	"fmt"

	gui "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	size = 30
)

func (g *GameState) getWidth() int {
	return size * g.cols
}

func (g *GameState) getHeight() int {
	return size*g.rows + size
}

func colorValue(c rl.Color) gui.PropertyValue {
	return gui.PropertyValue(rl.ColorToInt(c))
}

func setButtonStyle(selected bool) {
	if selected {
		gui.SetStyle(gui.BUTTON, gui.BASE_COLOR_NORMAL, colorValue(rl.Green))
		gui.SetStyle(gui.BUTTON, gui.BASE_COLOR_FOCUSED, colorValue(rl.Green))
		gui.SetStyle(gui.BUTTON, gui.BASE_COLOR_PRESSED, colorValue(rl.DarkGreen))

		return
	}

	gui.SetStyle(gui.BUTTON, gui.BASE_COLOR_NORMAL, colorValue(rl.Gray))
	gui.SetStyle(gui.BUTTON, gui.BASE_COLOR_FOCUSED, colorValue(rl.LightGray))
	gui.SetStyle(gui.BUTTON, gui.BASE_COLOR_PRESSED, colorValue(rl.DarkGray))
}

func (g *GameState) drawMenu() {
	var (
		rowSpacing float32 = 50
		baseY      float32 = 50
	)

	// BEGINNER
	setButtonStyle(g.selectedLevel == levelBeginner)

	if clicked := gui.Button(rl.NewRectangle(0, baseY, winWidth, size), "BEGINNER"); clicked {
		g.selectLevel(levelBeginner)
	}

	baseY += rowSpacing

	// INTERMEDIATE
	setButtonStyle(g.selectedLevel == levelIntermediate)

	if clicked := gui.Button(rl.NewRectangle(0, baseY, winWidth, size), "INTERMEDIATE"); clicked {
		g.selectLevel(levelIntermediate)
	}

	baseY += rowSpacing

	// EXPERT
	setButtonStyle(g.selectedLevel == levelExpert)

	if clicked := gui.Button(rl.NewRectangle(0, baseY, winWidth, size), "EXPERT"); clicked {
		g.selectLevel(levelExpert)
	}

	baseY += rowSpacing * 2

	// START
	setButtonStyle(false)

	if clicked := gui.Button(rl.NewRectangle(0, baseY, winWidth, size), "START"); clicked {
		g.start()
	}
}

func (g *GameState) drawCongratulations() {
	const (
		lineHeight int32 = 50
		fontSize   int32 = 40
	)

	text := "WELL DONE!"

	textWidth := rl.MeasureText(text, fontSize)
	textX := (winWidth - textWidth) / 2

	rl.DrawText(text, textX, lineHeight, fontSize, rl.White)

	buttonWidth := float32(200)
	buttonHeight := float32(size)

	buttonX := (float32(winWidth) - buttonWidth) / 2
	buttonY := float32(2 * lineHeight)

	clicked := gui.Button(rl.NewRectangle(buttonX, buttonY, buttonWidth, buttonHeight), "PLAY AGAIN")
	if clicked {
		g.reset()
	}
}

func (g *GameState) handleCellClick(x, y int, rect rl.Rectangle) {
	if !rl.CheckCollisionPointRec(rl.GetMousePosition(), rect) {
		return
	}

	left := rl.IsMouseButtonPressed(rl.MouseButtonLeft)
	right := rl.IsMouseButtonPressed(rl.MouseButtonRight)

	cell := &g.field[x][y]

	if left && right {
		if cell.opened {
			g.chordCell(x, y)
		}

		return
	}

	if right {
		if !cell.opened {
			cell.marked = !cell.marked
		}

		return
	}

	if left && !cell.opened && !cell.marked {
		g.revealTile(x, y)
	}
}

func (g *GameState) drawCell(x, y int, rect rl.Rectangle) {
	cell := &g.field[x][y]

	if cell.marked {
		rl.DrawText("M", 5+int32(x)*size, 5+int32(y)*size, 20, rl.Violet)

		return
	}

	if cell.opened {
		text := ""
		if cell.minesAround > 0 {
			text = fmt.Sprintf("%d", cell.minesAround)
		}

		rl.DrawText(text, 5+int32(x)*size, 5+int32(y)*size, 20, getTextColor(cell.minesAround))

		return
	}

	rl.DrawRectangleRec(rect, rl.Gray)
}

func (g *GameState) drawGameOverCell(x, y int) {
	var (
		text  string
		color rl.Color
	)

	cell := &g.field[x][y]
	if cell.hasMine {
		text = "*"
		color = rl.Red
	} else if cell.minesAround > 0 {
		text = fmt.Sprintf("%d", cell.minesAround)
		color = getTextColor(cell.minesAround)
	}

	rl.DrawText(text, 5+int32(x)*size, 5+int32(y)*size, 20, color)
}

func (g *GameState) drawField() {
	w := float32(g.getWidth())
	h := float32(g.getHeight())

	gui.StatusBar(
		rl.NewRectangle(0, h-size, w, size),
		g.getStatus(),
	)

	if restart := gui.Button(
		rl.NewRectangle(w-65, h-size+5, 60, size-10),
		"RESTART",
	); restart {
		g.reset()
		return
	}

	for x := range g.field {
		for y := range g.field[x] {
			if g.gameOver {
				g.drawGameOverCell(x, y)
				continue
			}

			rect := rl.NewRectangle(
				float32(x*size),
				float32(y*size),
				size,
				size,
			)

			g.handleCellClick(x, y, rect)
			g.drawCell(x, y, rect)
		}
	}
}
