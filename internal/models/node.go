package models

import (
	"fmt"
	"slices"
	"strings"
)

type Element[T any] struct {
	Key  int
	Data T
}

type Node[T any] struct {
	Parent             *Node[T]
	Elements           []*Element[T]
	ChildNodes         []*Node[T]
	MaxElementsPerNode int
}

func (n *Node[T]) AddChild(child *Node[T]) {
	n.ChildNodes = append(n.ChildNodes, child)
	child.Parent = n
}

// X if node has room for one more element, add in the correct order

func (n *Node[T]) AddElement(element *Element[T], root *Node[T]) {
	index := n.findIndexElementWithGreaterKey(element.Key)

	if index == -1 {
		n.Elements = append(n.Elements, element)
	} else {
		n.Elements = slices.Insert(n.Elements, index, element)
	}

	if n.isOverflowed() {
		fmt.Println("Partial Tree: ", root)
		n.splitElements(root)
		// n.updateChildParents()
	}
}

func (n *Node[T]) splitElements(root *Node[T]) {
	// get median and both halfs
	medianIndex := len(n.Elements) / 2
	medianElement := n.Elements[medianIndex]
	leftElements, leftNodes := n.sliceElementsAndChildsAtLeft(medianIndex)
	rightElements, rightNodes := n.sliceElementsAndChildsAtRight(medianIndex)

	n.Elements = leftElements
	n.ChildNodes = leftNodes

	if n.Parent == nil {
		newRoot := Node[T]{
			ChildNodes:         []*Node[T]{n},
			MaxElementsPerNode: n.MaxElementsPerNode,
		}
		n.Parent = &newRoot
	}

	var rightNode = n.copyNode(rightElements, rightNodes)

	n.Parent.AddElement(medianElement, root)
	n.Parent.AddChild(&rightNode)

}

func (n *Node[T]) updateChildParents() {
	newChilds := []*Node[T]{}
	for _, child := range n.ChildNodes {
		child.Parent = n
		newChilds = append(newChilds, child)
	}
	n.ChildNodes = newChilds
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

func (n *Node[T]) isOverflowed() bool { return len(n.Elements) > n.MaxElementsPerNode }

func (n *Node[T]) findIndexElementWithGreaterKey(key int) int {
	index := slices.IndexFunc(
		n.Elements,
		func(elem *Element[T]) bool { return elem.Key > key },
	)

	return index
}

func (n *Node[T]) sliceElementsAndChildsAtLeft(index int) ([]*Element[T], []*Node[T]) {
	sliceElements := slices.Clone(n.Elements)[:index]
	if len(n.ChildNodes) == 0 {
		return sliceElements, []*Node[T]{}
	}
	sliceNodes := slices.Clone(n.ChildNodes)[:index+1]

	return sliceElements, sliceNodes
}

func (n *Node[T]) sliceElementsAndChildsAtRight(index int) ([]*Element[T], []*Node[T]) {
	sliceElements := slices.Clone(n.Elements)[index+1:]
	if len(n.ChildNodes) == 0 {
		return sliceElements, []*Node[T]{}
	}
	sliceNodes := slices.Clone(n.ChildNodes)[index+1:]

	return sliceElements, sliceNodes
}

func (n *Node[T]) copyNode(elements []*Element[T], childs []*Node[T]) Node[T] {
	return Node[T]{
		Parent:             n.Parent,
		Elements:           elements,
		MaxElementsPerNode: n.MaxElementsPerNode,
		ChildNodes:         childs,
	}
}

func (e Element[T]) String() string {
	return fmt.Sprintf("{Element: %d}", e.Key)
}

func (n Node[T]) String() string {
	str := fmt.Sprintf("\nNode %p: {Parent: %p [%v]}", &n, n.Parent, n.Elements)

	var childStr string
	for _, v := range n.ChildNodes {
		childStr = childStr + v.String()
	}

	childStr = strings.ReplaceAll(childStr, "\n", "\n\t")

	return str + childStr
}
