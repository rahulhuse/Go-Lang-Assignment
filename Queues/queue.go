package main

import "fmt"

type Queue struct {
	items []int
}

func (q *Queue) Enqueue(i int) {

	q.items = append(q.items, i)
}

func (q *Queue) Dequeue() int {
	removedItems := q.items[0]
	q.items = q.items[1:]
	return removedItems
}

func main() {

	myQueue := Queue{}
	myQueue.Enqueue(70)
	myQueue.Enqueue(80)
	myQueue.Enqueue(90)
	fmt.Println(myQueue)
	myQueue.Dequeue()
	fmt.Println(myQueue)

}
