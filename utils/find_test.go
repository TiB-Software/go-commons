package utils

import "testing"

func TestFindIf(t *testing.T) {
	t.Run("find existing element", func(t *testing.T) {
		slice := []int{1, 2, 3, 4, 5}
		found := FindIf(slice, func(i int) bool { return i == 3 })
		if !found {
			t.Error("Expected to find element 3")
		}
	})

	t.Run("element not found", func(t *testing.T) {
		slice := []int{1, 2, 3, 4, 5}
		found := FindIf(slice, func(i int) bool { return i == 10 })
		if found {
			t.Error("Expected not to find element 10")
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		slice := []int{}
		found := FindIf(slice, func(i int) bool { return i == 1 })
		if found {
			t.Error("Expected not to find in empty slice")
		}
	})

	t.Run("find with string slice", func(t *testing.T) {
		slice := []string{"apple", "banana", "cherry"}
		found := FindIf(slice, func(s string) bool { return s == "banana" })
		if !found {
			t.Error("Expected to find banana")
		}
	})
}

func TestFindIndex(t *testing.T) {
	t.Run("find existing element", func(t *testing.T) {
		slice := []int{1, 2, 3, 4, 5}
		index := FindIndex(slice, 3)
		if index != 2 {
			t.Errorf("Expected index 2, got %d", index)
		}
	})

	t.Run("element not found", func(t *testing.T) {
		slice := []int{1, 2, 3, 4, 5}
		index := FindIndex(slice, 10)
		if index != -1 {
			t.Errorf("Expected index -1, got %d", index)
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		slice := []int{}
		index := FindIndex(slice, 1)
		if index != -1 {
			t.Errorf("Expected index -1, got %d", index)
		}
	})

	t.Run("first element", func(t *testing.T) {
		slice := []string{"a", "b", "c"}
		index := FindIndex(slice, "a")
		if index != 0 {
			t.Errorf("Expected index 0, got %d", index)
		}
	})

	t.Run("last element", func(t *testing.T) {
		slice := []string{"a", "b", "c"}
		index := FindIndex(slice, "c")
		if index != 2 {
			t.Errorf("Expected index 2, got %d", index)
		}
	})
}
