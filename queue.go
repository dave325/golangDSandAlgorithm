package main

type Queue struct {
	Items []int
	Size  int
	Max   int
}

func (q *Queue) Dequeue() int {
	if len(q.Items) == 0 {
		return -1
	}
	item := q.Items[0]
	q.Items = q.Items[1:]
	q.Size--
	return item
}

func (q *Queue) Enqueue(item int) {
	if q.Size == q.Max {
		tempArr := make([]int, 0, 10)
		q.Items = append(q.Items, tempArr...)
		q.Max += 10
	}
	q.Items = append(q.Items, item)
	q.Size++
}

func (q *Queue) IsEmpty() bool {
	return q.Size == 0
}

func (q *Queue) IsFull() bool {
	return q.Size == q.Max
}

func (q *Queue) Peek() interface{} {
	if q.Size == 0 {
		return nil
	}
	item := q.Items[0]
	return item
}
