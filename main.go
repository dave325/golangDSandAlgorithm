package main

import "fmt"

func main() {
	processStack()
	processQueue()
	processLinkedList()
}

func processStack() {
	fmt.Println("Start Stack")
	stack := Stack{}
	fmt.Println(stack.IsEmpty())
	stack.Push(2)
	fmt.Println(stack.Peek())
	fmt.Println(stack.IsEmpty())
	fmt.Println(stack.IsFull())
	stack.Pop()
	fmt.Println("After Pop")
	fmt.Println(stack.Peek())
	fmt.Println(stack.IsEmpty())
	fmt.Println(stack.IsFull())
}

func processQueue() {
	fmt.Println("Start Queue")
	queue := Queue{}
	fmt.Println(queue.IsEmpty())
	queue.Enqueue(2)
	queue.Enqueue(3)
	fmt.Println(queue.Peek())
	fmt.Println(queue.IsEmpty())
	fmt.Println(queue.IsFull())
	queue.Dequeue()
	fmt.Println("After Pop")
	fmt.Println(queue.Peek())
	fmt.Println(queue.IsEmpty())
	fmt.Println(queue.IsFull())
}

func processLinkedList() {
	fmt.Println("Start Linked List")

	ll := LinkedList{}
	ll.addItem(1)
	ll.addItem(2)
	ll.addItem(3)
	ll.printList()
	ll.deleteItem(2)
	ll.printList()

}
