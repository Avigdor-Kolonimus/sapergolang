package game

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func getTextColor(neighbors int) rl.Color {
	switch neighbors {
	case 1:
		return rl.Blue
	case 2:
		return rl.Green
	case 3:
		return rl.Red
	default:
		return rl.Black
	}
}

func centerWindow(width, height int) {
	rl.SetWindowSize(width, height)
	monitorWidth := rl.GetMonitorWidth(0)
	monitorHeight := rl.GetMonitorHeight(0)
	rl.SetWindowPosition((monitorWidth-width)/2, (monitorHeight-height)/2)
}
