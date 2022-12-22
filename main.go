package main

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

func main() {

}
