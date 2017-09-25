/*
The MIT License (MIT)
Copyright (c) 2013

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"),
to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense,
and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/

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
	for k := range n.Children {
		if k == key {
			return true
		}

	}
	return false
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
