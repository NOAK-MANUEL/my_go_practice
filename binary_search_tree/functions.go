package main

func (b *BST) Insert(val int) {
	b.Root = InsertNode(b.Root, val)
}
func InsertNode(node *Node, val int) *Node {

	if node == nil {
		return &Node{Value: val}
	}

	if val > node.Value {
		node.Right = InsertNode(node.Right, val)
	}
	if val < node.Value {
		node.Left = InsertNode(node.Left, val)
	}
	return node
}

func (b *BST) Search(val int) bool {
	return SearchNode(b.Root, val)
}

func SearchNode(node *Node, val int) bool {
	if node == nil {
		return false
	}
	if node.Value == val {
		return true
	}

	if val < node.Value {
		return SearchNode(node.Left, val)
	}
	if val > node.Value {
		return SearchNode(node.Right, val)
	}

	return false
}

func (b *BST) DeleteBST() {

}
