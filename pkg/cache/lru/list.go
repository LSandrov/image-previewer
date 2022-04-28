package lru

type List interface {
	Len() int
	Front() *Item
	Back() *Item
	PushFront(v interface{}) *Item
	PushBack(v interface{}) *Item
	Remove(i *Item)
	MoveToFront(i *Item)
}

type list struct {
	len         int
	front, back *Item
}

type Item struct {
	Value interface{}
	Next  *Item
	Prev  *Item
}

func (l *list) Len() int {
	return l.len
}

func (l *list) Front() *Item {
	return l.front
}

func (l *list) Back() *Item {
	return l.back
}

func (l *list) PushFront(v interface{}) *Item {
	listItem := &Item{Value: v, Next: l.front}

	if l.front == nil {
		l.back = listItem
	} else {
		l.front.Prev = listItem
	}

	l.front = listItem
	l.len++

	return listItem
}

func (l *list) PushBack(v interface{}) *Item {
	listItem := &Item{Value: v, Prev: l.back}

	if l.back == nil {
		l.back = listItem
	} else {
		l.back.Next = listItem
	}

	l.back = listItem
	l.len++

	return listItem
}

func (l *list) Remove(i *Item) {
	l.pos(i)

	l.len--
}

func (l *list) MoveToFront(i *Item) {
	if l.front == i {
		return
	}
	if l.back == i {
		l.back = i.Prev
		l.back.Next = nil
	} else {
		l.pos(i)
	}

	currentFront := l.front

	l.front = i
	l.front.Prev = nil
	l.front.Next = currentFront
	l.front.Next.Prev = i
}

func NewList() List {
	return new(list)
}

func (l *list) pos(i *Item) {
	if i.Next != nil {
		i.Next.Prev = i.Prev
	} else {
		l.back = i.Prev
	}

	if i.Prev != nil {
		i.Prev.Next = i.Next
	} else {
		l.front = i.Next
	}
}
