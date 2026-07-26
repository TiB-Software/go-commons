package commonsmodels

import (
	"encoding/json"
	"testing"
)

func TestNewResponseSuccess(t *testing.T) {
	response := NewResponseSuccess()

	if response.Message != "success" {
		t.Errorf("Expected message '%s', got '%s'", "success", response.Message)
	}
}

func TestResponseSuccessJSONSerialization(t *testing.T) {
	response := NewResponseSuccess()

	data, err := json.Marshal(response)
	if err != nil {
		t.Fatalf("Failed to marshal ResponseSuccess: %v", err)
	}

	expected := `{"message":"success"}`
	if string(data) != expected {
		t.Errorf("Expected JSON '%s', got '%s'", expected, string(data))
	}
}

func TestResponseSuccessJSONDeserialization(t *testing.T) {
	jsonData := `{"message":"success"}`

	var response ResponseSuccess
	err := json.Unmarshal([]byte(jsonData), &response)
	if err != nil {
		t.Fatalf("Failed to unmarshal ResponseSuccess: %v", err)
	}

	if response.Message != "success" {
		t.Errorf("Expected message 'success', got '%s'", response.Message)
	}
}
