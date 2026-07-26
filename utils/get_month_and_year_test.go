package utils

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/TB-Systems/go-commons/constants"

	"github.com/gin-gonic/gin"
)

func TestGetQueryMonthAndYear(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("success", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request, _ = http.NewRequest(http.MethodGet, "/?month=9&year=2025", nil)

		month, year, apiErr := GetQueryMonthAndYear(ctx)

		if apiErr != nil {
			t.Fatalf("Expected no error, got %v", apiErr)
		}
		if month != 9 || year != 2025 {
			t.Errorf("Expected month=9 and year=2025, got month=%d year=%d", month, year)
		}
	})

	t.Run("invalid month parse", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request, _ = http.NewRequest(http.MethodGet, "/?month=abc&year=2025", nil)

		_, _, apiErr := GetQueryMonthAndYear(ctx)

		if apiErr == nil {
			t.Fatal("Expected error, got nil")
		}
		if apiErr.GetStatus() != http.StatusBadRequest {
			t.Errorf("Expected status %d, got %d", http.StatusBadRequest, apiErr.GetStatus())
		}
		if len(apiErr.GetMessages()) == 0 || apiErr.GetMessages()[0].UserMessage != constants.MonthInvalidMsg {
			t.Errorf("Expected month invalid message, got %+v", apiErr.GetMessages())
		}
	})

	t.Run("invalid year parse", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request, _ = http.NewRequest(http.MethodGet, "/?month=9&year=abc", nil)

		_, _, apiErr := GetQueryMonthAndYear(ctx)

		if apiErr == nil {
			t.Fatal("Expected error, got nil")
		}
		if apiErr.GetStatus() != http.StatusBadRequest {
			t.Errorf("Expected status %d, got %d", http.StatusBadRequest, apiErr.GetStatus())
		}
		if len(apiErr.GetMessages()) == 0 || apiErr.GetMessages()[0].UserMessage != constants.YearInvalidMsg {
			t.Errorf("Expected year invalid message, got %+v", apiErr.GetMessages())
		}
	})

	t.Run("month lower than one", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request, _ = http.NewRequest(http.MethodGet, "/?month=0&year=2025", nil)

		_, _, apiErr := GetQueryMonthAndYear(ctx)

		if apiErr == nil {
			t.Fatal("Expected error, got nil")
		}
		if apiErr.GetStatus() != http.StatusBadRequest {
			t.Errorf("Expected status %d, got %d", http.StatusBadRequest, apiErr.GetStatus())
		}
		if len(apiErr.GetMessages()) == 0 || apiErr.GetMessages()[0].UserMessage != constants.MonthInvalidMsg {
			t.Errorf("Expected month invalid message, got %+v", apiErr.GetMessages())
		}
	})

	t.Run("month greater than twelve", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request, _ = http.NewRequest(http.MethodGet, "/?month=13&year=2025", nil)

		_, _, apiErr := GetQueryMonthAndYear(ctx)

		if apiErr == nil {
			t.Fatal("Expected error, got nil")
		}
		if apiErr.GetStatus() != http.StatusBadRequest {
			t.Errorf("Expected status %d, got %d", http.StatusBadRequest, apiErr.GetStatus())
		}
		if len(apiErr.GetMessages()) == 0 || apiErr.GetMessages()[0].UserMessage != constants.MonthInvalidMsg {
			t.Errorf("Expected month invalid message, got %+v", apiErr.GetMessages())
		}
	})

	t.Run("year lower than 1970", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request, _ = http.NewRequest(http.MethodGet, "/?month=9&year=1969", nil)

		_, _, apiErr := GetQueryMonthAndYear(ctx)

		if apiErr == nil {
			t.Fatal("Expected error, got nil")
		}
		if apiErr.GetStatus() != http.StatusBadRequest {
			t.Errorf("Expected status %d, got %d", http.StatusBadRequest, apiErr.GetStatus())
		}
		if len(apiErr.GetMessages()) == 0 || apiErr.GetMessages()[0].UserMessage != constants.YearMustBe1970OrLaterMsg {
			t.Errorf("Expected year lower bound message, got %+v", apiErr.GetMessages())
		}
	})
}
