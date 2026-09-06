package game

import (
	"testing"
	"time"
)

func TestGameState_IsGameWon(t *testing.T) {
	tests := []struct {
		name string
		game GameState
		want bool
	}{
		{
			name: "all safe cells are opened",
			game: GameState{
				rows:  2,
				cols:  2,
				mines: 1,
				field: [][]point{
					{
						{opened: true},
						{opened: true},
					},
					{
						{opened: true},
						{hasMine: true},
					},
				},
			},
			want: true,
		},
		{
			name: "not all safe cells are opened",
			game: GameState{
				rows:  2,
				cols:  2,
				mines: 1,
				field: [][]point{
					{
						{opened: true},
						{opened: true},
					},
					{
						{opened: false},
						{hasMine: true},
					},
				},
			},
			want: false,
		},
		{
			name: "mine is opened",
			game: GameState{
				rows:  2,
				cols:  2,
				mines: 1,
				field: [][]point{
					{
						{opened: true},
						{opened: true},
					},
					{
						{opened: true},
						{opened: true, hasMine: true},
					},
				},
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.game.isGameWon()

			if got != tt.want {
				t.Errorf("isGameWon() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalculateStatus(t *testing.T) {
	startedAt := time.Date(2026, 1, 1, 12, 0, 0, 0, time.UTC)

	tests := []struct {
		name       string
		gameOver   bool
		gameWon    bool
		fps        int32
		startedAt  time.Time
		finishedAt time.Time
		want       string
	}{
		{
			name:       "game over",
			gameOver:   true,
			gameWon:    false,
			fps:        60,
			startedAt:  startedAt,
			finishedAt: startedAt.Add(5 * time.Second),
			want:       "FPS: 60, TIME: 5.00",
		},
		{
			name:       "game won",
			gameOver:   false,
			gameWon:    true,
			fps:        120,
			startedAt:  startedAt,
			finishedAt: startedAt.Add(10 * time.Second),
			want:       "FPS: 120, TIME: 10.00",
		},
		{
			name:       "game in progress",
			gameOver:   false,
			gameWon:    false,
			fps:        60,
			startedAt:  time.Now().Add(-5 * time.Second),
			finishedAt: time.Time{},
			want:       "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := calculateStatus(
				tt.gameOver,
				tt.gameWon,
				tt.fps,
				tt.startedAt,
				tt.finishedAt,
			)

			if tt.want != "" && got != tt.want {
				t.Errorf("calculateStatus() = %q, want %q", got, tt.want)
			}
		})
	}
}
