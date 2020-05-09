// Package LinkedList implements a doubly linked list with type checking.
package LinkedList

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

// Node is a node type for containing keys for use in the LinkedList struct.
type Node struct {
	prev, next *Node
	key        interface{}
	keyType    reflect.Type
}

// LinkedList implements a doubly linked list.
type LinkedList struct {
	head, tail *Node
	size       uint
	keyType    reflect.Type
}

func newNode(key interface{}) Node {
	n := Node{nil, nil, key, reflect.TypeOf(key)}
	return n
}

// NewLinkedList is a factory function creating a new LinkedList with a given element as its keyType.
func NewLinkedList(key interface{}) LinkedList {
	l := LinkedList{nil, nil, 0, reflect.TypeOf(key)}
	return l
}

// InsertBack inserts an element with KeyType type into the back of the list.
func (l *LinkedList) InsertBack(key interface{}) (bool, error) {
	// Enforce same type restriction
	if reflect.TypeOf(key) != l.keyType {
		return false, fmt.Errorf("Expected key type %s", l.keyType.String())
	}

	// Check if list is empty
	if l.head == nil && l.tail == nil {
		l.head = &Node{nil, nil, key, reflect.TypeOf(key)}
		l.tail = l.head
		l.size++
		return true, nil
	}

	// Else we insert into the back of the list
	l.tail.next = &Node{l.tail, nil, key, reflect.TypeOf(key)}
	l.tail = l.tail.next
	l.size++
	return true, nil
}

// InsertFront inserts an element with KeyType type into the front of the list.
func (l *LinkedList) InsertFront(key interface{}) (bool, error) {
	// Enforce same type restriction
	if reflect.TypeOf(key) != l.keyType {
		return false, fmt.Errorf("Expected key type %s", l.keyType.String())
	}

	// Check if list is empty
	if l.head == nil && l.tail == nil {
		l.head = &Node{nil, nil, key, reflect.TypeOf(key)}
		l.tail = l.head
		l.size++
		return true, nil
	}

	// Else we insert into the front of the list
	l.head.prev = &Node{nil, l.head, key, reflect.TypeOf(key)}
	l.head = l.head.prev
	l.size++
	return true, nil
}

func (l *LinkedList) InsertAt(index uint, key interface{}) (bool, error) {
	// Check if index is out of range
	if index > l.size {
		return false, fmt.Errorf("Index %d out of range -- size is %d", index, l.size)
	}

	// Check if given key matches list's keyType
	if reflect.TypeOf(key) != l.keyType {
		return false, fmt.Errorf("Expected key type %s", l.keyType.String())
	}

	// Special cases:
	// - index = 0
	// - index = n (past-end)
	if index == 0 {
		return l.InsertFront(key)
	} else if index == l.size {
		return l.InsertBack(key)
	}

	// index is in [1, size)
	// get to the node at given index
	var current *Node
	if index < l.size/2 {
		current = l.head
		currentIndex := uint(0)
		for currentIndex < index {
			currentIndex++
			current = current.next
		}
	} else {
		current = l.tail
		currentIndex := uint(0)
		for currentIndex < l.size-index-1 {
			currentIndex++
			current = current.prev
		}
	}

	// shift every node from current onwards to the right,
	// and insert new node with given key
	inserted := &Node{current.prev, current, key, reflect.TypeOf(key)}
	fmt.Println(current)
	current.prev.next = inserted
	current.prev = inserted
	l.size++

	// Done.
	return true, nil
}

// At returns the key at given index from the list
func (l *LinkedList) At(index uint) (interface{}, error) {
	// Check if index is within [0, l.Size)
	if index >= l.size {
		return nil, fmt.Errorf("Index %d out of range -- size is %d", index, l.size)
	}

	// Check if index is within the first half of the range [0, l.Size),
	// if so we start at l.head, else we start at l.tail
	var key interface{}
	if index < l.size/2 {
		current := l.head
		currentIndex := uint(0)
		for currentIndex < index {
			currentIndex++
			current = current.next
		}
		key = current.key
	} else {
		current := l.tail
		currentIndex := uint(0)
		for currentIndex < l.size-index-1 {
			currentIndex++
			current = current.prev
		}
		key = current.key
	}

	// Key found
	return key, nil
}

