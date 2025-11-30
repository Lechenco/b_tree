package models

import (
	"fmt"
	"slices"
	"strings"
)

type Node[T any] struct {
	Elements           []*Element[T]
	ChildNodes         []*Node[T]
	MaxElementsPerNode int
}

func (n *Node[T]) AddChild(child *Node[T]) {
	childLastElement := child.Elements[len(child.Elements)-1]
	index := n.indexElementFunc(GreaterKeyComparator[T](childLastElement.Key))

	if index == -1 {
		n.ChildNodes = append(n.ChildNodes, child)
	} else {
		n.ChildNodes = slices.Insert(n.ChildNodes, index, child)
	}

}

func (n *Node[T]) AddElement(element *Element[T]) *Node[T] {
	index := n.indexElementFunc(GreaterKeyComparator[T](element.Key))

	if index == -1 {
		n.Elements = append(n.Elements, element)
	} else {
		n.Elements = slices.Insert(n.Elements, index, element)
	}

	return n
}

func (n *Node[T]) FindElement(element Element[T]) *Node[T] {
	index := n.indexElementFunc(EqualsKeyComparator[T](element.Key))
	if index != -1 {
		return n
	}
	if n.IsLeaf() {
		return nil
	}

	return n.nextNode(element).FindElement(element)
}

func (n *Node[T]) nextNode(element Element[T]) *Node[T] {
	nextNodeIndex := n.indexElementFunc(GreaterKeyComparator[T](element.Key))
	if nextNodeIndex == -1 {
		return n.ChildNodes[len(n.ChildNodes)-1]
	}

	return n.ChildNodes[nextNodeIndex]
}

func (n *Node[T]) IsLeaf() bool { return len(n.ChildNodes) == 0 }

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
		Elements:           elements,
		MaxElementsPerNode: n.MaxElementsPerNode,
		ChildNodes:         childs,
	}
}

func (n *Node[T]) String() string {
	str := fmt.Sprintf("\nNode %p: [%v]", n, n.Elements)

	var childStr string
	for _, v := range n.ChildNodes {
		childStr = childStr + v.String()
	}

	childStr = strings.ReplaceAll(childStr, "\n", "\n\t")

	return str + childStr
}
