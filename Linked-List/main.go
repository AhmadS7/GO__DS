package main

import "fmt"

type node struct {
	data int
	next *node
}

type linkedList struct {
	head   *node
	length int
}

func (list *linkedList) prepend(new *node) {
	second := list.head
	list.head = new
	list.head.next = second
	list.length++
}

func main() {
	myList := linkedList{}

	node1 := &node{data: 67}
	node2 := &node{data: 498}
	node3 := &node{data: 787}
	node4 := &node{data: 9263}

	myList.prepend(node1)
	myList.prepend(node2)
	myList.prepend(node3)
	myList.prepend(node4)

	fmt.Println(myList)

}
