package utils

import (
	"testing"
)

func TestColor(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		color     string
		expected  string
		wantError bool
	}{
		{"Red color", "Hello", "red", "\033[38;2;255;0;0mHello\033[0m", false},
		{"Green color", "World", "green", "\033[38;2;0;255;0mWorld\033[0m", false},
		{"Blue color", "Test", "blue", "\033[38;2;0;0;255mTest\033[0m", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Replace `tests` with the actual function you are testing
			result := Color(tt.input, tt.color)
			if tt.wantError {
				t.Errorf("expected an error, got nil")
			} else {
				if result != tt.expected {
					t.Errorf("expected %v, got %v", tt.expected, result)
				}
			}
		})
	}
}
