package utils

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/TB-Systems/go-commons/constants"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func init() {
	gin.SetMode(gin.TestMode)
}

func TestIDFromURLParam(t *testing.T) {
	t.Run("valid uuid", func(t *testing.T) {
		validUUID := uuid.New()
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Params = gin.Params{{Key: constants.ID, Value: validUUID.String()}}

		id, apiErr := IDFromURLParam(ctx)

		if apiErr != nil {
			t.Errorf("Expected no error, got %v", apiErr)
		}
		if id != validUUID {
			t.Errorf("Expected %v, got %v", validUUID, id)
		}
	})

	t.Run("invalid uuid", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Params = gin.Params{{Key: constants.ID, Value: "invalid-uuid"}}

		_, apiErr := IDFromURLParam(ctx)

		if apiErr == nil {
			t.Error("Expected error for invalid UUID")
		}
		if apiErr.GetStatus() != http.StatusBadRequest {
			t.Errorf("Expected status %d, got %d", http.StatusBadRequest, apiErr.GetStatus())
		}
	})

	t.Run("empty uuid", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Params = gin.Params{{Key: constants.ID, Value: ""}}

		_, apiErr := IDFromURLParam(ctx)

		if apiErr == nil {
			t.Error("Expected error for empty UUID")
		}
	})
}
