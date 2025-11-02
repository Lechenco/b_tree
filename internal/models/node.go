package models

import (
	"slices"
)

type Element[T any] struct {
	Key  int
	Data T
}

type Node[T any] struct {
	Elements           []Element[T]
	ChildNodes         []Node[T]
	MaxElementsPerNode int
}

// X if node has room for one more element, add in the correct order

func (n *Node[T]) AddElement(element Element[T]) {
	if n.isFull() {
		n.splitAndAdd(element)
	}
	index := n.findIndexElementWithGreaterKey(element.Key)
	n.Elements = slices.Insert(n.Elements, index, element)
}

// - split the node in two new nodes so:
//   - a median element is chosen from all node elements + new one
//   - values less then median to left, and other values to the right
//   - the median value is added to the parent node (recursively)
func (n *Node[T]) splitAndAdd(element Element[T]) {
	index := n.findIndexElementWithGreaterKey(element.Key)
	elements := slices.Insert(n.Elements, index, element)

	// get median and both halfs
	var medianIndex int = len(elements) / 2

}

func (n *Node[T]) FindNodeToAddElement(element Element[T]) *Node[T] {
	if n.isLeaf() {
		return n
	}

	nextNodeIndex := n.findIndexElementWithGreaterKey(element.Key)
	if nextNodeIndex == -1 {
		return n.ChildNodes[len(n.ChildNodes)-1].FindNodeToAddElement(element)
	}

	return n.ChildNodes[nextNodeIndex].FindNodeToAddElement(element)
}

func (n *Node[T]) isLeaf() bool { return len(n.ChildNodes) == 0 }

func (n *Node[T]) isFull() bool { return len(n.Elements) == n.MaxElementsPerNode }

func (n *Node[T]) findIndexElementWithGreaterKey(key int) int {
	index := slices.IndexFunc(
		n.Elements,
		func(elem Element[T]) bool { return elem.Key > key },
	)

	return index
}
