package models

import (
	"fmt"
	"slices"
	"strings"
)

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

func (n *Node[T]) AddElement(element *Element[T], root *Node[T]) *Node[T] {
	index := n.indexElementFunc(GreaterKeyComparator[T](element.Key))

	if index == -1 {
		n.Elements = append(n.Elements, element)
	} else {
		n.Elements = slices.Insert(n.Elements, index, element)
	}

	if n.isOverflowed() {
		n.splitElements(root)
		return root.findRoot().findNode(*element)
	}
	return n
}

func (n *Node[T]) splitElements(root *Node[T]) {
	medianIndex := len(n.Elements) / 2
	medianElement := n.Elements[medianIndex]

	n.addElementToParent(medianElement, root)

	leftElements, leftNodes := n.sliceElementsAndChildsAtLeft(medianIndex)
	rightElements, rightNodes := n.sliceElementsAndChildsAtRight(medianIndex)

	n.Elements = leftElements
	n.ChildNodes = leftNodes

	var rightNode = n.copyNode(rightElements, rightNodes)
	n.Parent.AddChild(&rightNode)
}

func (n *Node[T]) addElementToParent(element *Element[T], root *Node[T]) {
	if n.Parent == nil {
		newRoot := Node[T]{
			ChildNodes:         []*Node[T]{n},
			MaxElementsPerNode: n.MaxElementsPerNode,
		}
		n.Parent = &newRoot
	}
	n.Parent = n.Parent.AddElement(element, root)
}

func (n *Node[T]) FindNodeToAddElement(element Element[T]) *Node[T] {
	if n.isLeaf() {
		return n
	}

	return n.nextNode(element).FindNodeToAddElement(element)
}

func (n *Node[T]) findRoot() *Node[T] {
	if n.Parent == nil {
		return n
	}

	return n.Parent.findRoot()
}

func (n *Node[T]) findNode(element Element[T]) *Node[T] {
	index := n.indexElementFunc(EqualsKeyComparator[T](element.Key))

	if index != -1 {
		return n
	} else if len(n.ChildNodes) == 0 {
		return nil
	}

	return n.nextNode(element).findNode(element)
}

func (n *Node[T]) nextNode(element Element[T]) *Node[T] {
	nextNodeIndex := n.indexElementFunc(GreaterKeyComparator[T](element.Key))
	if nextNodeIndex == -1 {
		return n.ChildNodes[len(n.ChildNodes)-1]
	}

	return n.ChildNodes[nextNodeIndex]
}

func (n *Node[T]) isLeaf() bool { return len(n.ChildNodes) == 0 }

func (n *Node[T]) isOverflowed() bool { return len(n.Elements) > n.MaxElementsPerNode }

func (n *Node[T]) indexElementFunc(f func(*Element[T]) bool) int {
	return slices.IndexFunc(n.Elements, f)
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

func (n *Node[T]) String() string {
	str := fmt.Sprintf("\nNode %p: {Parent: %p [%v]}", n, n.Parent, n.Elements)

	var childStr string
	for _, v := range n.ChildNodes {
		childStr = childStr + v.String()
	}

	childStr = strings.ReplaceAll(childStr, "\n", "\n\t")

	return str + childStr
}
