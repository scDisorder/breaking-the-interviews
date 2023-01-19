package main

type DoublyListNode struct {
	data int
	next *DoublyListNode
	prev *DoublyListNode
}

type DoublyLinkedList struct {
	head *DoublyListNode
	tail *DoublyListNode
}

func (l *DoublyLinkedList) AddToHead(data int) {
	node := &DoublyListNode{data: data}
	if l.head == nil {
		l.head = node
		l.tail = node
	} else {
		l.head.prev = node
		node.next = l.head
		l.head = node
	}
}

func (l *DoublyLinkedList) AddToTail(data int) {
	node := &DoublyListNode{data: data}
	if l.tail == nil {
		l.tail = node
		l.head = node
	} else {
		l.tail.next = node
		node.prev = l.head
		l.tail = node
	}
}

func (l *DoublyLinkedList) Remove(data int) {
	current := l.head
	for current.next != nil {
		if current.data == data {
			if current.prev != nil {
				current.prev.next = current.next
			} else {
				l.head = current.next
			}
			if current.next != nil {
				current.next.prev = current.prev
			} else {
				l.tail = current.prev
			}
		}
		current = current.next
	}
}

func (l *DoublyLinkedList) Find(data int) *DoublyListNode {
	current := l.head
	for current.next != nil {
		if current.data == data {
			return current
		}
		current = current.next
	}
	return nil
}
