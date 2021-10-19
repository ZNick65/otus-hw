package hw04lrucache

import "fmt"

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
	PrintColletion()
}

type (
	ListItems []ListItem
	ListItem  struct {
		Value interface{}
		Next  *ListItem
		Prev  *ListItem
	}
)

type list struct {
	// Place your code here.
	firstNode *ListItem
	lastNode  *ListItem
	count     int
}

func NewList() List {
	return &list{}
}

func (l *list) Len() int {
	return l.count
}

func (l *list) Front() *ListItem {
	return l.firstNode
}

func (l *list) Back() *ListItem {
	return l.lastNode
}

func (l *list) PushFront(v interface{}) *ListItem {
	l.count++
	front := l.Front()

	li := &ListItem{
		Value: v,
		Next:  front,
	}

	if front != nil {
		front.Prev = li
	}

	l.firstNode = li

	if l.Back() == nil {
		l.lastNode = li
	}

	return li
}

func (l *list) PushBack(v interface{}) *ListItem {
	l.count++
	back := l.Back()

	li := &ListItem{
		Value: v,
		Prev:  back,
	}

	if back != nil {
		back.Next = li
	}

	l.lastNode = li

	if l.Front() == nil {
		l.firstNode = li
	}

	return li
}

func (l *list) Remove(i *ListItem) {
	if i.Prev != nil {
		i.Prev.Next = i.Next
	}

	if i.Next != nil {
		i.Next.Prev = i.Prev
	}

	if l.firstNode == i {
		l.firstNode = i.Next
	}

	if l.lastNode == i {
		l.lastNode = i.Prev
	}

	l.count--
}

func (l *list) MoveToFront(i *ListItem) {
	if l.firstNode == i {
		return
	}

	if l.lastNode == i {
		l.lastNode = i.Prev
	}

	if i.Prev != nil {
		i.Prev.Next = i.Next
	}

	if i.Next != nil {
		i.Next.Prev = i.Prev
	}

	l.firstNode.Prev = i
	i.Prev = nil
	i.Next = l.firstNode
	l.firstNode = i
}

func (l list) PrintColletion() {
	elems := make([]interface{}, 0, l.Len())
	for i := l.Front(); i != nil; i = i.Next {
		// fmt.Printf("current value: %v, next: %v, prev: %v\n", i, i.Next, i.Prev)
		elems = append(elems, i.Value)
	}
	fmt.Println(elems)
}
