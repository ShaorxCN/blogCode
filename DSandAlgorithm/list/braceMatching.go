package main

import (
	"container/list"
	"log"
)

func braceMatching(l string) bool {
	stack := list.New()
	for _, v := range l {
		top := stack.Back()

		if top == nil {
			stack.PushBack(v)
			continue
		}

		switch top.Value {

		case '(':
			if v == ')' {
				stack.Remove(top)
			} else {
				stack.PushBack(v)
			}

		case '[':
			if v == ']' {
				stack.Remove(top)
			} else {
				stack.PushBack(v)
			}

		default:
			stack.PushBack(v)
		}
	}

	return stack.Len() == 0
}

func main() {
	l := "[()[]()()]"

	log.Println(braceMatching(l))

	l2 := "[()[(][([[)]]]"

	log.Println(braceMatching(l2))

}
