package main

type ListNode struct {
	data int
	next *ListNode
}

type LinkedList struct {
	head *ListNode
}

func (l *LinkedList) Add(data int) {
	node := &ListNode{data: data}
	node.next = node
	l.head = node
}

func (l *LinkedList) Remove(data int) {
	if l.head == nil {
		return
	}

	if l.head.data == data {
		l.head = l.head.next
		return
	}

	current := l.head
	for current.next != nil {
		if current.data == data {
			current.next = current.next.next
			return
		}
		current = current.next
	}
}
