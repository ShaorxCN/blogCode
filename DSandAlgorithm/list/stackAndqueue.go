package main

import(
	"log"
	"container/list"
)


func main(){
	//stack
	l := list.New()
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)

	log.Println(l)
	stackEle := l.Back()
	l.Remove(stackEle)
	log.Println(stackEle,l)


	//queue
	q := list.New()
	q.PushBack("a")
	q.PushBack("b")
	q.PushBack("c")
	queueEle:=q.Front()
	q.Remove(queueEle)
	log.Println(q,queueEle)



}
