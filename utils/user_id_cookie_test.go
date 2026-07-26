package utils

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/TB-Systems/go-commons/constants"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func TestGetUserIDFromContext(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("valid user id", func(t *testing.T) {
		userID := uuid.New()
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Set(constants.UserID, userID)

		result, apiErr := GetUserIDFromContext(ctx)

		if apiErr != nil {
			t.Errorf("Expected no error, got %v", apiErr)
		}
		if result != userID {
			t.Errorf("Expected %v, got %v", userID, result)
		}
	})

	t.Run("user id not set", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)

		_, apiErr := GetUserIDFromContext(ctx)

		if apiErr == nil {
			t.Error("Expected error when user ID not set")
		}
		if apiErr.GetStatus() != http.StatusUnauthorized {
			t.Errorf("Expected status %d, got %d", http.StatusUnauthorized, apiErr.GetStatus())
		}
	})

	t.Run("user id wrong type", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Set(constants.UserID, "not-a-uuid")

		_, apiErr := GetUserIDFromContext(ctx)

		if apiErr == nil {
			t.Error("Expected error for wrong type")
		}
		if apiErr.GetStatus() != http.StatusUnauthorized {
			t.Errorf("Expected status %d, got %d", http.StatusUnauthorized, apiErr.GetStatus())
		}
	})
}
