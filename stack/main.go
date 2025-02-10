package main

import "log"

func main() {
	s := Stack{}
	s.Push(2)
	s.Push(4)
	s.Push(5)
	s.Display()

	s.Pop()
	s.Display()

}

type Stack struct {
	content []int
}

func (stack *Stack) Display() {
	log.Print(stack.content)
}

func (stack *Stack) Push(item int) {
	stack.content = append(stack.content, item)
}

func (stack *Stack) Pop() int {
	lastItemIndex := len(stack.content) - 1
	lastItem := stack.content[lastItemIndex]
	stack.content = stack.content[:lastItemIndex]
	return lastItem
}
