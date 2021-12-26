package main

import "fmt"

// QUeue holds slice
type Queue struct {
	items []int
}

// Enqueue
func (q *Queue) Enqueue(item int) {
	q.items = append(q.items, item)
}

// Dequeue
func (q *Queue) Dequeue() int {
	toRemove := q.items[0]
	q.items = q.items[1:]
	return toRemove

}

func main() {
	myQueue := Queue{}
	fmt.Println(myQueue)
	myQueue.Enqueue(40)
	myQueue.Enqueue(50)
	myQueue.Enqueue(60)
	myQueue.Enqueue(70)
	fmt.Println(myQueue)
	myQueue.Dequeue()
	fmt.Println(myQueue)
}
