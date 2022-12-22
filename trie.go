package main

import (
	"fmt"
	"strings"
)

type TrieNode struct {
	key      string
	children map[rune]*TrieNode
}

func NewTrie() *TrieNode {
	return &TrieNode{
		key:      "",
		children: make(map[rune]*TrieNode),
	}
}

func (n *TrieNode) Insert(word string) {
	word = strings.ToLower(word)
	runes := []rune(word)
	currNode := n

	for _, r := range runes {
		if nextNode, ok := currNode.children[r]; ok {
			currNode = nextNode
		} else {
			currNode.children[r] = &TrieNode{
				key:      string(r),
				children: make(map[rune]*TrieNode),
			}
			currNode = currNode.children[r]
		}
	}
}

func (n *TrieNode) Search(word string) bool {

	if len(n.children) == 0 && len(n.key) == 0 {
		return false
	}

	word = strings.ToLower(word)
	currNode := n

	for _, r := range word {
		if nextNode, ok := currNode.children[r]; ok {
			currNode = nextNode
		} else {
			return false
		}
	}
	return true
}

func (n *TrieNode) Traverse(word string) {
	currWord := word + n.key

	if len(n.children) == 0 {
		fmt.Println(currWord)
		return
	}

	for _, child := range n.children {
		child.Traverse(currWord)
	}
}

func TrieShow() {
	trie := NewTrie()

	trie.Insert("home")
	trie.Insert("hell")
	trie.Insert("hope")
	trie.Insert("meal")
	trie.Insert("burger")
	trie.Insert("sun")
	trie.Insert("shine")
	trie.Insert("kill")
	trie.Insert("bill")
	trie.Insert("homeless")

	trie.Traverse("")

	fmt.Println(trie.Search("homelesss"))

}
