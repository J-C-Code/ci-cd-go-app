package main

import "testing"

func TestAdd(t *testing.T) {
	result := sum(1, 5)
	if result != 6 {
		t.Errorf("add(1, 5) =%d; want 6", result)
	}
}
