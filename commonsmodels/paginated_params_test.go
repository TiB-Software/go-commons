package commonsmodels

import (
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestPaginatedParamsStruct(t *testing.T) {
	userID := uuid.New()
	params := PaginatedParams{
		UserID: userID,
		Limit:  10,
		Offset: 20,
		Page:   3,
	}

	if params.UserID != userID {
		t.Errorf("Expected UserID %v, got %v", userID, params.UserID)
	}
	if params.Limit != 10 {
		t.Errorf("Expected Limit 10, got %d", params.Limit)
	}
	if params.Offset != 20 {
		t.Errorf("Expected Offset 20, got %d", params.Offset)
	}
	if params.Page != 3 {
		t.Errorf("Expected Page 3, got %d", params.Page)
	}
}

func TestPaginatedParamsWithDateRangeStruct(t *testing.T) {
	userID := uuid.New()
	startDate := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2024, 12, 31, 23, 59, 59, 0, time.UTC)

	params := PaginatedParamsWithDateRange{
		UserID:    userID,
		Limit:     25,
		Offset:    50,
		Page:      3,
		StartDate: startDate,
		EndDate:   endDate,
	}

	if params.UserID != userID {
		t.Errorf("Expected UserID %v, got %v", userID, params.UserID)
	}
	if params.Limit != 25 {
		t.Errorf("Expected Limit 25, got %d", params.Limit)
	}
	if params.Offset != 50 {
		t.Errorf("Expected Offset 50, got %d", params.Offset)
	}
	if params.Page != 3 {
		t.Errorf("Expected Page 3, got %d", params.Page)
	}
	if !params.StartDate.Equal(startDate) {
		t.Errorf("Expected StartDate %v, got %v", startDate, params.StartDate)
	}
	if !params.EndDate.Equal(endDate) {
		t.Errorf("Expected EndDate %v, got %v", endDate, params.EndDate)
	}
}

func TestPaginatedParamsZeroValues(t *testing.T) {
	params := PaginatedParams{}

	if params.UserID != uuid.Nil {
		t.Errorf("Expected zero UUID, got %v", params.UserID)
	}
	if params.Limit != 0 {
		t.Errorf("Expected Limit 0, got %d", params.Limit)
	}
	if params.Offset != 0 {
		t.Errorf("Expected Offset 0, got %d", params.Offset)
	}
	if params.Page != 0 {
		t.Errorf("Expected Page 0, got %d", params.Page)
	}
}

func TestPaginatedParamsWithDateRangeZeroValues(t *testing.T) {
	params := PaginatedParamsWithDateRange{}

	if params.UserID != uuid.Nil {
		t.Errorf("Expected zero UUID, got %v", params.UserID)
	}
	if !params.StartDate.IsZero() {
		t.Errorf("Expected zero StartDate, got %v", params.StartDate)
	}
	if !params.EndDate.IsZero() {
		t.Errorf("Expected zero EndDate, got %v", params.EndDate)
	}
}
