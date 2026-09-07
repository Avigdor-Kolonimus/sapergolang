package game

import (
	"testing"
)

func TestGameState_SelectLevel(t *testing.T) {
	tests := []struct {
		name  string
		level int
		want  Level
	}{
		{
			name:  "beginner",
			level: levelBeginner,
			want:  levels[levelBeginner],
		},
		{
			name:  "intermediate",
			level: levelIntermediate,
			want:  levels[levelIntermediate],
		},
		{
			name:  "expert",
			level: levelExpert,
			want:  levels[levelExpert],
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GameState{}

			g.selectLevel(tt.level)

			if g.selectedLevel != tt.level {
				t.Errorf("selectedLevel = %d, want %d", g.selectedLevel, tt.level)
			}

			got := Level{
				rows:  g.rows,
				cols:  g.cols,
				mines: g.mines,
			}

			if got != tt.want {
				t.Errorf("level = %+v, want %+v", got, tt.want)
			}
		})
	}
}
