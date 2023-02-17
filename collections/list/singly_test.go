package list

import (
	"reflect"
	"testing"
)

func TestAddFirst(t *testing.T) {
	myList := NewLinkedList[int]()
	myList.AddFirst(1).AddFirst(3).AddFirst(5)

	expected := []int{5, 3, 1}
	values := myList.Values()

	if !reflect.DeepEqual(expected, values) {
		t.Errorf("Expected %v, got %v", expected, values)
	}
}

func TestAddLast(t *testing.T) {
	myList := NewLinkedList[int]()
	myList.AddLast(1).AddLast(3).AddLast(5)

	expected := []int{1, 3, 5}
	values := myList.Values()

	if !reflect.DeepEqual(expected, values) {
		t.Errorf("Expected %v, got %v", expected, values)
	}
}

func TestRemoveFirst(t *testing.T) {
	myList := NewLinkedList[int]()
	myList.AddFirst(5).AddFirst(3).AddFirst(7)
	myList.RemoveFirst()

	expected := []int{3, 5}
	values := myList.Values()

	if !reflect.DeepEqual(expected, values) {
		t.Errorf("Expected %v, got %v", expected, values)
	}
}

func TestRemoveLast(t *testing.T) {
	myList := NewLinkedList[int]()
	myList.AddFirst(5).AddFirst(3).AddFirst(7)
	myList.RemoveLast()

	expected := []int{7, 3}
	values := myList.Values()

	if !reflect.DeepEqual(expected, values) {
		t.Errorf("Expected %v, got %v", expected, values)
	}
}

func TestRemoveAt(t *testing.T) {
	myList := NewLinkedList[int]()
	myList.AddFirst(5).AddFirst(3).AddFirst(7)
	myList.RemoveAt(1)

	expected := []int{7, 5}
	values := myList.Values()

	if !reflect.DeepEqual(expected, values) {
		t.Errorf("Expected %v, got %v", expected, values)
	}
}

func TestInsertAt(t *testing.T) {
	myList := NewLinkedList[int]()
	myList.InsertAt(0, 1).InsertAt(1, 3).InsertAt(1, 5)

	expected := []int{1, 5, 3}
	values := myList.Values()

	if !reflect.DeepEqual(expected, values) {
		t.Errorf("Expected %v, got %v", expected, values)
	}
}
