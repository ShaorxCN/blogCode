package tree

import (
	"container/list"

	"bytes"
	"fmt"

	"log"
)

//由于修改父节点以及更改子节点影响较大这里主要写下遍历
type Node struct {
	data   interface{}
	parent *Node
	leftC  *Node
	rightN *Node
	height int //该节点为root的子树的深度
	size   int //该子树的节点总数
}

type binaryTree struct {
	root   *Node
	height int
	size   int
}

func NewNode(data interface{}) *Node {
	return &Node{data: data, size: 1}
}

func (n *Node) GetData() interface{} {
	if n == nil {
		return nil
	}

	return n.data
}

func (n *Node) SetData(data interface{}) *Node {
	n.data = data
	return n
}

func (n *Node) HasParent() bool {
	if n == nil {
		return false
	}
	return n.parent != nil
}

func (n *Node) GetParent() *Node {
	if !n.HasParent() {
		return nil
	}

	return n.parent
}

func (n *Node) AddLc(i *Node) *Node {
	if n == nil {
		return nil
	}

	n.leftC = i
	i.parent = n
	n.height++
	n.size++

	return n

}

func (n *Node) AddRc(i *Node) *Node {
	if n == nil {
		return nil
	}

	n.rightN = i
	i.parent = n
	n.height++
	n.size++

	return n

}

func (n *Node) GetLC() *Node {
	if n == nil || n.leftC == nil {
		return nil
	}
	return n.leftC
}

func (n *Node) GetRC() *Node {
	if n == nil || n.rightN == nil {
		return nil
	}
	return n.rightN
}

func NewBinaryTree(root *Node) *binaryTree {
	return &binaryTree{root: root}
}

func preOrder(n *Node, l *list.List) {
	if n == nil {
		return
	}

	l.PushBack(n)
	preOrder(n.GetLC(), l)
	preOrder(n.GetRC(), l)
}
func (b *binaryTree) PreOrder() *list.List {
	l := list.New()
	preOrder(b.root, l)
	return l

}

func inOrder(n *Node, l *list.List) {
	if n == nil {
		return
	}

	inOrder(n.GetLC(), l)
	l.PushBack(n)
	inOrder(n.GetRC(), l)

}

func (b *binaryTree) InOrder() *list.List {
	l := list.New()
	inOrder(b.root, l)
	return l
}

func postOrder(n *Node, l *list.List) {
	if n == nil {
		return
	}

	postOrder(n.GetLC(), l)
	postOrder(n.GetRC(), l)
	l.PushBack(n)
}

func (b *binaryTree) PostOrder() *list.List {
	l := list.New()
	postOrder(b.root, l)
	return l
}

func pt(l *list.List) (res []int) {
	if l == nil {
		return nil
	}

	for l.Front() != nil {
		n := l.Front()
		res = append(res, n.Value.(*Node).data.(int))
		l.Remove(n)
	}

	return res

}

func toptl(node *Node, buffer *bytes.Buffer) {
	if node == nil {
		return
	}

	if node.GetLC() != nil {
		buffer.WriteString(fmt.Sprintf("%d leftc is %d \n", node.data.(int), node.GetLC().data.(int)))

	}

	if node.GetRC() != nil {
		buffer.WriteString(fmt.Sprintf("%d rightc is %d \n", node.data.(int), node.GetRC().data.(int)))
	}

	toptl(node.GetLC(), buffer)

	toptl(node.GetRC(), buffer)
}

func ptl(t *binaryTree) string {
	b := new(bytes.Buffer)

	r := t.root

	toptl(r, b)
	return b.String()
}
