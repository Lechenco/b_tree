package models

type BTreeConfig struct {
	MaxElementsPerNode int
}

type BTree[T any] struct {
	Root   *Node[T]
	Config BTreeConfig
}

func (t *BTree[T]) FindNodeToAddElement(element Element[T]) *Node[T] {
	// find node
	if t.Root == nil {
		return t.createRootNode()
	}

	// find node to add
	targetNode := t.Root.FindNodeToAddElement(element)
	return targetNode
}

func (t *BTree[T]) createRootNode() *Node[T] {
	node := t.createNode()

	t.Root = &node
	return &node
}

func (t *BTree[T]) createNode() Node[T] {
	return Node[T]{MaxElementsPerNode: t.Config.MaxElementsPerNode}
}
