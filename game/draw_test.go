package game

import (
	"testing"
)

func TestGameState_GetWidth(t *testing.T) {
	tests := []struct {
		name string
		game GameState
		want int
	}{
		{
			name: "10 columns",
			game: GameState{cols: 10},
			want: size * 10,
		},
		{
			name: "zero columns",
			game: GameState{cols: 0},
			want: 0,
		},
		{
			name: "one column",
			game: GameState{cols: 1},
			want: size,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.game.getWidth()

			if got != tt.want {
				t.Errorf("getWidth() = %d, want %d", got, tt.want)
			}
		})
	}
}

func TestGameState_GetHeight(t *testing.T) {
	tests := []struct {
		name string
		game GameState
		want int
	}{
		{
			name: "10 rows",
			game: GameState{rows: 10},
			want: size*10 + size,
		},
		{
			name: "zero rows",
			game: GameState{rows: 0},
			want: size,
		},
		{
			name: "one row",
			game: GameState{rows: 1},
			want: size * 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.game.getHeight()

			if got != tt.want {
				t.Errorf("getHeight() = %d, want %d", got, tt.want)
			}
		})
	}
}
