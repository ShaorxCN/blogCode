package tree

//简单实现赫夫曼树

func huffmanDeal(tree []*binaryTree) *binaryTree {
	if len(tree) == 1 {
		return tree[0]
	}

	temp := tree[len(tree)-1]
	temp2 := tree[len(tree)-2]
	tree = tree[:len(tree)-2]
	root := NewNode(temp.root.data.(int) + temp2.root.data.(int))

	t := NewBinaryTree(root)
	t.root.AddLc(temp.root)
	t.root.AddRc(temp2.root)

	tree = append(tree, t)

	return huffmanDeal(tree)

}

func Huffman(w []int) *binaryTree {

	if w == nil || len(w) == 0 {
		return nil
	}

	trees := []*binaryTree{}
	for _, v := range w {
		r := NewNode(v)
		trees = append(trees, NewBinaryTree(r))
	}

	if len(trees) == 1 {
		return trees[0]
	}

	return huffmanDeal(trees)
}
