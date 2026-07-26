package utils

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

func TestPgTypeUUIDToUUID(t *testing.T) {
	t.Run("valid uuid", func(t *testing.T) {
		originalUUID := uuid.New()
		pgUUID := pgtype.UUID{Bytes: originalUUID, Valid: true}

		result := PgTypeUUIDToUUID(pgUUID)

		if result == nil {
			t.Error("Expected non-nil result")
		}
		if *result != originalUUID {
			t.Errorf("Expected %v, got %v", originalUUID, *result)
		}
	})

	t.Run("invalid pgtype uuid", func(t *testing.T) {
		pgUUID := pgtype.UUID{Valid: false}

		result := PgTypeUUIDToUUID(pgUUID)

		if result != nil {
			t.Error("Expected nil result for invalid pgtype UUID")
		}
	})
}

func TestUUIDToPgTypeUUID(t *testing.T) {
	t.Run("valid uuid", func(t *testing.T) {
		originalUUID := uuid.New()

		result := UUIDToPgTypeUUID(&originalUUID)

		if !result.Valid {
			t.Error("Expected valid result")
		}
		if result.Bytes != originalUUID {
			t.Errorf("Expected %v, got %v", originalUUID, result.Bytes)
		}
	})

	t.Run("nil uuid", func(t *testing.T) {
		result := UUIDToPgTypeUUID(nil)

		if result.Valid {
			t.Error("Expected invalid result for nil UUID")
		}
	})
}

func TestFloat64ToNumeric(t *testing.T) {
	tests := []struct {
		name     string
		input    float64
		expected float64
	}{
		{"positive float", 123.45, 123.45},
		{"negative float", -67.89, -67.89},
		{"zero", 0, 0},
		{"whole number", 100.00, 100.00},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			numeric := Float64ToNumeric(tt.input)
			if !numeric.Valid {
				t.Error("Expected valid numeric")
			}
			result := NumericToFloat64(numeric)
			if result != tt.expected {
				t.Errorf("Float64ToNumeric(%f) roundtrip = %f, want %f", tt.input, result, tt.expected)
			}
		})
	}
}

func TestNumericToFloat64(t *testing.T) {
	t.Run("valid numeric", func(t *testing.T) {
		numeric := Float64ToNumeric(123.45)
		result := NumericToFloat64(numeric)
		if result != 123.45 {
			t.Errorf("Expected 123.45, got %f", result)
		}
	})

	t.Run("invalid numeric returns zero", func(t *testing.T) {
		numeric := pgtype.Numeric{Valid: false}
		result := NumericToFloat64(numeric)
		if result != 0 {
			t.Errorf("Expected 0 for invalid numeric, got %f", result)
		}
	})
}

func TestTimeToPgTimestamptz(t *testing.T) {
	now := time.Now()
	result := TimeToPgTimestamptz(now)

	if !result.Valid {
		t.Error("Expected valid timestamptz")
	}
	if !result.Time.Equal(now) {
		t.Errorf("Expected %v, got %v", now, result.Time)
	}
}
