package utils

import (
	"testing"
)

func TestIsASCII(t *testing.T) {
	tests := []struct {
		input  string
		output bool
	}{
		{"Hello", true},
		{"456", true},
		{"", true},
		{"こんにちは、世界！", false}, // Japanese characters
		{"\n\t\r", false},    // Special characters
		{"😊", false},         // Emoji
	}

	for _, tt := range tests {
		output := IsASCII(tt.input)
		if output != tt.output {
			t.Errorf("IsASCII(%q) = %v, expected %v", tt.input, output, tt.output)
		}
	}
}
