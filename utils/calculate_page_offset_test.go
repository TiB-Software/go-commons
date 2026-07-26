package utils

import "testing"

func TestCalculateOffset(t *testing.T) {
	tests := []struct {
		name     string
		page     int32
		limit    int32
		expected int32
	}{
		{"first page", 1, 10, 0},
		{"second page", 2, 10, 10},
		{"third page", 3, 10, 20},
		{"different limit", 2, 25, 25},
		{"page zero edge case", 0, 10, -10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CalculateOffset(tt.page, tt.limit)
			if result != tt.expected {
				t.Errorf("CalculateOffset(%d, %d) = %d, want %d", tt.page, tt.limit, result, tt.expected)
			}
		})
	}
}
