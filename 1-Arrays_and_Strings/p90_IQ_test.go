package main

import "testing"

func TestTableDriven(t *testing.T) {
	var tests = []struct {
		chain string
		expected bool
	}{
		{"abcde", true},
		{"azertyuia", false},
		{"1'(-1", false},
	}

	for _, test := range tests {
		if isUnique(test.chain) != test.expected {
			t.Errorf("Should return %t", test.expected)
		}
	}
}
