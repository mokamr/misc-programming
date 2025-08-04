package main

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func addToEnd(list *LinkedList, a int) {
	newNode := &Node{data: a, next: nil}

	if list.head == nil {
		list.head = newNode
		return
	}

	current := list.head

	for current.next != nil {
		current = current.next
	}

	current.next = newNode
}

func addToFront(list *LinkedList, a int) {

	// if list.head == nil {
	// 	newNode := &Node{data: a, next: nil}
	// 	list.head = newNode
	// }
	newNode := &Node{data: a, next: list.head}

	list.head = newNode
}

func printList(list *LinkedList) {

	if list.head == nil {
		fmt.Print("Empty List")
	}

	current := list.head
	for current != nil {
		fmt.Println(current.data)

		current = current.next
	}

}

func main() {
	// node := &Node{data: 1, next: nil}
	// node2 := &Node{data: 2, next: nil}
	list := &LinkedList{head: nil}

	addToEnd(list, 5)
	addToEnd(list, 10)
	addToFront(list, 20)
	addToFront(list, 20)
	printList(list)
}
