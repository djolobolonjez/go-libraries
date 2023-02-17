package list

import (
	"errors"
	"fmt"
)

type ListNode[T any] struct {
	data T
	next *ListNode[T]
}

type LinkedList[T any] struct {
	head *ListNode[T]
	tail *ListNode[T]
	size int
}

func NewLinkedList[T any]() *LinkedList[T] {
	list := LinkedList[T]{
		nil,
		nil,
		0,
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

func (list *LinkedList[T]) Len() int {
	return list.size
}

func (list *LinkedList[T]) Front() *ListNode[T] {
	return list.head
}

func (node *ListNode[T]) Next() *ListNode[T] {
	return node.next
}

func (list *LinkedList[T]) AddLast(data T) *LinkedList[T] {
	node := list.newNode(data)

	if list.Front() == nil {
		list.head, list.tail = node, node
	} else {
		list.tail.next = node
		list.tail = list.tail.next
	}
	list.size++

	return list
}

func (list *LinkedList[T]) AddFirst(data T) *LinkedList[T] {
	node := list.newNode(data)
	node.next = list.head
	list.head = node
	list.size++

	return list
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

func (list *LinkedList[T]) RemoveFirst() (T, error) {
	return list.RemoveAt(0)
}

func (list *LinkedList[T]) RemoveLast() (T, error) {
	return list.RemoveAt(list.Len() - 1)
}

func (list *LinkedList[T]) RemoveAt(pos int) (T, error) {
	var data T

	if pos >= list.Len() {
		return data, errors.New("index out of bounds")
	}

	index := 0
	curr := list.Front()
	var prev *ListNode[T] = nil
	for index < pos {
		prev = curr
		curr = curr.Next()
		index++
	}

	data = curr.data
	if prev == nil {
		list.head = list.head.next
	} else {
		prev.next = curr.next
	}
	list.size--

	return data, nil
}
