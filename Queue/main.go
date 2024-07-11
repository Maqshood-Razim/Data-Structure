package main

import "fmt"

func main() {

	queue := queue{}

	queue.enqueue(1)
	queue.enqueue(2)
	queue.enqueue(3)
	queue.enqueue(4)
	queue.enqueue(5)
	queue.enqueue(6)

	fmt.Println(queue.item)

	for {

		item, err := queue.dequeue()

		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println("dequeued item ->", item)
		fmt.Println("current item->", queue.item)

	}



}

type queue struct {
	item []int
}

func (q *queue) enqueue(data int) {
	q.item = append(q.item, data)
}

func (q *queue) dequeue() (int, error) {

	if len(q.item) == 0 {
		return 0, fmt.Errorf("queue is empty")
	}

	front := q.item[0]
	q.item = q.item[1:]

	return front, nil

}

