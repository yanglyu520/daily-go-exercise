package main

import "fmt"

// Node represents a node in the circular linked list
type Node struct {
	data int
	next *Node
}

// createCircularLinkedList creates a circular linked list with n nodes
func createCircularLinkedList(n int) *Node {
	if n <= 0 {
		return nil
	}

	head := &Node{data: 0}
	current := head

	for i := 1; i < n; i++ {
		current.next = &Node{data: i}
		current = current.next
	}

	current.next = head // Make it circular
	return head
}

// josephusUsingLinkedList solves the Josephus problem using a circular linked list
func josephusUsingLinkedList(n, k int) int {
	if n <= 0 || k <= 0 {
		return -1
	}

	head := createCircularLinkedList(n)
	current := head

	for current.next != current {
		for i := 0; i < k-1; i++ {
			current = current.next
		}
		fmt.Printf("Removing %d\n", current.next.data)
		current.next = current.next.next // Remove the k-th node
	}

	return current.data
}

func main() {
	n := 7 // Number of people
	k := 3 // Step count

	result := josephusUsingLinkedList(n, k)
	fmt.Printf("The last remaining person is at position %d\n", result)
}
