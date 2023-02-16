package list

import (
	"reflect"
	"testing"
)

func TestAddFirst(t *testing.T) {
	myList := NewLinkedList[int]()
	myList.AddFirst(1)
	myList.AddFirst(3)
	myList.AddFirst(5)

	expected := []int{5, 3, 1}
	values := myList.Values()

	if !reflect.DeepEqual(expected, values) {
		t.Errorf("Expected %v, got %v", expected, values)
	}
}

func TestAddLast(t *testing.T) {
	myList := NewLinkedList[int]()
	myList.AddLast(1)
	myList.AddLast(3)
	myList.AddLast(5)

	expected := []int{1, 3, 5}
	values := myList.Values()

	if !reflect.DeepEqual(expected, values) {
		t.Errorf("Expected %v, got %v", expected, values)
	}
}
