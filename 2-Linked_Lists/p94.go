package main

import "fmt"
import "math"

type node struct {
	value int
	next *node
}

type head struct {
	h *node
}

func addNode(head *node, value int) {
	var n = head
	for n.next != nil {
		n = n.next
	}
	var newNode = node{value: value, next: nil}

	n.next = &newNode
}

func addNodes(head *node, value []int) {
	for _, v := range value {
		addNode(head, v)
	}
}

func displayList(n *node) {
	for n != nil {
		fmt.Print(" -> ", n.value)
		n = n.next
	}
	fmt.Println("")
}

func getSize(list *head) int {
	size := 0
	n := list.h
	for n != nil {
		size++
		n = n.next
	}
	return size
}

func contain(arr []int, value int) bool {
	for _, v := range arr {
		if v == value {
			return true
		}
	}
	return false
}

func equals(h1 *head, h2 *head) bool {
	n1 := h1.h
	n2 := h2.h

	for n1 != nil && n2 != nil {
		if (n1.value != n2.value) {
			return false
		}
		n1 = n1.next
		n2 = n2.next
	}

	return n1 == nil && n2 == nil
}

func createHead(arr []int) *head {
	var n = node{value: arr[0], next: nil}
	addNodes(&n, arr[1:])
	var h = head{h: &n}

	return &h
}

// Q 2.1 - Remove Dups
func removeDuplicate(n *node) {
	values := []int{}
	var previous *node
	for n != nil {
		if previous != nil && contain(values, n.value) {
			previous.next = n.next
		} else {
			values = append(values, n.value)
		}
		previous = n
		n = n.next
	}
}

// Q 2.1 - Remove Dups (follow up)
func removeDuplicateWithoutBuffer(n *node) {
	for n != nil {
		var previous *node = n
		nextNodes := n.next
		for nextNodes != nil  {
			if nextNodes.value == n.value {
				previous.next = nextNodes.next
			}
			previous = nextNodes
			nextNodes = nextNodes.next
		}

		n = n.next
	}
}

// Q 2.2 - Return Kth to Last
func getKToLast(n *node, k int) *node {
	// first step: count number of element in the linked list
	nbNodes := -1
	list := n

	for list != nil {
		nbNodes++
		list = list.next
	}

	stop := nbNodes - k
	if stop < 0 {
		stop = 0
	}

	if stop > nbNodes {
		stop = nbNodes
	}

	// second step: iterate until being at Kth to last
	for stop > 0 {
		n = n.next
		stop --
	}
	return n
}

// Q 2.2 - Return Kth to Last (recursive)
func printKToLastRec(n *node, k int) int {
	if n == nil {
		return 0
	}
	i := printKToLastRec(n.next, k) + 1
	if i == k {
		fmt.Println(k, "th to the last node is", n.value)
	}
	return i
}

// Q 2.3
func deleteMiddleNode(h *head) {
	size := getSize(h)
	if size < 3 {
		// there is no middle
		return
	}
	middle := int(math.Round(float64(size) / 2))

	var previousNode, node *node
	node = h.h

	for i := 0 ; i < middle - 1 ; i++ {
		previousNode = node
		node = node.next
	}
	previousNode.next = node.next
}

// Q 2.4
func partition(h *head, value int) {
	var p1, p2, headP1, headP2 *node
	n := h.h

	for n != nil {
		if n.value < value {
			if p1 == nil {
				p1 = n
				headP1 = n
			} else {
				p1.next = n
				p1 = p1.next
			}
		} else {
			if p2 == nil {
				p2 = n
				headP2 = n
			} else {
				p2.next = n
				p2 = p2.next
			}
		}
		n = n.next
	}

	// Remove old queue at the end of the partition
	p2.next = nil

	// Reallocate memory to avoid infinite loop
	newHeadNode := node{value: headP1.value, next: headP1.next}
	h.h = &newHeadNode

	// Concat partitions
	headP1 = headP1.next
	for headP1.next != nil {
		headP1 = headP1.next
	}
	headP1.next = headP2
}

// For Q2.5 - 1->2->3 return 321
func sumListAux(h *head) int {
	total := 0
	coeff := 1

	n := h.h
	for n != nil {
		total += n.value * coeff
		coeff *= 10
		n = n.next
	}

	return total
}

// Q2.5
func sumListsReverse(h1 *head, h2 *head) *head {
	total := sumListAux(h1) + sumListAux(h2)
	divider := 1

	var n = node{value: total % 10, next: nil}
	var h = head{h: &n}

	for divider * 10 < total {
		divider *= 10
		addNode(&n, int(math.Round(float64(total) / float64(divider))) % 10)
	}

	return &h
}

func reverse(list *head) *head {
	n := list.h
	var newNode *node

	for n != nil {
		oldNode := newNode
		newNode = &node{value: n.value, next: oldNode}
		n = n.next
	}

	return &head{h: newNode}
}

// Q2.6
// Optimization: not necessary to loop on all the linked list
// We could stop at half of it
func palindrome(list *head) bool {
	reversedList := reverse(list)
	return equals(list, reversedList)
}
