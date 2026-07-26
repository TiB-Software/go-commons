package errors

import (
	"testing"
)

func TestNewApiError(t *testing.T) {
	message := ApiErrorItem{
		UserMessage:   "user message",
		SystemMessage: "system message",
		SystemDetail:  "detail",
	}

	err := NewApiError(400, message)

	if err.GetStatus() != 400 {
		t.Errorf("Expected status 400, got %d", err.GetStatus())
	}

	messages := err.GetMessages()
	if len(messages) != 1 {
		t.Errorf("Expected 1 message, got %d", len(messages))
	}

	if messages[0].UserMessage != "user message" {
		t.Errorf("Expected UserMessage 'user message', got '%s'", messages[0].UserMessage)
	}
	if messages[0].SystemMessage != "system message" {
		t.Errorf("Expected SystemMessage 'system message', got '%s'", messages[0].SystemMessage)
	}
	if messages[0].SystemDetail != "detail" {
		t.Errorf("Expected SystemDetail 'detail', got '%s'", messages[0].SystemDetail)
	}
}

func TestNewApiErrorWithErrors(t *testing.T) {
	messages := []ApiErrorItem{
		{
			UserMessage:   "error 1",
			SystemMessage: "sys error 1",
			SystemDetail:  "detail 1",
		},
		{
			UserMessage:   "error 2",
			SystemMessage: "sys error 2",
			SystemDetail:  "detail 2",
		},
	}

	err := NewApiErrorWithErrors(422, messages)

	if err.GetStatus() != 422 {
		t.Errorf("Expected status 422, got %d", err.GetStatus())
	}

	result := err.GetMessages()
	if len(result) != 2 {
		t.Errorf("Expected 2 messages, got %d", len(result))
	}

	if result[0].UserMessage != "error 1" {
		t.Errorf("Expected first UserMessage 'error 1', got '%s'", result[0].UserMessage)
	}
	if result[1].UserMessage != "error 2" {
		t.Errorf("Expected second UserMessage 'error 2', got '%s'", result[1].UserMessage)
	}
}

func TestApiErrorGetStatus(t *testing.T) {
	tests := []struct {
		name           string
		status         int
		expectedStatus int
	}{
		{"Bad Request", 400, 400},
		{"Unauthorized", 401, 401},
		{"Not Found", 404, 404},
		{"Internal Server Error", 500, 500},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := NewApiError(tc.status, ApiErrorItem{})
			if err.GetStatus() != tc.expectedStatus {
				t.Errorf("Expected status %d, got %d", tc.expectedStatus, err.GetStatus())
			}
		})
	}
}

func TestApiErrorGetMessages(t *testing.T) {
	t.Run("single message", func(t *testing.T) {
		message := ApiErrorItem{UserMessage: "test"}
		err := NewApiError(400, message)

		messages := err.GetMessages()
		if len(messages) != 1 {
			t.Errorf("Expected 1 message, got %d", len(messages))
		}
	})

	t.Run("multiple messages", func(t *testing.T) {
		messages := []ApiErrorItem{
			{UserMessage: "test1"},
			{UserMessage: "test2"},
			{UserMessage: "test3"},
		}
		err := NewApiErrorWithErrors(400, messages)

		result := err.GetMessages()
		if len(result) != 3 {
			t.Errorf("Expected 3 messages, got %d", len(result))
		}
	})

	t.Run("empty messages", func(t *testing.T) {
		err := NewApiErrorWithErrors(400, []ApiErrorItem{})

		messages := err.GetMessages()
		if len(messages) != 0 {
			t.Errorf("Expected 0 messages, got %d", len(messages))
		}
	})
}
