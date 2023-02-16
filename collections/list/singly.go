package list

import "fmt"

type ListNode[T any] struct {
	data T
	next *ListNode[T]
}

type LinkedList[T any] struct {
	head *ListNode[T]
	tail *ListNode[T]
}

func NewLinkedList[T any]() *LinkedList[T] {
	list := LinkedList[T]{
		nil,
		nil,
	}
	return &list
}

func (list *LinkedList[T]) newNode(data T) *ListNode[T] {
	n := ListNode[T]{
		data,
		nil,
	}
	return &n
}

func (list *LinkedList[T]) Front() *ListNode[T] {
	return list.head
}

func (node *ListNode[T]) Next() *ListNode[T] {
	return node.next
}

func (list *LinkedList[T]) AddLast(data T) {
	node := list.newNode(data)

	if list.Front() == nil {
		list.head, list.tail = node, node
	} else {
		list.tail.next = node
		list.tail = list.tail.next
	}
}

func (list *LinkedList[T]) AddFirst(data T) {
	node := list.newNode(data)
	node.next = list.head
	list.head = node
}

func (list *LinkedList[T]) PrintList() {
	for n := list.Front(); n != nil; n = n.Next() {
		fmt.Println(n.data)
	}
}

func (list *LinkedList[T]) Values() []T {
	values := []T{}

	for node := list.Front(); node != nil; node = node.Next() {
		values = append(values, node.data)
	}

	return values
}
