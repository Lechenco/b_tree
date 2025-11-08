package models

import "strings"

type BTreeConfig struct {
	MaxElementsPerNode int
}

type BTree[T any] struct {
	Root   *Node[T]
	Config BTreeConfig
}

func (t *BTree[T]) FindNodeToAddElement(element Element[T]) error {
	if t.Root == nil {
		t.createRootNode()
		t.Root.AddElement(&element, t.Root)

		return nil
	}

	// find node to add
	targetNode := t.Root.FindNodeToAddElement(element)

	targetNode.AddElement(&element, t.Root)

	// check for new root
	if t.Root.Parent != nil {
		t.Root = t.Root.Parent
	}

	return nil
}

func (t *BTree[T]) createRootNode() *Node[T] {
	node := t.createNode()

	t.Root = &node
	return &node
}

func (t *BTree[T]) createNode() Node[T] {
	return Node[T]{MaxElementsPerNode: t.Config.MaxElementsPerNode}
}

func (t BTree[T]) String() string {
	str := "Tree:"
	nodestr := strings.ReplaceAll(t.Root.String(), "\n", "\n\t")

	return str + nodestr
}
