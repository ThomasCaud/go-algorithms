package main

import "testing"

func checkNextExists(t *testing.T, n *node) {
	if n.next == nil {
		t.Errorf("Should have a next node defined")
	}
}

func TestAddNode(t *testing.T) {
	var head = node{value: 1, next: nil}
	addNode(&head, 2)
	addNode(&head, 3)

	checkNextExists(t, &head)

	if head.next.value != 2 {
		t.Errorf("Next head should be equal to 2")
	}
	head = *head.next
	if head.next == nil || head.next.value != 3 {
		t.Errorf("Error with last node")
	}
}

func checkNodeValue(t *testing.T, n *node, expected int) {
	if n.value != expected {
		t.Errorf("Expected value %d, got %d", expected, n.value)
	}
}

func checkNextIsNil(t *testing.T, n *node) {
	if n.next != nil {
		t.Errorf("Expected next node to be nil")
	}
}

type fn func(n *node)

func RemoveDuplicateTests(t *testing.T, f fn) {
	var head = node{value: 1, next: nil}

	addNodes(&head, []int{2,2,3,2})

	f(&head)

	checkNodeValue(t, &head, 1)

	head = *head.next
	checkNodeValue(t, &head, 2)

	head = *head.next
	checkNodeValue(t, &head, 3)

	checkNextIsNil(t, &head)
}

func TestRemoveDuplicate(t *testing.T) {
	RemoveDuplicateTests(t, removeDuplicate)
}

func TestRemoveDuplicateWithoutBuffer(t *testing.T) {
	RemoveDuplicateTests(t, removeDuplicateWithoutBuffer)
}
