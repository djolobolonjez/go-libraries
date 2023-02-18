package list

import (
	"errors"
	"fmt"
)

type ListNode[T any] struct {
	data T
	next *ListNode[T]
	prev *ListNode[T]
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
		nil,
	}
	return &n
}

func (this *LinkedList[T]) findByIndex(pos int) *ListNode[T] {
	curr := this.Front()
	for index := 0; index < pos; index++ {
		curr = curr.Next()
	}

	return curr
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

func (this *LinkedList[T]) PeekFirst() (T, error) {
	var data T
	if this.Len() == 0 {
		return data, errors.New("list is empty")
	}

	data = this.findByIndex(0).data

	return data, nil
}

func (this *LinkedList[T]) PeekLast() (T, error) {
	var data T
	if this.Len() == 0 {
		return data, errors.New("list is empty")
	}
	data = this.tail.data
	return data, nil
}

func (this *LinkedList[T]) PeekAt(pos int) (T, error) {
	var data T
	if pos >= this.Len() {
		return data, errors.New("index out of bounds")
	}
	data = this.findByIndex(pos).data
	return data, nil
}

func (this *LinkedList[T]) AddLast(data T) *LinkedList[T] {
	node := this.newNode(data)

	if this.Front() == nil {
		this.head, this.tail = node, node
	} else {
		this.tail.next = node
		node.prev = this.tail
		this.tail = this.tail.next
	}
	this.size++

	return this
}

func (this *LinkedList[T]) AddFirst(data T) *LinkedList[T] {
	node := this.newNode(data)
	node.next = this.head
	if this.head != nil {
		this.head.prev = node
	} else {
		this.tail = node
	}
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

	node := this.findByIndex(pos)

	data = node.data

	if node.prev == nil {
		this.head = this.head.next
		if this.head != nil {
			this.head.prev = nil
		}
	} else if node.next == nil {
		this.tail = this.tail.prev
		this.tail.next = nil
	} else {
		node.next.prev = node.prev
		node.prev.next = node.next
		node.prev, node.next = nil, nil
	}

	this.size--

	return data, nil
}

func (this *LinkedList[T]) InsertAt(pos int, data T) *LinkedList[T] {

	if pos > this.Len() {
		pos = this.Len()
	}

	if pos == this.Len() {
		return this.AddLast(data)
	} else if pos == 0 {
		return this.AddFirst(data)
	}

	node := this.newNode(data)
	curr := this.findByIndex(pos)

	node.next = curr
	node.prev = curr.prev
	if node.next != nil {
		node.next.prev = node
	}
	if node.prev != nil {
		node.prev.next = node
	}

	return this
}
