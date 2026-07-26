package utils

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/TB-Systems/go-commons/errors"
	"github.com/TB-Systems/go-commons/validator"

	"github.com/gin-gonic/gin"
)

type testRequest struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

func (r testRequest) Validate() []errors.ApiErrorItem {
	var errs []errors.ApiErrorItem
	if r.Name == "" {
		errs = append(errs, errors.BadRequestError("name is required"))
	}
	return errs
}

var _ validator.Validator = testRequest{}

func TestDecodeJson(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("valid json", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)

		body := `{"name": "test", "value": 42}`
		ctx.Request, _ = http.NewRequest(http.MethodPost, "/", bytes.NewBufferString(body))
		ctx.Request.Header.Set("Content-Type", "application/json")

		data, errs := DecodeJson[testRequest](ctx)

		if len(errs) > 0 {
			t.Errorf("Expected no errors, got %v", errs)
		}
		if data.Name != "test" || data.Value != 42 {
			t.Errorf("Expected {test, 42}, got {%s, %d}", data.Name, data.Value)
		}
	})

	t.Run("invalid json", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)

		ctx.Request, _ = http.NewRequest(http.MethodPost, "/", bytes.NewBufferString("invalid json"))
		ctx.Request.Header.Set("Content-Type", "application/json")

		_, errs := DecodeJson[testRequest](ctx)

		if len(errs) == 0 {
			t.Error("Expected errors for invalid JSON")
		}
	})
}

func TestDecodeValidJson(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("valid json with valid data", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)

		body := `{"name": "test", "value": 42}`
		ctx.Request, _ = http.NewRequest(http.MethodPost, "/", bytes.NewBufferString(body))
		ctx.Request.Header.Set("Content-Type", "application/json")

		data, apiErr := DecodeValidJson[testRequest](ctx)

		if apiErr != nil {
			t.Errorf("Expected no error, got %v", apiErr)
		}
		if data.Name != "test" {
			t.Errorf("Expected name 'test', got '%s'", data.Name)
		}
	})

	t.Run("valid json with invalid data", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)

		body := `{"name": "", "value": 42}`
		ctx.Request, _ = http.NewRequest(http.MethodPost, "/", bytes.NewBufferString(body))
		ctx.Request.Header.Set("Content-Type", "application/json")

		_, apiErr := DecodeValidJson[testRequest](ctx)

		if apiErr == nil {
			t.Error("Expected validation error")
		}
		if apiErr.GetStatus() != http.StatusUnprocessableEntity {
			t.Errorf("Expected status %d, got %d", http.StatusUnprocessableEntity, apiErr.GetStatus())
		}
	})

	t.Run("invalid json", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)

		ctx.Request, _ = http.NewRequest(http.MethodPost, "/", bytes.NewBufferString("invalid"))
		ctx.Request.Header.Set("Content-Type", "application/json")

		_, apiErr := DecodeValidJson[testRequest](ctx)

		if apiErr == nil {
			t.Error("Expected decode error")
		}
		if apiErr.GetStatus() != http.StatusBadRequest {
			t.Errorf("Expected status %d, got %d", http.StatusBadRequest, apiErr.GetStatus())
		}
	})
}
