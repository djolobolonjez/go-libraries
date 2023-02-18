package queue

import (
	"reflect"
	"testing"
)

func TestPushFront(t *testing.T) {
	myDeque := NewDeque[int]()
	myDeque.PushFront(1).PushFront(2).PushFront(3)

	expected := []int{3, 2, 1}
	values := myDeque.data

	if !reflect.DeepEqual(expected, values) {
		t.Errorf("Expected %v, got %v", expected, values)
	}
}

func TestPushBack(t *testing.T) {
	myDeque := NewDeque[int]()
	myDeque.PushBack(1).PushBack(2).PushBack(3)

	expected := []int{1, 2, 3}
	values := myDeque.data

	if !reflect.DeepEqual(expected, values) {
		t.Errorf("Expected %v, got %v", expected, values)
	}
}

func TestPopBack(t *testing.T) {
	myDeque := NewDeque[int]()
	myDeque.PushBack(1).PushBack(2).PushBack(3)
	myDeque.PopBack()

	expected := []int{1, 2}
	values := myDeque.data

	if !reflect.DeepEqual(expected, values) {
		t.Errorf("Expected %v, got %v", expected, values)
	}
}

func TestPopFront(t *testing.T) {
	myDeque := NewDeque[int]()
	myDeque.PushBack(1).PushBack(2).PushBack(3)
	myDeque.PopFront()

	expected := []int{2, 3}
	values := myDeque.data

	if !reflect.DeepEqual(expected, values) {
		t.Errorf("Expected %v, got %v", expected, values)
	}
}

func TestPeekFront(t *testing.T) {
	myDeque := NewDeque[int]()
	myDeque.PushBack(1).PushBack(2).PushBack(3)

	expected := 1
	value, err := myDeque.PeekFront()
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(expected, value) {
		t.Errorf("Expected %v, got %v", expected, value)
	}
}

func TestPeekBack(t *testing.T) {
	myDeque := NewDeque[int]()
	myDeque.PushBack(1).PushBack(2).PushBack(3)

	expected := 3
	value, err := myDeque.PeekBack()
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(expected, value) {
		t.Errorf("Expected %v, got %v", expected, value)
	}
}
