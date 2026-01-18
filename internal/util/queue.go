package util

import (
	"container/list"
)

type Queue struct {
	q *list.List
}

func NewQueue() *Queue {
	return &Queue{list.New()}
}

func (oq *Queue) Push(order interface{}) {
	oq.q.PushBack(order)
}

func (oq *Queue) Pop() interface{} {
	front := oq.q.Front()
	if front == nil {
		return nil
	}

	return oq.q.Remove(front)
}

func (oq *Queue) Len() int {
	return oq.q.Len()
}

func (oq *Queue) ForEach(fn func(v interface{})) {
	for e := oq.q.Front(); e != nil; e = e.Next() {
		fn(e.Value)
	}
}
