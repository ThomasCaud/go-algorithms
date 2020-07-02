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

func displayList(n *node) {
	for n != nil {
		fmt.Print(" -> ", n.value)
		n = n.next
	}
}

func contain(arr []int, value int) bool {
	for _, v := range arr {
		if v == value {
			return true
		}
	}
	return false
}

func removeDuplicate(n *node) {
	values := []int{}
	var previous *node
	for n != nil {
		if previous != nil && contain(values, n.value) && n.next != nil {
			previous.next = n.next
		} else {
			values = append(values, n.value)
		}
		previous = n
		n = n.next
	}
}

/*
func removeDuplicateWithoutBuffer(n *node) {
	for n != nil {
		var previous *node
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
}*/
