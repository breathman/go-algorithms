package queue

import "errors"

type Queue struct {
	elements []int
}

func New() *Queue {
	return &Queue{}
}

func (q *Queue) Push(item int) {
	q.elements = append(q.elements, item)
}

func (q *Queue) Pop() (int, error) {
	if !q.IsEmpty() {
		res := q.elements[0]
		q.elements = q.elements[1:]
		return res, nil
	}
	return 0, errors.New("queue is empty")
}

func (q *Queue) Top() (int, error) {
	if !q.IsEmpty() {
		return q.elements[0], nil
	}
	return 0, errors.New("queue is empty")
}

func (q *Queue) Size() int {
	return len(q.elements)
}

func (q *Queue) IsEmpty() bool {
	return q.Size() == 0
}
