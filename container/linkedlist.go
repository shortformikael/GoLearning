package container

import (
	"errors"
	"fmt"
)

type Node struct {
	data string
	next *Node
}

type LinkedList struct {
	head *Node
}

func (l *LinkedList) Append(data string) {
	newNode := &Node{data: data}
	if l.head == nil {
		l.head = newNode
		return
	}
	current := l.head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
}

func (l *LinkedList) Prepend(data string) {
	newNode := &Node{data: data, next: l.head}
	l.head = newNode
}

func (l *LinkedList) Print() {
	current := l.head
	for {
		if current != nil {
			fmt.Println(current.data)
			current = current.next
		} else {
			return
		}
	}
}

func (l *LinkedList) Delete(s string) error {
	if s == l.head.data {
		l.head = l.head.next
		return nil
	}
	current := l.head.next
	previous := l.head
	for current != nil {
		if current.data == s {
			previous.next = current.next
			return nil
		}
		previous = current
		current = current.next
	}
	return errors.New("ITEM DOESNT EXIST IN LIST")
}

func (l *LinkedList) DeleteAt(n int) error {
	if 0 > n || n >= l.Length() {
		return errors.New("INTEGER OUT OF BOUNDS")
	} else if n == 0 {
		l.head = l.head.next
		return nil
	}
	current := l.head.next
	previous := l.head
	for i := 0; i < n; i++ {
		previous = current
		current = current.next
	}
	previous.next = current.next
	return nil
}

func (l *LinkedList) Get(n int) (string, error) {
	if 0 > n || n >= l.Length() {
		return "", errors.New("INTERGER OUT OF BOUNDS")
	}
	current := l.head
	for i := 0; i < n; i++ {
		if current != nil {
			current = current.next
		}
	}
	return current.data, nil
}

func (l *LinkedList) Exists(s string) bool {
	current := l.head
	for current != nil {
		if current.data == s {
			return true
		}
		current = current.next
	}
	return false
}

func (l *LinkedList) Length() int {
	count := 0
	current := l.head
	for current != nil {
		count++
		current = current.next
	}
	return count
}

func (l *LinkedList) GetArray() []string {
	var rArray []string
	for i := 0; i < l.Length(); i++ {
		item, _ := l.Get(i)
		rArray = append(rArray, item)
	}
	return rArray
}
