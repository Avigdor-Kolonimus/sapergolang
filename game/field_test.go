package game

import (
	"testing"
)

func TestGameState_RevealTile_Mine(t *testing.T) {
	g := &GameState{
		rows: 1,
		cols: 1,
		field: [][]point{
			{{hasMine: true}},
		},
	}

	g.revealTile(0, 0)

	if !g.field[0][0].opened {
		t.Error("mine should be opened")
	}

	if !g.gameOver {
		t.Error("gameOver should be true")
	}

	if g.finishedAt.IsZero() {
		t.Error("finishedAt should be set")
	}
}

func TestGameState_RevealTile_MarkedCell(t *testing.T) {
	g := &GameState{
		rows: 1,
		cols: 1,
		field: [][]point{
			{{marked: true}},
		},
	}

	g.revealTile(0, 0)

	if g.field[0][0].opened {
		t.Error("marked cell should not be opened")
	}
}

func TestGameState_RevealTile_OpenedCell(t *testing.T) {
	g := &GameState{
		rows: 1,
		cols: 1,
		field: [][]point{
			{{opened: true}},
		},
	}

	g.revealTile(0, 0)

	if !g.field[0][0].opened {
		t.Error("opened cell should remain opened")
	}
}

func TestGameState_RevealTile_GameOver(t *testing.T) {
	g := &GameState{
		rows:     1,
		cols:     1,
		gameOver: true,
		field: [][]point{
			{{}},
		},
	}

	g.revealTile(0, 0)

	if g.field[0][0].opened {
		t.Error("cell should not be opened after game over")
	}
}

func TestGameState_RevealTile_GameWon(t *testing.T) {
	g := &GameState{
		rows: 1,
		cols: 1,
		field: [][]point{
			{{}},
		},
	}

	g.revealTile(0, 0)

	if !g.gameWon {
		t.Error("gameWon should be true")
	}

	if g.finishedAt.IsZero() {
		t.Error("finishedAt should be set")
	}
}

func TestGameState_RevealTile_OpensNeighbors(t *testing.T) {
	g := &GameState{
		rows: 2,
		cols: 2,
		field: [][]point{
			{
				{minesAround: 0},
				{minesAround: 1},
			},
			{
				{minesAround: 1},
				{hasMine: true, minesAround: 2},
			},
		},
	}

	g.revealTile(0, 0)

	if !g.field[0][0].opened {
		t.Error("cell [0][0] should be opened")
	}

	if !g.field[0][1].opened {
		t.Error("cell [0][1] should be opened")
	}

	if !g.field[1][0].opened {
		t.Error("cell [1][0] should be opened")
	}

	if g.field[1][1].opened {
		t.Error("mine [1][1] should not be opened")
	}
}

func TestGameState_ChordCell_WrongFlags(t *testing.T) {
	g := &GameState{
		rows: 3,
		cols: 3,
		field: [][]point{
			{{}, {}, {}},
			{{}, {opened: true, minesAround: 2}, {}},
			{{}, {marked: true}, {}},
		},
	}

	g.chordCell(1, 1)

	if g.field[0][0].opened {
		t.Error("neighbor [0][0] should not be opened")
	}

	if g.field[0][1].opened {
		t.Error("neighbor [0][1] should not be opened")
	}

	if g.field[0][2].opened {
		t.Error("neighbor [0][2] should not be opened")
	}

	if !g.field[1][1].opened {
		t.Error("center cell should remain opened")
	}
}

func TestGameState_ChordCell_OpensNeighbors(t *testing.T) {
	g := &GameState{
		rows: 3,
		cols: 3,
		field: [][]point{
			{{}, {marked: true}, {}},
			{{}, {opened: true, minesAround: 1}, {}},
			{{}, {}, {}},
		},
	}

	g.chordCell(1, 1)

	if !g.field[0][0].opened {
		t.Error("neighbor [0][0] should be opened")
	}

	if g.field[0][1].opened {
		t.Error("marked neighbor [0][1] should not be opened")
	}

	if !g.field[0][2].opened {
		t.Error("neighbor [0][2] should be opened")
	}
}

func TestGameState_ChordCell_ClosedCell(t *testing.T) {
	g := &GameState{
		rows: 2,
		cols: 2,
		field: [][]point{
			{{}, {}},
			{{}, {}},
		},
	}

	g.chordCell(0, 0)

	for x := range g.field {
		for y := range g.field[x] {
			if g.field[x][y].opened {
				t.Errorf("field[%d][%d] should not be opened", x, y)
			}
		}
	}
}

func TestGameState_DoForNeighbours(t *testing.T) {
	tests := []struct {
		name string
		x    int
		y    int
		want int
	}{
		{
			name: "corner",
			x:    0,
			y:    0,
			want: 3,
		},
		{
			name: "edge",
			x:    0,
			y:    1,
			want: 5,
		},
		{
			name: "center",
			x:    1,
			y:    1,
			want: 8,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GameState{
				rows: 3,
				cols: 3,
			}

			count := 0

			g.doForNeighbours(tt.x, tt.y, func(x, y int) {
				count++
			})

			if count != tt.want {
				t.Errorf("neighbors = %d, want %d", count, tt.want)
			}
		})
	}
}
