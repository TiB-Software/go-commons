package utils

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/TB-Systems/go-commons/constants"

	"github.com/gin-gonic/gin"
)

func TestGetQueryLimit(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("valid limit within max", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request, _ = http.NewRequest(http.MethodGet, "/?limit=5", nil)

		limit := GetQueryLimit(ctx)

		if limit != 5 {
			t.Errorf("Expected limit 5, got %d", limit)
		}
	})

	t.Run("default limit", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request, _ = http.NewRequest(http.MethodGet, "/", nil)

		limit := GetQueryLimit(ctx)

		if limit != int32(constants.LimitDefault) {
			t.Errorf("Expected default limit %d, got %d", constants.LimitDefault, limit)
		}
	})

	t.Run("limit exceeds max", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request, _ = http.NewRequest(http.MethodGet, "/?limit=999999", nil)

		limit := GetQueryLimit(ctx)

		if limit != int32(constants.LimitDefault) {
			t.Errorf("Expected default limit for exceeded max, got %d", limit)
		}
	})

	t.Run("invalid limit falls back to default", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request, _ = http.NewRequest(http.MethodGet, "/?limit=invalid", nil)

		limit := GetQueryLimit(ctx)

		if limit != int32(constants.LimitDefault) {
			t.Errorf("Expected default limit for invalid input, got %d", limit)
		}
	})
}
