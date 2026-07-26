package utils

import "testing"

func TestIsBlank(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"empty string", "", true},
		{"only spaces", "   ", true},
		{"only tabs", "\t\t", true},
		{"spaces and tabs", " \t ", true},
		{"non-blank", "hello", false},
		{"spaces around text", "  hello  ", false},
		{"newline only", "\n", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsBlank(tt.input)
			if result != tt.expected {
				t.Errorf("IsBlank(%q) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}
