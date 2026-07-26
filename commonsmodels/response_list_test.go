package commonsmodels

import (
	"encoding/json"
	"testing"
)

func TestResponseListStruct(t *testing.T) {
	response := ResponseList[string]{
		Items: []string{"item1", "item2", "item3"},
		Total: 3,
	}

	if len(response.Items) != 3 {
		t.Errorf("Expected 3 items, got %d", len(response.Items))
	}
	if response.Total != 3 {
		t.Errorf("Expected Total 3, got %d", response.Total)
	}
}

func TestResponseListWithDifferentTypes(t *testing.T) {
	t.Run("string type", func(t *testing.T) {
		response := ResponseList[string]{
			Items: []string{"a", "b"},
			Total: 2,
		}
		if len(response.Items) != 2 {
			t.Errorf("Expected 2 items, got %d", len(response.Items))
		}
	})

	t.Run("int type", func(t *testing.T) {
		response := ResponseList[int]{
			Items: []int{1, 2, 3, 4, 5},
			Total: 5,
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
		response := ResponseList[TestItem]{
			Items: []TestItem{
				{ID: 1, Name: "test1"},
				{ID: 2, Name: "test2"},
			},
			Total: 2,
		}
		if len(response.Items) != 2 {
			t.Errorf("Expected 2 items, got %d", len(response.Items))
		}
		if response.Items[0].Name != "test1" {
			t.Errorf("Expected first item name 'test1', got '%s'", response.Items[0].Name)
		}
	})
}

func TestResponseListJSONSerialization(t *testing.T) {
	response := ResponseList[string]{
		Items: []string{"item1", "item2"},
		Total: 2,
	}

	data, err := json.Marshal(response)
	if err != nil {
		t.Fatalf("Failed to marshal ResponseList: %v", err)
	}

	expected := `{"items":["item1","item2"],"total":2}`
	if string(data) != expected {
		t.Errorf("Expected JSON '%s', got '%s'", expected, string(data))
	}
}

func TestResponseListJSONDeserialization(t *testing.T) {
	jsonData := `{"items":["a","b","c"],"total":3}`

	var response ResponseList[string]
	err := json.Unmarshal([]byte(jsonData), &response)
	if err != nil {
		t.Fatalf("Failed to unmarshal ResponseList: %v", err)
	}

	if len(response.Items) != 3 {
		t.Errorf("Expected 3 items, got %d", len(response.Items))
	}
	if response.Total != 3 {
		t.Errorf("Expected Total 3, got %d", response.Total)
	}
}

func TestResponseListEmptyItems(t *testing.T) {
	response := ResponseList[string]{
		Items: []string{},
		Total: 0,
	}

	data, err := json.Marshal(response)
	if err != nil {
		t.Fatalf("Failed to marshal ResponseList: %v", err)
	}

	expected := `{"items":[],"total":0}`
	if string(data) != expected {
		t.Errorf("Expected JSON '%s', got '%s'", expected, string(data))
	}
}

func TestResponseListTotalMismatch(t *testing.T) {
	// Test that Total can differ from actual items count
	// (might represent total in DB vs current page)
	response := ResponseList[string]{
		Items: []string{"item1", "item2"},
		Total: 100, // Total in database
	}

	if len(response.Items) != 2 {
		t.Errorf("Expected 2 items in current page, got %d", len(response.Items))
	}
	if response.Total != 100 {
		t.Errorf("Expected Total 100, got %d", response.Total)
	}
}
