package generic

import (
	"testing"
)

func TestMapItem(t *testing.T) {
	m := map[string]any{
		"key1": "value1",
		"key2": 42,
		"props": map[string]int{
			"nestedKey": 100,
		},
	}

	if val := MapItem[string](m, "key1"); val != "value1" {
		t.Errorf("Expected 'value1', got '%s'", val)
	}

	if val := MapItem[int](m, "key2"); val != 42 {
		t.Errorf("Expected 42, got %d", val)
	}

	if val := MapItem[string](m, "nonexistent"); val != "" {
		t.Errorf("Expected empty string, got '%s'", val)
	}

	if val := MapItem[int](m, "key1"); val != 0 {
		t.Errorf("Expected 0 for type mismatch, got %d", val)
	}

	if val := MapItem[map[string]int](m, "props"); val["nestedKey"] != 100 {
		t.Errorf("Expected nestedKey to be 100, got %d", val["nestedKey"])
	}
}
