package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkList struct {
	head *Node
	tail *Node
	len  int
}

func (l *LinkList) Insert(datta int) {
	newNode := Node{}
	newNode.data = datta

	if l.len == 0 {
		l.head = &newNode
		l.tail = l.head
		l.len++
		return
	}
	ptr := l.head

	for ; ptr.next != nil; ptr = ptr.next {
	}

	ptr.next = &newNode
	l.len++
}

func (l *LinkList) PrintLL() {
	if l.len == 0 {
		fmt.Println("LinkList is empty")
		return
	}
	ptr := l.head

	for ; ptr != nil; ptr = ptr.next {
		fmt.Printf("%+v -> ", ptr.data)
	}
	fmt.Println("END")
}

func main() {
	head := LinkList{}
	head.Insert(34)
	head.Insert(36)
	head.Insert(4)
	head.Insert(8)

	head.PrintLL()

}
