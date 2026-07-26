package utils

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
)

func TestGetQueryDatesIfHas(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("no dates provided", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request, _ = http.NewRequest(http.MethodGet, "/", nil)

		_, _, hasDates, apiErr := GetQueryDatesIfHas(ctx)

		if apiErr != nil {
			t.Errorf("Expected no error, got %v", apiErr)
		}
		if hasDates {
			t.Error("Expected hasDates to be false")
		}
	})

	t.Run("valid dates", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request, _ = http.NewRequest(http.MethodGet, "/?start_date=2024-01-01&end_date=2024-12-31", nil)

		startDate, endDate, hasDates, apiErr := GetQueryDatesIfHas(ctx)

		if apiErr != nil {
			t.Errorf("Expected no error, got %v", apiErr)
		}
		if !hasDates {
			t.Error("Expected hasDates to be true")
		}
		expectedStart, _ := time.Parse(time.DateOnly, "2024-01-01")
		expectedEnd, _ := time.Parse(time.DateOnly, "2024-12-31")
		if !startDate.Equal(expectedStart) {
			t.Errorf("Expected start date %v, got %v", expectedStart, startDate)
		}
		if !endDate.Equal(expectedEnd) {
			t.Errorf("Expected end date %v, got %v", expectedEnd, endDate)
		}
	})

	t.Run("invalid start date", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request, _ = http.NewRequest(http.MethodGet, "/?start_date=invalid&end_date=2024-12-31", nil)

		_, _, _, apiErr := GetQueryDatesIfHas(ctx)

		if apiErr == nil {
			t.Error("Expected error for invalid start date")
		}
		if apiErr.GetStatus() != http.StatusBadRequest {
			t.Errorf("Expected status %d, got %d", http.StatusBadRequest, apiErr.GetStatus())
		}
	})

	t.Run("invalid end date", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request, _ = http.NewRequest(http.MethodGet, "/?start_date=2024-01-01&end_date=invalid", nil)

		_, _, _, apiErr := GetQueryDatesIfHas(ctx)

		if apiErr == nil {
			t.Error("Expected error for invalid end date")
		}
		if apiErr.GetStatus() != http.StatusBadRequest {
			t.Errorf("Expected status %d, got %d", http.StatusBadRequest, apiErr.GetStatus())
		}
	})
}
