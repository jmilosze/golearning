package datastructs

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

func (tree *Tree) Insert(value int) *Tree {
	insertNode(&(tree.RootNode), value)
	return tree
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

func (tree *Tree) Delete(value int) *Tree {
	deleteNode(&(tree.RootNode), value)
	return tree
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
	if currentNode.Parent != nil {
		if currentNode.Parent.LeftChild == currentNode {
			currentNode.Parent.LeftChild = newNode
		} else {
			currentNode.Parent.RightChild = newNode
		}
	} else {
		// this means we are deleting the root node
		// we do not allow to delete root node if it's the only node (case when newNode == nil)
		if newNode != nil {
			*currentNode = *newNode
			currentNode.Parent = nil
		}
	}
}

func leftLeaf(currentNode *Node) *Node {
	if currentNode.LeftChild != nil {
		return leftLeaf(currentNode.LeftChild)
	} else {
		return currentNode
	}
}

func (tree *Tree) ReturnInOrder() []int {
	orderedValues := make([]int, 100)[0:0]
	return returnInOrder(&(tree.RootNode), orderedValues)
}

func returnInOrder(node *Node, orderedValues []int) []int {
	if node != nil {
		orderedValues = returnInOrder(node.LeftChild, orderedValues)
		orderedValues = append(orderedValues, node.Value)
		orderedValues = returnInOrder(node.RightChild, orderedValues)
	}
	return orderedValues
}
