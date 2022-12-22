package main

import (
	"fmt"
	"math"
	"strings"
)

type RadixTreeNode struct {
	key      string
	children []*RadixTreeNode
}

func NewRadixTreeNode() *RadixTreeNode {
	return &RadixTreeNode{
		key:      "",
		children: make([]*RadixTreeNode, 0),
	}
}

func (n *RadixTreeNode) Insert(word string) {
	children := n.children
	for _, child := range children {

		generalStart := getGeneralStart(word, child.key)
		lenGeneralStartn := len(generalStart)

		if child.key == generalStart && child.key != "" {
			child.Insert(word[lenGeneralStartn:])
			return
		}

		if lenGeneralStartn != 0 {
			child.divideWord(lenGeneralStartn)
			child.Insert(word[lenGeneralStartn:])
			return
		}
	}

	n.children = append(n.children, &RadixTreeNode{
		key:      word,
		children: make([]*RadixTreeNode, 0),
	})
}

func getGeneralStart(s1, s2 string) string {
	minLength := int(math.Min(float64(len(s1)), float64(len(s2))))
	var i int
	for i = 0; i < minLength; i += 1 {
		if s1[i] != s2[i] {
			return s1[:i]
		}
	}
	if i == minLength {
		return s1[:i]
	}

	return ""
}

func (n *RadixTreeNode) divideWord(amountLetters int) {

	key := n.key
	children := n.children
	n.key = key[:amountLetters]

	n.children = nil
	n.children = append(n.children, &RadixTreeNode{
		key:      key[amountLetters:],
		children: children,
	})
}

func (n *RadixTreeNode) Traverse(tab int) {
	if n.key == "" && tab != 0 || len(n.children) == 0 {
		return
	}

	for _, child := range n.children {
		if child.key == "" {
			continue
		}

		fmt.Println(strings.Repeat(" ", tab+2) + "*" + child.key)
		child.Traverse(tab + 4)
	}
}

func TestRadixTree() {
	tree := NewRadixTreeNode()

	tree.Insert("home")
	tree.Traverse(0)
	fmt.Println("#############")

	tree.Insert("homeless")
	tree.Traverse(0)
	fmt.Println("#############")

	tree.Insert("hell")
	tree.Traverse(0)
	fmt.Println("#############")

	tree.Insert("hell")
	tree.Traverse(0)
	fmt.Println("#############")

	tree.Insert("hope")
	tree.Traverse(0)
	fmt.Println("#############")

	tree.Insert("meal")
	tree.Traverse(0)
	fmt.Println("#############")

	tree.Insert("burger")
	tree.Traverse(0)
	fmt.Println("#############")

	tree.Insert("sun")
	tree.Traverse(0)
	fmt.Println("#############")

	tree.Insert("shine")
	tree.Traverse(0)
	fmt.Println("#############")

	tree.Insert("kill")
	tree.Traverse(0)
	fmt.Println("#############")

	tree.Insert("bill")
	tree.Traverse(0)
	fmt.Println("#############")
}

func ShowRadixTree() {
	tree := NewRadixTreeNode()

	tree.Insert("hello")
	tree.Insert("help")
	tree.Insert("hell")
	tree.Insert("henk")
	tree.Insert("hellicopter")
	tree.Insert("hellicoptar")
	tree.Insert("hellic")
	tree.Insert("hank")
	tree.Insert("find")
	tree.Insert("found")
	tree.Insert("something")
	tree.Insert("someone")
	tree.Insert("somewho")
	tree.Insert("same")
	tree.Insert("fast")
	tree.Insert("facility")
	tree.Insert("fuck")
	tree.Insert("fi")
	tree.Insert("somethin")

	tree.Traverse(0)
}
