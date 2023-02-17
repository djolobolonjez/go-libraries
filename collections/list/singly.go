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

func (this *LinkedList[T]) newNode(data T) *ListNode[T] {
	n := ListNode[T]{
		data,
		nil,
	}
	return &n
}

func (this *LinkedList[T]) Len() int {
	return this.size
}

func (this *LinkedList[T]) Front() *ListNode[T] {
	return this.head
}

func (this *ListNode[T]) Next() *ListNode[T] {
	return this.next
}

func (this *LinkedList[T]) AddLast(data T) *LinkedList[T] {
	node := this.newNode(data)

	if this.Front() == nil {
		this.head, this.tail = node, node
	} else {
		this.tail.next = node
		this.tail = this.tail.next
	}
	this.size++

	return this
}

func (this *LinkedList[T]) AddFirst(data T) *LinkedList[T] {
	node := this.newNode(data)
	node.next = this.head
	this.head = node
	this.size++

	return this
}

func (this *LinkedList[T]) PrintList() {
	for n := this.Front(); n != nil; n = n.Next() {
		fmt.Println(n.data)
	}
}

func (this *LinkedList[T]) Values() []T {
	values := []T{}

	for node := this.Front(); node != nil; node = node.Next() {
		values = append(values, node.data)
	}

	return values
}

func (this *LinkedList[T]) RemoveFirst() (T, error) {
	return this.RemoveAt(0)
}

func (this *LinkedList[T]) RemoveLast() (T, error) {
	return this.RemoveAt(this.Len() - 1)
}

func (this *LinkedList[T]) RemoveAt(pos int) (T, error) {
	var data T

	if pos >= this.Len() {
		return data, errors.New("index out of bounds")
	}

	index := 0
	curr := this.Front()
	var prev *ListNode[T] = nil
	for index < pos {
		prev = curr
		curr = curr.Next()
		index++
	}

	data = curr.data
	if prev == nil {
		this.head = this.head.next
	} else {
		prev.next = curr.next
	}
	this.size--

	return data, nil
}

func (this *LinkedList[T]) InsertAt(pos int, data T) *LinkedList[T] {
	if pos == this.Len() {
		return this.AddLast(data)
	} else if pos == 0 {
		return this.AddFirst(data)
	}

	node := this.newNode(data)
	var prev *ListNode[T]
	curr := this.Front()

	for index := 0; index < pos; index++ {
		prev = curr
		curr = curr.Next()
	}

	node.next = curr
	prev.next = node

	return this
}
