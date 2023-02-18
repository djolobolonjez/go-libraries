package queue

import "errors"

type Deque[T any] struct {
	data []T
	size int
}

func NewDeque[T any]() *Deque[T] {
	q := Deque[T]{
		make([]T, 0),
		0,
	}

	return &q
}

func (this *Deque[T]) Len() int {
	return this.size
}

func (this *Deque[T]) PushFront(value T) *Deque[T] {
	q := make([]T, this.Len()+1)
	q[0] = value
	copy(q[1:], this.data)
	this.data = q
	this.size++

	return this
}

func (this *Deque[T]) PushBack(value T) *Deque[T] {
	this.data = append(this.data, value)
	this.size++

	return this
}

func (this *Deque[T]) PopBack() {
	this.data = this.data[:this.Len()-1]
	this.size--
}

func (this *Deque[T]) PopFront() {
	this.data = this.data[1:]
	this.size--
}

func (this *Deque[T]) PeekBack() (T, error) {
	var value T
	if this.Len() == 0 {
		return value, errors.New("queue is empty")
	}
	value = this.data[this.Len()-1]
	return value, nil
}

func (this *Deque[T]) PeekFront() (T, error) {
	var value T
	if this.Len() == 0 {
		return value, errors.New("queue is empty")
	}
	value = this.data[0]
	return value, nil
}
