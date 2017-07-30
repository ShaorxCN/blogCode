package graph

import (
	"fmt"
	"log"
	"strconv"
)

type VNode struct {
	adjVex  int         //该点在图中的位置
	nextarc *VNode      //下一个条边/弧的节点
	info    interface{} //相关信息，比如权值
}

type HNode struct {
	data     string
	firstarc *VNode
}

type adjListGraph struct {
	hNodes []*HNode
	num    int
}

//初始化邻接表,存储顶点信息，这里就用string
func NewAdjListGraph(nodes []string) adjListGraph {
	hNodes := make([]*HNode, 0)
	for _, v := range nodes {
		hNode := &HNode{data: v}
		hNodes = append(hNodes, hNode)
	}
	return adjListGraph{hNodes: hNodes, num: len(nodes)}
}

//提供VNODE
func CreateVNode(s string) *VNode {
	n, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return &VNode{adjVex: n, info: s}
}

//为头结点添加第一个邻接点,成功返回*VNode，失败返回HNode
func (h *HNode) AddVNode(v *VNode) interface{} {

	if h.firstarc != nil {
		log.Println("error:this hNode already has a VNode")
		return h
	}

	h.firstarc = v
	return v
}

//为vNode增加新的VNode，返回追加的VNode
func (v *VNode) AddVVode(u *VNode) *VNode {
	if v == nil {
		log.Fatal("VNode can not be nil")
	}

	v.nextarc = u
	return u
}

func DFS(l *adjListGraph, start int, visited []int) {

	if visited[start] == 0 {
		visited[start] = 1
		fmt.Print(l.hNodes[start].data)
		n := l.hNodes[start].firstarc
		for n != nil {
			if visited[n.adjVex] == 0 {
				DFS(l, n.adjVex, visited)
			}
			n = n.nextarc
		}
	}

}

//针对非连通图情况所以遍历
func (l *adjListGraph) PDFS(i int) {
	visited := make([]int, l.num)

	for i := 0; i < l.num; i++ {
		DFS(l, i, visited)
	}
}

func BFS(l *adjListGraph, start int, hvisited, visited []int) {

	if hvisited[start] == 0 {
		hvisited[start] = 1

		if visited[start] == 0 {
			visited[start] = 1
			fmt.Print(l.hNodes[start].data)

		}

		n := l.hNodes[start].firstarc
		m := l.hNodes[start].firstarc
		for n != nil {
			if visited[n.adjVex] == 0 {
				visited[n.adjVex] = 1
				fmt.Print(n.info)

			}
			n = n.nextarc
		}

		for m != nil {
			BFS(l, m.adjVex, hvisited, visited)
			m = m.nextarc
		}

	}
}

func (l *adjListGraph) PBFS(i int) {
	visited := make([]int, l.num)
	hvisited := make([]int, l.num)
	for i := 0; i < l.num; i++ {
		BFS(l, i, hvisited, visited)
	}
}

func (l *adjListGraph) Print() {
	for _, v := range l.hNodes {
		fmt.Print(v.data)
		n := v.firstarc
		if n != nil {
			fmt.Print(",", n.info)
			m := n.nextarc
			for m != nil {
				fmt.Print(",", m.info)
				m = m.nextarc
			}
		}
		fmt.Println()
	}

}
