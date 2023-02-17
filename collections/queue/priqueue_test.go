package queue

import (
	"reflect"
	"testing"
)

func TestPush(t *testing.T) {
	myPriorityQueue := NewPriorityQueue[int]()
	myPriorityQueue.Push(1).Push(3).Push(2).Push(4)

	expected := []int{4, 3, 2, 1}
	values := myPriorityQueue.data

	if !reflect.DeepEqual(expected, values) {
		t.Errorf("Expected %v, got %v", expected, values)
	}
}

func TestPop(t *testing.T) {
	myPriorityQueue := NewPriorityQueue[int]()
	myPriorityQueue.Push(1).Push(3).Push(2)

	myPriorityQueue.Pop()
	expected := []int{2, 1}
	values := myPriorityQueue.data

	if !reflect.DeepEqual(expected, values) {
		t.Errorf("Expected %v, got %v", expected, values)
	}
}

func TestTop(t *testing.T) {
	myPriorityQueue := NewPriorityQueue[int]()
	myPriorityQueue.Push(1).Push(3).Push(2)

	value, err := myPriorityQueue.Top()
	if err != nil {
		t.Error(err)
	}

	expected := 3
	if value != expected {
		t.Errorf("Expected %v, got %v", expected, value)
	}
}
