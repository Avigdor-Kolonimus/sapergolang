package game

const (
	beginnerRows  = 9
	beginnerCols  = 9
	beginnerMines = 10

	intermediateRows  = 16
	intermediateCols  = 16
	intermediateMines = 40

	expertRows  = 30
	expertCols  = 30
	expertMines = 99
)

type Level struct {
	rows  int
	cols  int
	mines int
}

var levels = map[int]Level{
	levelBeginner: {
		rows:  beginnerRows,
		cols:  beginnerCols,
		mines: beginnerMines,
	},
	levelIntermediate: {
		rows:  intermediateRows,
		cols:  intermediateCols,
		mines: intermediateMines,
	},
	levelExpert: {
		rows:  expertRows,
		cols:  expertCols,
		mines: expertMines,
	},
}

func (g *GameState) selectLevel(level int) {
	g.selectedLevel = level

	config := levels[level]

	g.rows = config.rows
	g.cols = config.cols
	g.mines = config.mines
}
