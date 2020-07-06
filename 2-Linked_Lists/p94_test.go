package main

import "testing"

func checkNextExists(t *testing.T, n *node) {
	if n.next == nil {
		t.Errorf("Should have a next node defined")
	}
}

func TestGetSize(t *testing.T) {
	var head = head{h: nil}
	if getSize(&head) != 0 {
		t.Errorf("GetSize should return 0")
	}

	n := node{value: 1, next: nil}
	head.h = &n
	if getSize(&head) != 1 {
		t.Errorf("GetSize should return 1, returned %d", getSize(&head))
	}
}

func TestEquals(t *testing.T) {
	var n1 = node{value: 1, next: nil}
	var h1 = head{h: &n1}
	addNodes(&n1, []int{2,3,4,5})

	var n2 = node{value: 1, next: nil}
	var h2 = head{h: &n2}

	if equals(&h1, &h2) {
		t.Errorf("Equals should return false")
	}

	addNodes(&n2, []int{2,3,4,5})
	if !equals(&h1, &h2) {
		t.Errorf("Equals should return true")
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

func TestGetKToLast(t *testing.T) {
	var head = node{value: 1, next: nil}
	addNodes(&head, []int{2,3,4,5,6,7,8,9,10})

	lastNode := getKToLast(&head, 0)
	checkNodeValue(t, lastNode, 10)

	lastNode = getKToLast(&head, 4)
	checkNodeValue(t, lastNode, 6)

	lastNode = getKToLast(&head, 789)
	checkNodeValue(t, lastNode, 1)

	lastNode = getKToLast(&head, -50)
	checkNodeValue(t, lastNode, 10)
}

func TestDeleteMiddleNode(t *testing.T) {
	var n = node{value: 1, next: nil}
	addNodes(&n, []int{2,3})
	var h1 = head{h: &n}

	deleteMiddleNode(&h1)

	var expected = node{value: 1, next: nil}
	addNodes(&expected, []int{3})
	var h2 = head{h: &expected}

	if !equals(&h1, &h2) {
		t.Errorf("DeleteMiddleNode does not work as expected")
		displayList(h1.h)
		displayList(h2.h)
	}
}
