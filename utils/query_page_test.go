package utils

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGetQueryPage(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("valid page", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request, _ = http.NewRequest(http.MethodGet, "/?page=5", nil)

		page, apiErr := GetQueryPage(ctx)

		if apiErr != nil {
			t.Errorf("Expected no error, got %v", apiErr)
		}
		if page != 5 {
			t.Errorf("Expected page 5, got %d", page)
		}
	})

	t.Run("default page", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request, _ = http.NewRequest(http.MethodGet, "/", nil)

		page, apiErr := GetQueryPage(ctx)

		if apiErr != nil {
			t.Errorf("Expected no error, got %v", apiErr)
		}
		if page != 1 {
			t.Errorf("Expected page 1, got %d", page)
		}
	})

	t.Run("page zero becomes one", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request, _ = http.NewRequest(http.MethodGet, "/?page=0", nil)

		page, apiErr := GetQueryPage(ctx)

		if apiErr != nil {
			t.Errorf("Expected no error, got %v", apiErr)
		}
		if page != 1 {
			t.Errorf("Expected page 1 for zero input, got %d", page)
		}
	})

	t.Run("invalid page", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request, _ = http.NewRequest(http.MethodGet, "/?page=invalid", nil)

		_, apiErr := GetQueryPage(ctx)

		if apiErr == nil {
			t.Error("Expected error for invalid page")
		}
		if apiErr.GetStatus() != http.StatusBadRequest {
			t.Errorf("Expected status %d, got %d", http.StatusBadRequest, apiErr.GetStatus())
		}
	})
}
