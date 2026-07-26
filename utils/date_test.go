package utils

import (
	"testing"
	"time"
)

func TestNormalizeDay(t *testing.T) {
	tests := []struct {
		name     string
		year     int
		month    time.Month
		day      int
		expected int
	}{
		{"normal day in january", 2024, time.January, 15, 15},
		{"day 31 in february", 2024, time.February, 31, 29}, // 2024 is leap year
		{"day 30 in february", 2024, time.February, 30, 29},
		{"day 31 in april", 2024, time.April, 31, 30},       // April has 30 days
		{"day 31 in december", 2024, time.December, 31, 31}, // December has 31 days
		{"day 1", 2024, time.January, 1, 1},
		{"day 0 becomes 1", 2024, time.January, 0, 1},
		{"negative day becomes 1", 2024, time.January, -5, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NormalizeDay(tt.year, tt.month, tt.day)
			if result != tt.expected {
				t.Errorf("NormalizeDay(%d, %v, %d) = %d, want %d", tt.year, tt.month, tt.day, result, tt.expected)
			}
		})
	}
}

func TestCreateDateWithNormalizedDay(t *testing.T) {
	t.Run("creates date with normalized day for february", func(t *testing.T) {
		result := CreateDateWithNormalizedDay(2024, time.February, 31)
		expected := time.Date(2024, time.February, 29, 0, 0, 0, 0, time.UTC) // 2024 is leap year
		if !result.Equal(expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("creates date with normal day", func(t *testing.T) {
		result := CreateDateWithNormalizedDay(2024, time.June, 15)
		expected := time.Date(2024, time.June, 15, 0, 0, 0, 0, time.UTC)
		if !result.Equal(expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("january 31 stays 31", func(t *testing.T) {
		result := CreateDateWithNormalizedDay(2024, time.January, 31)
		expected := time.Date(2024, time.January, 31, 0, 0, 0, 0, time.UTC)
		if !result.Equal(expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})
}
