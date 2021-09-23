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
	li ListItems
}

func NewList() List {
	return &list{
		li: ListItems{},
	}
}

func (l *list) Len() int {
	return len(l.li)
}

func (l *list) Front() *ListItem {
	for i, item := range l.li {
		if item.Prev == nil {
			fmt.Println("Front", l.li[i].Value)
			return &(l.li[i])
		}
	}
	return nil
}

func (l *list) Back() *ListItem {
	for i, item := range l.li {
		if item.Next == nil {
			fmt.Println("Back", l.li[i].Value)
			return &(l.li[i])
		}
	}
	return nil
}

func (l *list) PushFront(v interface{}) *ListItem {
	fmt.Println("PushFront")
	l.PrintColletion()

	front := l.Front()
	li := ListItem{
		Value: v,
		Next:  front,
		Prev:  nil,
	}
	l.li = append(l.li, li)

	if front != nil {
		l.setPrev(front.Value, &li)
	}

	l.PrintColletion()
	fmt.Printf("list - %v\n", l)
	return &li
}

func (l *list) setNext(v interface{}, next *ListItem) {
	for i, item := range l.li {
		if item.Value == v {
			l.li[i].Next = next
		}
	}
}

func (l *list) setPrev(v interface{}, prev *ListItem) {
	for i, item := range l.li {
		if item.Value == v {
			fmt.Printf("before %v\n", l.li[i])
			l.li[i].Prev = prev
			fmt.Printf("after %v\n", l.li[i])
		}
	}
}

func (l *list) PushBack(v interface{}) *ListItem {
	fmt.Println("PushBack")
	l.PrintColletion()

	back := l.Back()

	li := ListItem{
		Value: v,
		Next:  nil,
		Prev:  l.Back(),
	}

	l.li = append(l.li, li)

	if back != nil {
		l.setNext(back.Value, &(l.li[l.Len()-1]))
		fmt.Printf("back - %v\n", l.Back())
		fmt.Printf("front - %v\n", l.Front())
	}

	l.PrintColletion()
	fmt.Printf("list after pushback - %v\n", l)
	return &li
}

func (l *list) Remove(i *ListItem) {
	fmt.Println("Remove")
	l.PrintColletion()

	for j, item := range l.li {
		if &item == i {
			if l.li[j].Next != nil {
				l.setNext(l.li[j].Next.Value, item.Next)
			}

			if l.li[j].Prev != nil {
				l.setPrev(l.li[j].Prev.Value, item.Prev)
			}

			//copy(l.li[j:], l.li[j+1:])   // Shift a[i+1:] left one index.
			//l.li[l.Len()-1] = ListItem{} // Erase last element (write zero value).
			l.li = append(l.li[j:], l.li[j+1:]...)
			break
		}
	}
	l.PrintColletion()
}

func (l *list) MoveToFront(i *ListItem) {
	if i.Prev != nil {
		i.Prev.Next = i.Next
	}
	if i.Next != nil {
		i.Next.Prev = i.Prev
	}
	front := l.Front()
	i.Next = front
	i.Prev = nil
	front.Prev = i
	l.PrintColletion()
}

func (l *list) PrintColletion() {
	elems := make([]interface{}, 0, l.Len())
	for i := l.Front(); i != nil; i = i.Next {
		fmt.Printf("next value: %v, next: %v, prev: %v\n", i.Value, i.Next, i.Prev)
		elems = append(elems, i.Value.(int))
	}
	fmt.Println(elems)
}
