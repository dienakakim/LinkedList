# LinkedList

## Description

This project contains code implementing a generic doubly linked list in Go. Takes a placeholder element to use as the list's key type, and requires all argument keys to have that type.

## Exposed methods

```go
func NewLinkedList(t interface{}) LinkedList
```

NewLinkedList is a factory function creating a new LinkedList with a given element as its keyType.

```go
func (l *LinkedList) InsertBack(key interface{}) (bool, error)
func (l *LinkedList) InsertFront(key interface{}) (bool, error)
func (l *LinkedList) InsertAt(index uint, key interface{}) (bool, error)
```

Insert methods -- InsertBack inserts at the end of the list, InsertFront inserts at the front, and InsertAt inserts at a given index.

```go
func (l *LinkedList) At(index uint) (interface{}, error)
```

Retrieves an element at the given index.

```go
func (l *LinkedList) IndexOf(t interface{}) (uint, error)
```

Finds a given element and returns its index if it is in the list.

```go
func (l *LinkedList) String() string
```

Returns the list's string representation.

```go
func (l *LinkedList) RemoveBack() (interface{}, error)
func (l *LinkedList) RemoveFront() (interface{}, error)
func (l *LinkedList) RemoveAt(index uint) (interface{}, error)
```

Remove methods -- same semantics as detailed in insert methods.

```go
func (l *LinkedList) Remove(key interface{}) (bool, error)
```

Removes a given element from the list if found.
