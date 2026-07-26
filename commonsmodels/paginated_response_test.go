package commonsmodels

import (
	"encoding/json"
	"testing"
)

func TestPaginatedResponseStruct(t *testing.T) {
	response := PaginatedResponse[string]{
		Items:     []string{"item1", "item2", "item3"},
		PageCount: 5,
		Page:      2,
	}

	if len(response.Items) != 3 {
		t.Errorf("Expected 3 items, got %d", len(response.Items))
	}
	if response.PageCount != 5 {
		t.Errorf("Expected PageCount 5, got %d", response.PageCount)
	}
	if response.Page != 2 {
		t.Errorf("Expected Page 2, got %d", response.Page)
	}
}

func TestPaginatedResponseWithDifferentTypes(t *testing.T) {
	t.Run("string type", func(t *testing.T) {
		response := PaginatedResponse[string]{
			Items:     []string{"a", "b"},
			PageCount: 1,
			Page:      1,
		}
		if len(response.Items) != 2 {
			t.Errorf("Expected 2 items, got %d", len(response.Items))
		}
	})

	t.Run("int type", func(t *testing.T) {
		response := PaginatedResponse[int]{
			Items:     []int{1, 2, 3, 4, 5},
			PageCount: 10,
			Page:      1,
		}
		if len(response.Items) != 5 {
			t.Errorf("Expected 5 items, got %d", len(response.Items))
		}
	})

	t.Run("struct type", func(t *testing.T) {
		type TestItem struct {
			ID   int
			Name string
		}
		response := PaginatedResponse[TestItem]{
			Items: []TestItem{
				{ID: 1, Name: "test1"},
				{ID: 2, Name: "test2"},
			},
			PageCount: 3,
			Page:      1,
		}
		if len(response.Items) != 2 {
			t.Errorf("Expected 2 items, got %d", len(response.Items))
		}
		if response.Items[0].Name != "test1" {
			t.Errorf("Expected first item name 'test1', got '%s'", response.Items[0].Name)
		}
	})
}

func TestPaginatedResponseJSONSerialization(t *testing.T) {
	response := PaginatedResponse[string]{
		Items:     []string{"item1", "item2"},
		PageCount: 5,
		Page:      2,
	}

	data, err := json.Marshal(response)
	if err != nil {
		t.Fatalf("Failed to marshal PaginatedResponse: %v", err)
	}

	expected := `{"items":["item1","item2"],"page_count":5,"page":2}`
	if string(data) != expected {
		t.Errorf("Expected JSON '%s', got '%s'", expected, string(data))
	}
}

func TestPaginatedResponseJSONDeserialization(t *testing.T) {
	jsonData := `{"items":["a","b","c"],"page_count":10,"page":3}`

	var response PaginatedResponse[string]
	err := json.Unmarshal([]byte(jsonData), &response)
	if err != nil {
		t.Fatalf("Failed to unmarshal PaginatedResponse: %v", err)
	}

	if len(response.Items) != 3 {
		t.Errorf("Expected 3 items, got %d", len(response.Items))
	}
	if response.PageCount != 10 {
		t.Errorf("Expected PageCount 10, got %d", response.PageCount)
	}
	if response.Page != 3 {
		t.Errorf("Expected Page 3, got %d", response.Page)
	}
}

func TestPaginatedResponseEmptyItems(t *testing.T) {
	response := PaginatedResponse[string]{
		Items:     []string{},
		PageCount: 0,
		Page:      1,
	}

	data, err := json.Marshal(response)
	if err != nil {
		t.Fatalf("Failed to marshal PaginatedResponse: %v", err)
	}

	expected := `{"items":[],"page_count":0,"page":1}`
	if string(data) != expected {
		t.Errorf("Expected JSON '%s', got '%s'", expected, string(data))
	}
}
