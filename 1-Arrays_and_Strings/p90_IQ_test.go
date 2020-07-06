package main

import "testing"

func TestIsUnique(t *testing.T) {
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

func TestCheckPermutation(t *testing.T) {
	var tests = []struct {
		chain1 string
		chain2 string
		expected bool
	} {
		{"a", "ba", false},
		{"kayak", "kayak", true},
		{"Thomas", "samohT", true},
		{"Thomas", "samoth", false},
	}

	for _, test := range tests {
		if checkPermutation(test.chain1, test.chain2) != test.expected {
			t.Errorf("Should return %t", test.expected)
		}
	}
}

func TestURLify(t *testing.T) {
	var tests = []struct {
		chain1 string
		expected string
	} {
		{"a","a"},
		{"a a", "a%20a"},
		{"a ", "a"},
		{"a %20", "a%20%20"},
	}

	for _, test := range tests {
		if URLify(test.chain1) != test.expected {
			t.Errorf("Should return %s, %s returned", test.expected, URLify(test.chain1))
		}
	}
}
