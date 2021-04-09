package datastructs

import "fmt"

type Node struct {
	Value      int
	LeftChild  *Node
	RightChild *Node
	Parent     *Node
}

type Tree struct {
	RootNode Node
}

func MakeTree(rootVal int) Tree {
	return Tree{RootNode: Node{Value: rootVal, LeftChild: nil, RightChild: nil, Parent: nil}}
}

func (tree *Tree) Insert(value int) {
	insertNode(&(tree.RootNode), value)
}

func insertNode(currentNode *Node, value int) {
	if value == currentNode.Value {
		//	node already exists, do nothing
	} else if value < currentNode.Value {
		if currentNode.LeftChild != nil {
			insertNode(currentNode.LeftChild, value)
		} else {
			currentNode.LeftChild = &Node{Value: value, LeftChild: nil, RightChild: nil, Parent: currentNode}
		}
	} else {
		if currentNode.RightChild != nil {
			insertNode(currentNode.RightChild, value)
		} else {
			currentNode.RightChild = &Node{Value: value, LeftChild: nil, RightChild: nil, Parent: currentNode}
		}
	}
}

func (tree *Tree) Delete(value int) {
	deleteNode(&(tree.RootNode), value)
}

func deleteNode(currentNode *Node, value int) {
	if value == currentNode.Value {
		if currentNode.LeftChild == nil && currentNode.RightChild == nil {
			replaceCurrent(currentNode, nil)
		} else if currentNode.LeftChild != nil && currentNode.RightChild == nil {
			replaceCurrent(currentNode, currentNode.LeftChild)
		} else if currentNode.LeftChild == nil && currentNode.RightChild != nil {
			replaceCurrent(currentNode, currentNode.RightChild)
		} else {
			inOrderSuccessor := leftLeaf(currentNode.RightChild)
			currentNode.Value = inOrderSuccessor.Value
			deleteNode(inOrderSuccessor, inOrderSuccessor.Value)
		}
	} else if value < currentNode.Value {
		if currentNode.LeftChild != nil {
			deleteNode(currentNode.LeftChild, value)
		} else {
			//	node does not exist, do nothing
		}
	} else {
		if currentNode.RightChild != nil {
			deleteNode(currentNode.RightChild, value)
		} else {
			//	node does not exist, do nothing
		}
	}
}

func replaceCurrent(currentNode *Node, newNode *Node) {
	if currentNode.Parent.LeftChild == currentNode {
		currentNode.Parent.LeftChild = newNode
	} else {
		currentNode.Parent.RightChild = newNode
	}
}

func leftLeaf(currentNode *Node) *Node {
	if currentNode.LeftChild != nil {
		return leftLeaf(currentNode.LeftChild)
	} else {
		return currentNode
	}
}

func PrintInOrder(node *Node) {
	if node != nil {
		PrintInOrder(node.LeftChild)
		fmt.Println(node.Value)
		PrintInOrder(node.RightChild)
	}
}
