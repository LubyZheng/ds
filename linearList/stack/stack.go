package stack

import (
	"container/list"
	"fmt"
)

type Stack struct {
	line list.List
}

func (stack *Stack) Push(val int) {
	stack.line.PushFront(val)
}

func (stack *Stack) Pop() {
	stack.line.Remove(stack.line.Front())
}

func (stack *Stack) Top() interface{} {
	return stack.line.Front().Value
}

func (stack *Stack) Length() int {
	return stack.line.Len()
}

func (stack *Stack) Clear() {
	stack.line.Init()
}

func (stack Stack) ShowStack() {
	for e := stack.line.Front(); e != nil; e = e.Next() {
		fmt.Printf("%d ", e.Value)
	}
	fmt.Println()
}
