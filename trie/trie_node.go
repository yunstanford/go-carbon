package trie

// Node trie tree node container carries value and children.
type Node struct {
	isLeaf bool
	value  string
	sep    string
	child  map[string]*Node
}


// Insert - insert a child node.
func (n *Node) Insert(child_node *Node) {

}

// Delete - delete a child_node.
func (n *Node) Delete(child_name string) bool {

}

// IsLeaf - check if node is leaf node.
func (n *Node) IsLeaf() bool {
	return n.isLeaf
}

// GetAllNode - get all child nodes based on wild card query.
func (n *Node) GetAllNode(pattern string) []*Node{

}
