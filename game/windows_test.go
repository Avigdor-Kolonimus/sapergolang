package game

import (
	"testing"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func TestGetTextColor(t *testing.T) {
	tests := []struct {
		name      string
		neighbors int
		expected  rl.Color
	}{
		{
			name:      "one neighbor",
			neighbors: 1,
			expected:  rl.Blue,
		},
		{
			name:      "two neighbors",
			neighbors: 2,
			expected:  rl.Green,
		},
		{
			name:      "three neighbors",
			neighbors: 3,
			expected:  rl.Red,
		},
		{
			name:      "zero neighbors",
			neighbors: 0,
			expected:  rl.Black,
		},
		{
			name:      "more than three neighbors",
			neighbors: 4,
			expected:  rl.Black,
		},
		{
			name:      "negative neighbors",
			neighbors: -1,
			expected:  rl.Black,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getTextColor(tt.neighbors)

			if got != tt.expected {
				t.Errorf("getTextColor(%d) = %v, want %v",
					tt.neighbors, got, tt.expected)
			}
		})
	}
}
