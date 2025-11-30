package models

import (
	"fmt"
	"strings"
)

type BTreeConfig struct {
	MaxElementsPerNode int
}

type BTree[T any] struct {
	Root   *Node[T]
	Config BTreeConfig
}

func (t *BTree[T]) AddElement(element Element[T]) error {
	if t.Root == nil {
		t.createRootNode()
	}

	if t.elementExists(element) {
		fmt.Println("Element already in the tree, ignoring add action")
		return nil
	}

	addElementRecursively(t.Root, element)

	if t.Root.isOverflowed() {
		t.splitTreeRoot()
	}
	return nil
}

func (t *BTree[T]) UpdateElement(element Element[T]) error {
	panic("Not Implemented")
	// find element by key
	// update data
	// error if not found
}

func (t *BTree[T]) DeleteElement(element Element[T]) error {
	panic("Not Implemented")
	/*
		https://www.programiz.com/dsa/deletion-from-a-b-tree

		Case 1: delete from a  leaf node
			Case 1.a: delete does not cause underflow
				delete element
			Case 1.b: delete causes a underflow
				borrow a element from a sibling node (largest from left and lowest from rigth)
				replace the separator with the choosen element and add the separator to the child node
			Case 1.c: both siblings have a minimum number of keys
				merge the node with the left or right sibling plus the separator

		Case 2: delete from internal node
			Case 2.a: the deleted key is replaced by a predecessor
			Case 2.b: the deleted key is replaced by a successor
			Case 2.c: if either child has exactly a minimum number of keys
				merge the left and the right children
				if the deleted key node has less than minimum number of keys, look for siblings for borrow

		Case 3: The tree shrinks


	*/

}

func (t *BTree[T]) splitTreeRoot() error {
	previousRoot := t.Root
	root := t.createRootNode()
	root.AddChild(previousRoot)
	splitChildrenNode(root, previousRoot)

	return nil
}

func addElementRecursively[T any](node *Node[T], element Element[T]) error {
	if node.IsLeaf() {
		node.AddElement(&element)
		return nil
	}

	nextNode := node.nextNode(element)
	addElementRecursively(nextNode, element)

	if !nextNode.isOverflowed() {
		return nil
	}

	splitChildrenNode(node, nextNode)

	return nil
}

func splitChildrenNode[T any](parent, children *Node[T]) {
	medianIndex := len(children.Elements) / 2
	medianElement := children.Elements[medianIndex]

	leftElements, leftNodes := children.sliceElementsAndChildsAtLeft(medianIndex)
	rightElements, rightNodes := children.sliceElementsAndChildsAtRight(medianIndex)

	children.Elements = leftElements
	children.ChildNodes = leftNodes

	parent.AddElement(medianElement)

	rightNode := children.copyNode(rightElements, rightNodes)
	parent.AddChild(&rightNode)
}

func (t *BTree[T]) elementExists(element Element[T]) bool {
	return t.Root.FindElement(element) != nil
}

func (t *BTree[T]) createRootNode() *Node[T] {
	node := t.createNode()

	t.Root = &node
	return &node
}

func (t *BTree[T]) createNode() Node[T] {
	return Node[T]{MaxElementsPerNode: t.Config.MaxElementsPerNode}
}

func (t *BTree[T]) String() string {
	str := "Tree:"
	nodestr := strings.ReplaceAll(t.Root.String(), "\n", "\n\t")

	return str + nodestr
}
