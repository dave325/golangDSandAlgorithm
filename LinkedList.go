package main

import "fmt"

type LinkedList struct {
	Data int
	Next *LinkedList
}

func (ll *LinkedList) addItem(item int) {

	var head *LinkedList
	fmt.Println(ll.Next)
	for ll.Next != nil {
		head = ll.Next
	}
	newItem := LinkedList{}
	newItem.Data = item
	newItem.Next = nil
	head.Next = &newItem
}

func (ll *LinkedList) deleteItem(item int) {

	var head *LinkedList
	var prevItem *LinkedList
	var nextItem *LinkedList
	for ll.Next != nil {
		if head.Next.Data == item {
			prevItem = head
			nextItem = head.Next.Next
			prevItem.Next = nextItem
			break
		}
		head = ll.Next
	}
}
