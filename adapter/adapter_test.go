package adapter

import "testing"

func TestAdapter(t *testing.T) {
	adapter := NewAdapter()

	t.Run("GetUsername with valid ID", func(t *testing.T) {
		username := adapter.GetUsername("1")
		if username != "Alice" {
			t.Errorf("Expected 'Alice' but got '%s'", username)
		}
	})

	t.Run("GetUsername with invalid ID", func(t *testing.T) {
		username := adapter.GetUsername("999")
		if username != "Unknown User" {
			t.Errorf("Expected 'Unknown User' but got '%s'", username)
		}
	})
}
