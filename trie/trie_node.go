package trie

import (
    "path/filepath"
    "strings"
    "bytes"
)

// Node trie tree node container carries name and children.
type Node struct {
    isLeaf bool
    name  string
    sep    rune
    children  map[string]*Node
}


// NewNode
func NewNode(isLeaf bool, name string, sep rune) *Node {
    return &Node{
        isLeaf: isLeaf,
        name: name,
        sep: sep,
        children: make(map[string]*Node),
    }
}


// Get - get a child node.
func (n *Node) Get(childName string) {
    if child, ok := n.children[childName]; ok {
        return child
    } else {
        return nil
    }
}

// Insert - insert a child node.
func (n *Node) Insert(childNode *Node) {
    if child, ok := n.children[childNode.name]; ok {
        if child.isLeaf 
    } else {
        n.children[childNode.name] = childNode
    }
}

// Delete - delete a child_node.
func (n *Node) Delete(childName string) bool {
    if _, ok := n.children[childName]; ok {
        delete(n.children, childName)
        return true
    } else {
        return false
    }
}

// IsLeaf - check if node is leaf node.
func (n *Node) IsLeaf() bool {
    return n.isLeaf
}

// GetAllNode - get all child nodes based on wild card query.
func (n *Node) GetAllNode(pattern string) []*Node {
    // TODO: Add expand braces logic
    var matches []*Node

    for childName, childNode := range n.children {
        matched, err := Match(pattern, childName)
        if err != nil {
            // logging
            continue
        }

        if matched {
            append(matches, n.children[childName])
        }
    }
    return matches
}

// ExpandQuery - expand a wildcard query.
func (n *Node) ExpandQuery(query string) []string{
    var queries []string
    sepIndex := strings.IndexRune(query, n.sep)

    if sepIndex < 0 {
        for node := range n.GetAllNode(query) {
            if node.isLeaf {
                append(queries, node.name)
            }
        }
    } else {
        curPart := query[:sepIndex]
        curMatches := n.GetAllNode(curPart)
        subQuery := query[sepIndex + 1]
        for m := range curMatches {
            subQueries := m.ExpandQuery(subQuery)
            for sq := range subQueries {
                var buffer bytes.Buffer
                buffer.WriteString(m.name)
                buffer.WriteRune(n.sep)
                buffer.WriteString(sq)
                append(queries, buffer.String())
            }   
        }
    }
    return queries
}

// Count - return children count
func (n *Node) Count() int {
    return len(n.children)
}
