package main

import (
	"fmt"
)

type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

type BinarySearchTree struct {
	Root *Node
}

func (b *BinarySearchTree) insert(key int) {
	if b.Root == nil {
		b.Root = &Node{
			Key:   key,
			Left:  nil,
			Right: nil,
		}
	} else {
		recursiveInsert(b.Root, &Node{
			Key:   key,
			Left:  nil,
			Right: nil,
		})
	}
}

func recursiveInsert(targetNode *Node, newNode *Node) {
	if newNode.Key < targetNode.Key {
		if targetNode.Left == nil {
			targetNode.Left = newNode
		} else {
			recursiveInsert(targetNode.Left, newNode)
		}
	} else {
		if targetNode.Right == nil {
			targetNode.Right = newNode
		} else {
			recursiveInsert(targetNode.Right, newNode)
		}
	}
}

func (b *BinarySearchTree) remove(key int) {
	recursiveRemove(b.Root, key)
}

func recursiveRemove(targetNode *Node, key int) *Node {
	if targetNode == nil {
		return nil
	}

	if key < targetNode.Key {
		targetNode.Left = recursiveRemove(targetNode.Left, key)
		return targetNode
	}

	if key > targetNode.Key {
		targetNode.Right = recursiveRemove(targetNode.Right, key)
		return targetNode
	}

	if targetNode.Left == nil && targetNode.Right == nil {
		targetNode = nil
		return nil
	}

	if targetNode.Left == nil {
		targetNode = targetNode.Right
		return targetNode
	}

	if targetNode.Right == nil {
		targetNode = targetNode.Left
		return targetNode
	}

	// find min in right node
	leftNodeOfMostRightNode := targetNode.Right

	for {
		if leftNodeOfMostRightNode != nil && leftNodeOfMostRightNode.Left != nil {
			leftNodeOfMostRightNode = leftNodeOfMostRightNode.Left
		} else {
			break
		}
	}

	targetNode.Key = leftNodeOfMostRightNode.Key
	targetNode.Right = recursiveRemove(targetNode.Right, targetNode.Key)
	return targetNode

	/*
		// find max in left node
		rightNodeOfMostLeftNode = targetNode.Left

		for {
			if rightNodeOfMostLeftNode != nil && rightNodeOfMostLeftNode.Right != nil {
				rightNodeOfMostLeftNode = rightNodeOfMostLeftNode.Left
			} else {
				break
			}
		}

		targetNode.key = rightNodeOfMostLeftNode.key
		targetNode.Left = removeRecursive(targetNode.Left, targetNode.Key)
		return targetNode

	*/
}

func (b *BinarySearchTree) search(key int) bool {
	return recursiveSearch(b.Root, key)
}

func recursiveSearch(node *Node, key int) bool {
	if node == nil {
		return false
	}

	if node.Key > key {
		return recursiveSearch(node.Right, key)
	}

	if node.Key < key {
		return recursiveSearch(node.Left, key)
	}

	return true
}

func (b *BinarySearchTree) PreOrderTraverse() {
	recursivePreOrderTraverse(b.Root)
}

func recursivePreOrderTraverse(node *Node) {
	if node != nil {
		fmt.Print(node.Key, " ")
		recursivePreOrderTraverse(node.Left)
		recursivePreOrderTraverse(node.Right)
	}
}

func (b *BinarySearchTree) InOrderTraverse() {
	recursiveInOrderTraverse(b.Root)
}

func recursiveInOrderTraverse(node *Node) {
	if node != nil {
		recursivePreOrderTraverse(node.Left)
		fmt.Print(node.Key, " ")
		recursivePreOrderTraverse(node.Right)
	}
}

func (b *BinarySearchTree) PostOrderTraverse() {
	recursivePostOrderTraverse(b.Root)
}

func recursivePostOrderTraverse(node *Node) {
	if node != nil {
		recursivePreOrderTraverse(node.Left)
		recursivePreOrderTraverse(node.Right)
		fmt.Print(node.Key, " ")
	}
}

func (b *BinarySearchTree) BFS() {
	BFS(b.Root)
}

func BFS(root *Node) {
	order := make([]*Node, 0)

	order = append(order, root)

	for indexNode := 0; indexNode < len(order); indexNode += 1 {
		if order[indexNode] != nil {
			fmt.Print(order[indexNode].Key, " ")
			order = append(order, order[indexNode].Left, order[indexNode].Right)
			// fmt.Println(len(order))
		}
	}

}

func BSTShow() {
	tree := &BinarySearchTree{}

	tree.insert(10)
	tree.insert(15)
	tree.insert(18)
	tree.insert(16)
	tree.insert(14)
	tree.insert(2)
	tree.insert(3)

	tree.PreOrderTraverse()
	fmt.Println()
	tree.InOrderTraverse()
	fmt.Println()
	tree.PostOrderTraverse()
	fmt.Println()

	tree.BFS()
	fmt.Println()

	tree.remove(10)

	tree.BFS()
	fmt.Println()
}

func main() {
	TrieShow()
}
