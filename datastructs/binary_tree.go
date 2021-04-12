package datastructs

type BSTNode struct {
	Value      int
	LeftChild  *BSTNode
	RightChild *BSTNode
	Parent     *BSTNode
}

type BSTree struct {
	RootNode BSTNode
}

func MakeBSTree(rootVal int) BSTree {
	return BSTree{RootNode: BSTNode{Value: rootVal, LeftChild: nil, RightChild: nil, Parent: nil}}
}

func (tree *BSTree) Insert(value int) *BSTree {
	insertNode(&(tree.RootNode), value)
	return tree
}

func insertNode(currentNode *BSTNode, value int) {
	if value == currentNode.Value {
		//	node already exists, do nothing
	} else if value < currentNode.Value {
		if currentNode.LeftChild != nil {
			insertNode(currentNode.LeftChild, value)
		} else {
			currentNode.LeftChild = &BSTNode{Value: value, LeftChild: nil, RightChild: nil, Parent: currentNode}
		}
	} else {
		if currentNode.RightChild != nil {
			insertNode(currentNode.RightChild, value)
		} else {
			currentNode.RightChild = &BSTNode{Value: value, LeftChild: nil, RightChild: nil, Parent: currentNode}
		}
	}
}

func (tree *BSTree) Delete(value int) *BSTree {
	deleteNode(&(tree.RootNode), value)
	return tree
}

func deleteNode(currentNode *BSTNode, value int) {
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

func replaceCurrent(currentNode *BSTNode, newNode *BSTNode) {
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

func leftLeaf(currentNode *BSTNode) *BSTNode {
	if currentNode.LeftChild != nil {
		return leftLeaf(currentNode.LeftChild)
	} else {
		return currentNode
	}
}

func (tree *BSTree) ReturnInOrder() []int {
	orderedValues := make([]int, 100)[0:0]
	return returnInOrder(&(tree.RootNode), orderedValues)
}

func returnInOrder(node *BSTNode, orderedValues []int) []int {
	if node != nil {
		orderedValues = returnInOrder(node.LeftChild, orderedValues)
		orderedValues = append(orderedValues, node.Value)
		orderedValues = returnInOrder(node.RightChild, orderedValues)
	}
	return orderedValues
}
