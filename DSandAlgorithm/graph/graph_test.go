package graph

import (
	"testing"
	"fmt"
)

func TestDfs(t *testing.T) {
	s := []string{"0", "1", "2", "3", "4", "5", "6"}
	l := NewAdjListGraph(s)

	l.hNodes[0].AddVNode(CreateVNode("1")).(*VNode).AddVVode(CreateVNode("3")).AddVVode(CreateVNode("2"))
	l.hNodes[1].AddVNode(CreateVNode("2")).(*VNode).AddVVode(CreateVNode("6")).AddVVode(CreateVNode("0"))
	l.hNodes[2].AddVNode(CreateVNode("1")).(*VNode).AddVVode(CreateVNode("0"))
	l.hNodes[3].AddVNode(CreateVNode("0"))
	l.hNodes[4].AddVNode(CreateVNode("5"))
	l.hNodes[5].AddVNode(CreateVNode("4"))
	l.hNodes[6].AddVNode(CreateVNode("1"))
	l.PDFS(0)
	fmt.Println()
	l.PBFS(0)
}
