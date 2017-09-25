package bktree

import (
	"github.com/antzucaro/matchr"
	"strings"
)

type node struct {
	Word     string
	Children map[int]*node
}

func NewNode(word string) *node {
	return &node{
		Word:     strings.ToLower(word),
		Children: nil,
	}
}

func (n *node) AddChild(key int, word string) {
	if n.Children == nil {
		n.Children = make(map[int]*node)
	}
	n.Children[key] = NewNode(word)
}

func (n *node) Keys() []int {
	if n.Children == nil {
		return make([]int, 0)
	}
	var keys []int
	for key := range n.Children {
		keys = append(keys, key)
	}
	return keys
}

func (n *node) Node(key int) *node {
	if n.Children == nil {
		return nil
	}
	return n.Children[key]
}

func (n *node) ContainsKey(key int) bool {
	if n.Children == nil {
		return false
	}
	_, ok := n.Children[key]
	return ok
}

type Tree struct {
	Root *node
	Size int
}

func (tree *Tree) Add(word string) {
	word = strings.ToLower(word)

	if tree.Root == nil {
		tree.Root = NewNode(word)
		tree.Size++
		return
	}

	curNode := tree.Root

	dist := matchr.DamerauLevenshtein(curNode.Word, word)
	for curNode.ContainsKey(dist) {
		if dist == 0 {
			return
		}

		curNode = curNode.Node(dist)
		dist = matchr.DamerauLevenshtein(curNode.Word, word)
	}

	curNode.AddChild(dist, word)
	tree.Size++
}

func (tree *Tree) Search(word string, distance int) []string {
	var rtn = make([]string, 0, tree.Size)
	word = strings.ToLower(word)

	tree.RecursiveSearch(tree.Root, &rtn, word, distance)

	return rtn
}

func (tree *Tree) RecursiveSearch(node *node, rtn *[]string, word string, distance int) {
	curDist := matchr.DamerauLevenshtein(node.Word, word)
	minDist := curDist - distance
	maxDist := curDist + distance

	if curDist <= distance {
		*rtn = append(*rtn, node.Word)
	}

	for _, v := range node.Keys() {
		if minDist <= v && v <= maxDist {
			tree.RecursiveSearch(node.Node(v), rtn, word, distance)
		}
	}
}