// IndexOf attempts to look for a given key in the list. If found, it returns its index.
func (l *LinkedList) IndexOf(t interface{}) (uint, error) {
	// Check if list is empty
	if l.size == 0 {
		return 0, errors.New("List is empty")
	}

	// Check if types match
	if reflect.TypeOf(t) != l.keyType {
		return l.size, fmt.Errorf("Unmatched types: got %s, expected %s", reflect.TypeOf(t), l.keyType)
	}

	// Else we can begin searching. Start from l.head
	current := l.head
	index := uint(0)
	for current != nil {
		if current.key == t {
			return index, nil
		}
		index++
		current = current.next
	}

	// If we get here, then t is not found
	return l.size, errors.New("Key not found")
}

// Size returns the list's size.
func (l *LinkedList) Size() uint {
	return l.size
}

// String returns the list's string representation. Uses fmt's Sprintf.
func (l *LinkedList) String() string {
	var output strings.Builder
	output.WriteRune('[')
	current := l.head
	for current != nil {
		output.WriteString(fmt.Sprintf("%v", current.key))
		if current.next != nil {
			output.WriteString(", ")
		}
		current = current.next
	}
	output.WriteRune(']')
	return output.String()
}

// RemoveBack unlinks the last element of the list.
func (l *LinkedList) RemoveBack() (interface{}, error) {
	// Check if list is empty first
	if l.size == 0 {
		return nil, errors.New("List is empty")
	}

	// If l.size == 1, this is the only node in the list
	var removed *Node
	if l.size == 1 {
		removed = l.head
		l.head = nil
		l.tail = nil
		l.size = 0
	} else {
		// We have at least 2 nodes
		removed = l.tail
		l.tail = l.tail.prev
		l.tail.next = nil
		l.size--
	}

	// Done.
	return removed.key, nil
}

// RemoveFront unlinks the first element in the list.
func (l *LinkedList) RemoveFront() (interface{}, error) {
	// Check if list is empty first
	if l.size == 0 {
		return false, errors.New("List is empty")
	}

	// If l.size == 1, this is the only node in the list
	var removed *Node
	if l.size == 1 {
		removed = l.head
		l.head = nil
		l.tail = nil
		l.size = 0
	} else {
		// We have at least 2 nodes
		removed = l.head
		l.head = l.head.next
		l.head.prev = nil
		l.size--
	}

	// Done.
	return removed.key, nil
}

func (l *LinkedList) RemoveAt(index uint) (interface{}, error) {
	// Check if index is out of range
	if index >= l.size {
		return nil, errors.New("Index out of range")
	}

	// Special cases:
	// - index = 0
	// - index = n-1
	if index == 0 {
		return l.RemoveFront()
	} else if index == l.size-1 {
		// index == 0 cannot happen, this implies l.size >= 2
		// Remove tail
		return l.RemoveBack()
	}

	// index is in interior
	// Check if index is within the first half of the range [0, l.Size),
	// if so we start at l.head, else we start at l.tail
	var current *Node
	if index < l.size/2 {
		current = l.head
		currentIndex := uint(0)
		for currentIndex < index {
			currentIndex++
			current = current.next
		}
	} else {
		current = l.tail
		currentIndex := uint(0)
		for currentIndex < l.size-index-1 {
			currentIndex++
			current = current.prev
		}
	}

	// current is the node to be removed, and as it is in the interior,
	// it has prev and next nodes
	fmt.Println(current)
	current.prev.next = current.next // detach prev
	current.next.prev = current.prev // detach next
	l.size--

	// Done.
	return current.key, nil
}

func (l *LinkedList) Remove(key interface{}) (bool, error) {
	// Check if list is empty
	if l.size == 0 {
		return false, errors.New("List is empty")
	}

	// Check if types match
	if reflect.TypeOf(key) != l.keyType {
		return false, fmt.Errorf("Expected type %s", l.keyType.String())
	}

	// Now linear search the list
	current := l.head
	for current != nil {
		if current.key == key {
			break
		}
		current = current.next
	}

	// Check if we found the key yet
	if current == nil {
		return false, errors.New("Key not found")
	}

	// We found the key
	// Special cases:
	// - key is at head
	// - key is at tail
	// If key is at head, call RemoveFront
	if current == l.head {
		val, err := l.RemoveFront()
		return val != nil, err
	} else if current == l.tail {
		val, err := l.RemoveBack()
		return val != nil, err
	}

	// Key is in interior, so prev and next nodes exist
	// Unlink prev and next
	current.prev.next = current.next
	current.next.prev = current.prev
	l.size--

	// Done.
	return true, nil
}
