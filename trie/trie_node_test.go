package trie

import (
	"testing"
)

// Simply make sure creating a new tree works.
func TestNewTrieNode(t *testing.T) {
	var node *Node
	node = NewNode(true, "carbon", '.')

	if node.Count() != 0 {
		t.Errorf("expected size 0, got: %d", node.Count())
	}

	if node.isLeaf != true {
		t.Errorf("expected isLeaf true, got: %t", node.isLeaf)
	}

	if node.name != "carbon" {
		t.Errorf("expected name carbon, got: %s", node.name)
	}
}
