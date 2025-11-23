package models

import (
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
	addElementRecursively(t.Root, element)

	if t.Root.isOverflowed() {
		previousRoot := t.Root
		root := t.createRootNode()
		root.AddChild(previousRoot)
		splitChildrenNode(root, previousRoot)
	}
	return nil
}

func addElementRecursively[T any](node *Node[T], element Element[T]) error {
	if node.isLeaf() {
		node.AddElement(&element, nil)
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

	parent.AddElement(medianElement, nil)

	rightNode := children.copyNode(rightElements, rightNodes)
	parent.AddChild(&rightNode)
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
