package queue

import (
	"container/list"
	"fmt"
)

type Queue struct {
	line list.List
}

func (queue *Queue) Add(val int) {
	queue.line.PushBack(val)
}

func (queue Queue) Length() int {
	return queue.line.Len()
}

func (queue *Queue) Quit() interface{} {
	return queue.line.Remove(queue.line.Front())
}

func (queue *Queue) Clear(){
	queue.line.Init()
}

func (queue Queue) ShowQueue() {
	for e := queue.line.Front(); e != nil; e = e.Next() {
		fmt.Printf("%d ", e.Value)
	}
	fmt.Println()
}
