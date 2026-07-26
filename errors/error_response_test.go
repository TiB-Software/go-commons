package errors

import (
	"encoding/json"
	"testing"
)

func TestErrorResponseStruct(t *testing.T) {
	response := ErrorResponse{
		Status:   400,
		Messages: []string{"error 1", "error 2"},
	}

	if response.Status != 400 {
		t.Errorf("Expected Status 400, got %d", response.Status)
	}

	if len(response.Messages) != 2 {
		t.Errorf("Expected 2 messages, got %d", len(response.Messages))
	}

	if response.Messages[0] != "error 1" {
		t.Errorf("Expected first message 'error 1', got '%s'", response.Messages[0])
	}

	if response.Messages[1] != "error 2" {
		t.Errorf("Expected second message 'error 2', got '%s'", response.Messages[1])
	}
}

func TestErrorResponseJSONSerialization(t *testing.T) {
	response := ErrorResponse{
		Status:   404,
		Messages: []string{"not found"},
	}

	data, err := json.Marshal(response)
	if err != nil {
		t.Fatalf("Failed to marshal ErrorResponse: %v", err)
	}

	expected := `{"status":404,"messages":["not found"]}`
	if string(data) != expected {
		t.Errorf("Expected JSON '%s', got '%s'", expected, string(data))
	}
}

func TestErrorResponseJSONDeserialization(t *testing.T) {
	jsonData := `{"status":500,"messages":["internal error","database error"]}`

	var response ErrorResponse
	err := json.Unmarshal([]byte(jsonData), &response)
	if err != nil {
		t.Fatalf("Failed to unmarshal ErrorResponse: %v", err)
	}

	if response.Status != 500 {
		t.Errorf("Expected Status 500, got %d", response.Status)
	}

	if len(response.Messages) != 2 {
		t.Errorf("Expected 2 messages, got %d", len(response.Messages))
	}

	if response.Messages[0] != "internal error" {
		t.Errorf("Expected first message 'internal error', got '%s'", response.Messages[0])
	}

	if response.Messages[1] != "database error" {
		t.Errorf("Expected second message 'database error', got '%s'", response.Messages[1])
	}
}

func TestErrorResponseEmptyMessages(t *testing.T) {
	response := ErrorResponse{
		Status:   400,
		Messages: []string{},
	}

	data, err := json.Marshal(response)
	if err != nil {
		t.Fatalf("Failed to marshal ErrorResponse: %v", err)
	}

	expected := `{"status":400,"messages":[]}`
	if string(data) != expected {
		t.Errorf("Expected JSON '%s', got '%s'", expected, string(data))
	}
}

func TestErrorResponseJSONTags(t *testing.T) {
	jsonData := `{"status":401,"messages":["unauthorized"]}`

	var response ErrorResponse
	err := json.Unmarshal([]byte(jsonData), &response)
	if err != nil {
		t.Fatalf("Failed to unmarshal ErrorResponse: %v", err)
	}

	if response.Status != 401 {
		t.Errorf("Expected Status 401, got %d", response.Status)
	}

	if len(response.Messages) != 1 || response.Messages[0] != "unauthorized" {
		t.Errorf("Expected messages ['unauthorized'], got %v", response.Messages)
	}
}
