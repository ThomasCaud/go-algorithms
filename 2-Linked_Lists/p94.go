package main

import "fmt"

type node struct {
	value int
	next *node
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

func contain(arr []int, value int) bool {
	for _, v := range arr {
		if v == value {
			return true
		}
	}
	return false
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
