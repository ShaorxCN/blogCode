package tree

import (
	"testing"
)

func TestBuffman(t *testing.T){
	//这里为了方便处理直接用有序slice。
	weights := []int{8,7,6,5,4,3,2,1}
	t.Log(ptl(Huffman(weights)))
}

func TestBinaryTree(t *testing.T) {
	root := NewNode(0)
		tree := NewBinaryTree(root)

		n1 := NewNode(1)
		n2 := NewNode(2)
		n3 := NewNode(3)
		n4 := NewNode(4)
		n5 := NewNode(5)
		n6 := NewNode(6)
		n7 := NewNode(7)
		n8 := NewNode(8)
		n9 := NewNode(9)
		n10 := NewNode(10)
		n11 := NewNode(11)
		n13 := NewNode(13)
		root.AddLc(n1).AddRc(n2)
		n1.AddLc(n3)
		n2.AddLc(n4).AddRc(n5)

		n3.AddLc(n6).AddRc(n7)
		n4.AddRc(n8)
		n5.AddLc(n9).AddRc(n10)
		n1.AddRc(n11)
		n11.AddRc(n13)

		t.Log(pt(tree.PreOrder()))
		t.Log(pt(tree.InOrder()))
		t.Log(pt(tree.PostOrder()))
}