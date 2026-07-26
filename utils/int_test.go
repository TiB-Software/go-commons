package utils

import "testing"

func TestStringToInt64(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		expected  int64
		expectErr bool
	}{
		{"valid positive", "123", 123, false},
		{"valid negative", "-456", -456, false},
		{"zero", "0", 0, false},
		{"invalid string", "abc", 0, true},
		{"empty string", "", 0, true},
		{"float string", "12.34", 0, true},
		{"large number", "9223372036854775807", 9223372036854775807, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := StringToInt64(tt.input)
			if tt.expectErr {
				if err == nil {
					t.Errorf("Expected error for input %s", tt.input)
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error for input %s: %v", tt.input, err)
				}
				if result != tt.expected {
					t.Errorf("StringToInt64(%s) = %d, want %d", tt.input, result, tt.expected)
				}
			}
		})
	}
}
