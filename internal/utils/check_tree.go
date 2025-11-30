package utils

import (
	"math"

	"github.com/Lechenco/b_tree/internal/models"
)

func CheckTree[T any](tree models.BTree[T]) bool {
	return checkNodeDeep(tree.Root, 0, math.MaxInt)
}

func checkNodeDeep[T any](node *models.Node[T], leftSeparator, rightSeparator int) bool {
	isValid := checkNode(node, leftSeparator, rightSeparator)

	if !isValid {
		return isValid
	}

	for index, children := range node.ChildNodes {
		leftSeparator, rightSeparator := getSeparators(index, leftSeparator, rightSeparator, node.Elements)
		childrenIsValid := checkNodeDeep(children, leftSeparator, rightSeparator)
		isValid = isValid && childrenIsValid

		if !isValid {
			break
		}
	}

	return isValid
}

func checkNode[T any](node *models.Node[T], leftSeparator, rightSeparator int) bool {
	isValid := true
	for _, elem := range node.Elements {
		key := elem.Key
		isValid = isValid && leftSeparator < key && key < rightSeparator

		if !isValid {
			break
		}
	}

	return isValid
}

func getSeparators[T any](index, leftSeparator, rightSeparator int, elements []*models.Element[T]) (int, int) {
	if index == 0 {
		return leftSeparator, elements[index].Key
	}
	if index == len(elements) {
		return elements[index-1].Key, rightSeparator
	}

	return elements[index-1].Key, elements[index].Key
}
