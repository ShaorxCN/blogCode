package main

import (
	"log"
	"math"
)

//链表实现多项式的求和.本来是做多项式求和的，这里写的是指数式求和

type PNode struct{
	coef float64
	expn float64
	belong *PolymomialList
	next *PNode
}


type PolymomialList struct{
	root PNode
	length int
	last *PNode
}


func (l *PolymomialList)Init() *PolymomialList{
	l.length = 0
	l.root.next = nil
	l.last = &l.root
	return l
}

func NewPolynomialList() *PolymomialList{
	return new(PolymomialList).Init()

}

func NewPNode(coef float64,expn float64)*PNode{
	return &PNode{coef:coef,expn:expn}
}


func (l *PolymomialList)AddNode(n *PNode)*PolymomialList{

	l.last.next = n
	n.belong = l
	l.length++
	l.last = n
	return l
}


func (l *PolymomialList)Sum() float64{
	sum := 0.0

	if l.length == 0{
		return 0
	}


	n := l.root.next


	for n != nil{

		sum += math.Pow(n.coef,n.expn)
		n = n.next

	}

	return sum
}


func main(){
	l := NewPolynomialList()

	pa := NewPNode(3,2)
	pb := NewPNode(3,3)
	pc := NewPNode(2,4)
	pd := NewPNode(1,1)

	l.AddNode(pa).AddNode(pb).AddNode(pc).AddNode(pd)

	log.Println(l.Sum())
}


