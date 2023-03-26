package list

import (
	"errors"
	"fmt"
)

type Node struct {
	value int
	next  *Node
}

type LinkedList struct {
	head *Node
	size int
}

func (ll *LinkedList) Init() {
	ll.head = nil
	ll.size = 0
}

func (ll *LinkedList) addNode(value int) {
	newNode := &Node{value: value}

	/*
		newNode := &Node{}
		newNode.value = value
	*/

	if ll.head == nil {
		ll.head = newNode
	} else {
		link := ll.head
		for link.next != nil {
			link = link.next
		}
		link.next = newNode
	}
}

func (ll *LinkedList) addElementStart(value int) {
	if ll.head == nil {
		ll.head = &Node{value: value}
	} else {
		newNode := &Node{value: value, next: ll.head}
		ll.head = newNode
	}
	ll.size++
}

func (ll *LinkedList) addElementOnIndex(index int, value int) {
	newNode := &Node{value, nil}
	if index < 0 || index > ll.size {

	} else if index == 0 {
		ll.addElementStart(value)
	} else {
		node := ll.head
		for i := 0; i < index-1; i++ {
			node = node.next
		}
		newNode.next = node.next
		node.next = newNode
	}
	ll.size++
}

func (ll *LinkedList) RemoveElement() {
	if ll.size > 0 {
		ll.size--
	}
}

func (ll *LinkedList) RemoveElementOnIndex(index int) {
	link := ll.head
	if index > ll.size || index < 0 || link == nil || link.next == nil {
		return
	} else if index == 0 {
		ll.head = ll.head.next
	} else {
		for i := 0; i < index-1; i++ {
			link = link.next
		}
		link.next = link.next.next
	}
	ll.size--
}

func (ll *LinkedList) Get(index int) (int, error) {
	if index < 0 || index > ll.size {
		return 0, errors.New("Indice Iválido")
	} else {
		value := ll.head
		if value == nil {
			return 0, errors.New("Indice Iválido")
		}
		for i := 0; i < index; i++ {
			value = value.next
		}
		return value.value, nil
	}
}

func (ll *LinkedList) SetElement(value int, index int) {
	if index < 0 || index > ll.size {
		return
	} else {
		element := ll.head
		for i := 0; i < index; i++ {
			element = element.next
		}
		element.value = value
	}
}

func (ll *LinkedList) Sizell() int {
	return ll.size
}

func (ll *LinkedList) show() {
	link := ll.head
	for link != nil {
		fmt.Println(link.value)
		link = link.next
	}
}
