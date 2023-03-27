package list

import "fmt"

type node struct {
	value int
	next  *node
	prev  *node
}

type doublyLinkedList struct {
	head *node
	tail *node
	size int
}

func (dll *doublyLinkedList) Init() {
	dll.head = nil
	dll.tail = nil
	dll.size = 0
}

func (dll *doublyLinkedList) add(value int) {
	newNode := &node{value: value}

	if dll.head == nil {
		dll.head = newNode
		dll.tail = newNode
	} else {
		newNode.prev = dll.tail
		dll.tail.next = newNode
		dll.tail = newNode
	}
	dll.size++
}

func (dll *doublyLinkedList) addElementStart(value int) {
	newNode := &node{value: value, prev: nil}
	if dll.head == nil {
		dll.head = newNode
		dll.tail = newNode
	} else {
		dll.head.prev = newNode
		newNode.next = dll.head
		dll.head = newNode
	}
	dll.size++
}

func (dll *doublyLinkedList) addElementOnIndex(value int, index int) {
	newNode := &node{value: value}
	if index < 0 || index > dll.size {
		return
	}
	if index == 0 {
		dll.addElementStart(value)
	} else if index > 0 && index < dll.size/2 {
		link := dll.head
		for i := 0; i < index-1; i++ {
			link = link.next
		}
		newNode.next = link.next
		newNode.prev = link
		link.next = newNode
		link.next.prev = newNode
	} else if index >= dll.size/2 && index < dll.size {
		link := dll.tail
		for i := dll.size; index-1 < i; i-- {
			link = link.prev
		}
		newNode.next = link.next
		newNode.prev = link
		link.next = newNode
		link.next.prev = newNode
	}
	dll.size++
}

func (dll *doublyLinkedList) Remove() {
	if dll.tail == nil {
		return
	}
	dll.tail = dll.tail.prev
	dll.tail.prev.next = nil
	dll.size--
}

func (dll *doublyLinkedList) RemoveOnIndex(index int) {

}

func (dll *doublyLinkedList) Get(index int) {

}

func (dll *doublyLinkedList) Set(value int, index int) {

}

func (dll *doublyLinkedList) Size() int {
	return dll.size
}

func (dll *doublyLinkedList) show() {
	link := dll.head

	for link != nil {
		fmt.Print(link.value, " ")
		link = link.next
	}
	fmt.Println()
}
