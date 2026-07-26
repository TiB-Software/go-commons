package utils

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/TB-Systems/go-commons/errors"

	"github.com/gin-gonic/gin"
)

func TestSendResponse(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("send json response", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)

		data := map[string]string{"message": "hello"}
		SendResponse(ctx, data, http.StatusOK)

		if w.Code != http.StatusOK {
			t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
		}

		var response map[string]string
		json.Unmarshal(w.Body.Bytes(), &response)
		if response["message"] != "hello" {
			t.Errorf("Expected message 'hello', got '%s'", response["message"])
		}
	})

	t.Run("send with created status", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)

		data := map[string]int{"id": 1}
		SendResponse(ctx, data, http.StatusCreated)

		if w.Code != http.StatusCreated {
			t.Errorf("Expected status %d, got %d", http.StatusCreated, w.Code)
		}
	})
}

func TestSendErrorResponse(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("send error response", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)

		apiErr := errors.NewApiError(http.StatusBadRequest, errors.BadRequestError("test error"))
		SendErrorResponse(ctx, apiErr)

		if w.Code != http.StatusBadRequest {
			t.Errorf("Expected status %d, got %d", http.StatusBadRequest, w.Code)
		}

		var response errors.ErrorResponse
		json.Unmarshal(w.Body.Bytes(), &response)
		if response.Status != http.StatusBadRequest {
			t.Errorf("Expected status in body %d, got %d", http.StatusBadRequest, response.Status)
		}
		if len(response.Messages) == 0 {
			t.Error("Expected at least one message")
		}
	})

	t.Run("send multiple error messages", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)

		apiErr := errors.NewApiErrorWithErrors(http.StatusBadRequest, []errors.ApiErrorItem{
			errors.BadRequestError("error 1"),
			errors.BadRequestError("error 2"),
		})
		SendErrorResponse(ctx, apiErr)

		var response errors.ErrorResponse
		json.Unmarshal(w.Body.Bytes(), &response)
		if len(response.Messages) != 2 {
			t.Errorf("Expected 2 messages, got %d", len(response.Messages))
		}
	})
}
