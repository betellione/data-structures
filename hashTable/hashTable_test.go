package hashTable

import "testing"

func TestHashTable(t *testing.T) {
	h := NewHashTable(10)

	h.Set("key1", "value1")
	h.Set("key2", "value2")

	if v, ok := h.Get("key1"); !ok || v != "value1" {
		t.Errorf("Expected 'value1', got '%s'", v)
	}

	if v, ok := h.Get("key2"); !ok || v != "value2" {
		t.Errorf("Expected 'value2', got '%s'", v)
	}

	if _, ok := h.Get("nonexistent"); ok {
		t.Errorf("Expected 'false', got 'true'")
	}

	h.Set("key1", "newValue")
	if v, ok := h.Get("key1"); !ok || v != "newValue" {
		t.Errorf("Expected 'newValue', got '%s'", v)
	}
}
