package singleton

import "testing"

func TestSingleton(t *testing.T) {
	s := New()
	s["this"] = "that"
	s2 := New()

	if s2["this"] != "that" {
		t.Errorf("This should be that")
	}
}