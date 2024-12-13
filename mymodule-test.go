package mymodule

import "testing"

func TestGreet(t *testing.T) {
	expected := "Hello, John!"
	result := Greet("John")
	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}
}
