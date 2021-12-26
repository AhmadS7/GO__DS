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

func (list linkedList) printThisData() {
	nodeToPrint := list.head

	for list.length != 0 {
		fmt.Printf("%d  ", nodeToPrint.data)
		nodeToPrint = nodeToPrint.next
		list.length--
	}

	fmt.Printf("\n")
}

func (list *linkedList) deleteWithValue(value int) {
	if list.length == 0 {
		return
	}

	if list.head.data == value {
		list.head = list.head.next
		list.length--
		return
	}
	previousToDelete := list.head
	for previousToDelete.next.data != value {
		if previousToDelete.next.next == nil {
			return
		}
		previousToDelete = previousToDelete.next
	}
	previousToDelete.next = previousToDelete.next.next
	list.length--
}

func main() {
	myList := linkedList{}

	node1 := &node{data: 67}
	node2 := &node{data: 498}
	node3 := &node{data: 787}
	node4 := &node{data: 9263}
	node5 := &node{data: 5456}
	node6 := &node{data: 7889}

	myList.prepend(node6)
	myList.prepend(node1)
	myList.prepend(node2)
	myList.prepend(node5)
	myList.prepend(node3)
	myList.prepend(node4)

	myList.deleteWithValue(500)
	myList.printThisData()
	myList.deleteWithValue(9263)
	myList.printThisData()
	myList.deleteWithValue(67)

	emptyList := linkedList{}
	emptyList.deleteWithValue(10)

}
