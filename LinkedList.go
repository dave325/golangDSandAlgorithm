package main

import "fmt"

type Node struct {
	Data int
	Next *Node
}
type LinkedList struct {
	Node *Node
	Size int
}

func (ll *LinkedList) addItem(item int) {

	if ll.Size == 0 {
		node := Node{item, nil}
		ll.Node = &node
		ll.Size++
		return
	}
	head := ll.Node
	for head.Next != nil {
		head = head.Next
	}
	newItem := Node{item, nil}
	head.Next = &newItem
	ll.Size++
}

func (ll *LinkedList) deleteItem(item int) {

	head := ll.Node
	var prevItem *Node
	var nextItem *Node
	for head.Next != nil {
		if head.Next.Data == item {
			prevItem = head
			nextItem = head.Next.Next
			prevItem.Next = nextItem
			ll.Size--
			break
		}
		head = head.Next
	}
}

func (ll *LinkedList) insertAfterIdx(item int, itemToLookup int) {
	if ll.Size == 0 {
		node := Node{item, nil}
		ll.Node = &node
		ll.Size++
		return
	}

	head := ll.Node
	var currentItem *Node
	var nextItem *Node
	fmt.Println(ll.Node)
	for head.Next != nil {
		if head.Data == itemToLookup {
			currentItem = head
			node := Node{item, nil}
			nextItem = head.Next
			currentItem.Next = &node
			node.Next = nextItem
			ll.Size++
			break
		}
		head = head.Next
	}
}

func (ll LinkedList) printList() {
	head := ll.Node
	for head != nil {
		fmt.Println(head.Data)
		head = head.Next
	}
}
