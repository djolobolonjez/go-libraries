package queue

import (
	"errors"

	"golang.org/x/exp/constraints"
)

type PriorityQueue[T constraints.Ordered] struct {
	data []T
}

func NewPriorityQueue[T constraints.Ordered]() *PriorityQueue[T] {
	q := PriorityQueue[T]{make([]T, 0)}

	return &q
}

func getLeftChild(index int) int {
	return 2*index + 1
}

func getRightChild(index int) int {
	return 2*index + 2
}

func getParent(index int) int {
	return (index - 1) / 2
}

func (this *PriorityQueue[T]) Len() int {
	return len(this.data)
}

func (this *PriorityQueue[T]) moveUp(index int) {
	parent := getParent(index)

	if index > 0 && this.data[index] > this.data[parent] {
		this.data[index], this.data[parent] = this.data[parent], this.data[index]
		this.moveUp(parent)
	}
}

func (this *PriorityQueue[T]) moveDown(index int) {
	max := index
	left, right := getLeftChild(index), getRightChild(index)

	if left < this.Len() && this.data[left] > this.data[max] {
		max = left
	}
	if right < this.Len() && this.data[right] > this.data[max] {
		max = right
	}

	if max != index {
		this.data[max], this.data[index] = this.data[index], this.data[max]
		this.moveDown(max)
	}
}

func (this *PriorityQueue[T]) Top() (T, error) {
	var value T
	if this.Len() == 0 {
		return value, errors.New("queue is empty")
	}

	value = this.data[0]
	return value, nil
}

func (this *PriorityQueue[T]) Push(value T) *PriorityQueue[T] {
	this.data = append(this.data, value)

	this.moveUp(this.Len() - 1)
	return this
}

func (this *PriorityQueue[T]) Pop() {
	if this.Len() == 0 {
		return
	}

	rear := this.Len() - 1

	this.data[0] = this.data[rear]
	this.data = this.data[:rear]

	this.moveDown(0)
}
