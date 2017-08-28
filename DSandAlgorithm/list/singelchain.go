package main

import (
	"log"
)

//node定义
type Node struct {
	Data   interface{}  //数据域
	next   *Node        //指针域
	belong *SingelChain //Node归属链

}

type SingelChain struct {
	root   Node //存储头结点
	length int  //存储当前长度
}

func (l *SingelChain) Init() *SingelChain {
	l.root.next = nil
	l.length = 0

	return l
}

//返回链表的长度
func (l *SingelChain) Len() int {
	return l.length
}

//返回第一个节点
func (l *SingelChain) Front() *Node {
	if l.length == 0 {
		return nil
	}

	return l.root.next
}

//返回最后一个节点
func (l *SingelChain) Back() *Node {
	if l.length == 0 {
		return nil
	}

	n := l.root.next

	for n.next != nil {
		n = n.next
	}

	return n
}

//在at节点后面插入e,并返回e
func (l *SingelChain) insert(e, at *Node) *Node {
	n := at.next
	at.next = e
	e.next = n
	l.length++
	e.belong = l
	return e
}

func (l *SingelChain) Insert(e, at *Node) *Node {
	return l.insert(e, at)
}

func (l *SingelChain) InsertValue(v interface{}, at *Node) *Node {
	return l.insert(&Node{Data: v}, at)
}

//去除l中的e节点
func (l *SingelChain) remove(e *Node) interface{} {
	n := e.next
	temp := &l.root
	for temp.next != e {
		temp = temp.next
	}
	temp.next = n
	l.length--
	e.belong = nil
	return e.Data

}

func (l *SingelChain) Remove(e *Node) interface{} {
	//e.gelong == l: l must have been initialized when e was inserted
	//and root.belong is nil.if e == nil
	if e.belong == l {
		return l.remove(e)
	}

	return nil
}
