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
	var n = node{value: 1, next: nil}
	addNode(&n, 2)
	addNode(&n, 3)

	checkNextExists(t, &n)

	if n.next.value != 2 {
		t.Errorf("Next head should be equal to 2")
	}
	n = *n.next
	if n.next == nil || n.next.value != 3 {
		t.Errorf("Error with last node")
	}

	// Test if head is nil
	// var h = head{h: nil}
	// var test = head{h: &n}
	// h.h = nil
	// addNode(h.h, 1)
	// checkNodeValue(t, h.h, 1)
}

func checkNodeValue(t *testing.T, n *node, expected int) {
	if n == nil {
		t.Errorf("Nil node")
	}
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

func TestPartition(t *testing.T) {
	var n = node{value: 3, next: nil}
	addNodes(&n, []int{5,8,5,10,2,1})
	var h1 = head{h: &n}
	partition(&h1, 5)

	var expected = node{value: 3, next: nil}
	addNodes(&expected, []int{2,1,5,8,5,10})
	var h2 = head{h: &expected}

	if !equals(&h1, &h2) {
		t.Errorf("TestPartition does not work as expected")
		displayList(h1.h)
		displayList(h2.h)
	}
}

func TestSumListAux(t *testing.T) {
	var n = node{value: 7, next: nil}
	addNodes(&n, []int{1,6})
	var h = head{h: &n}

	if sumListAux(&h) != 617 {
		t.Errorf("sumListAux does not work as expected")
	}
}

func TestSumListsReverse(t *testing.T) {
	h1 := createHead([]int{7,1,6})
	h2 := createHead([]int{5,9,2})
	expected := createHead([]int{2,1,9})
	res := sumListsReverse(h1, h2)

	if !equals(expected, res) {
		t.Errorf("SumListsReverse does not work as expected")
		displayList(expected.h)
		displayList(res.h)
	}
}
