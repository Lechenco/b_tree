package services

import (
	"fmt"

	"github.com/Lechenco/b_tree/internal/models"
	"github.com/Lechenco/b_tree/internal/utils"
)

type BTreeService[T any] struct {
	tree models.BTree[T]
}

func (s *BTreeService[T]) InitService(config models.BTreeConfig) {
	s.tree = models.BTree[T]{
		Config: config,
	}
}

func (s *BTreeService[T]) Add(element models.Element[T]) (*models.Element[T], error) {
	err := s.tree.AddElement(element)

	validTree := utils.CheckTree(s.tree)
	if !validTree {
		fmt.Println(s.tree)
		fmt.Println(element)
		panic("Error adding element to Tree")
	}
	return &element, err
}

func (s *BTreeService[T]) Update(element models.Element[T]) (*models.Element[T], error) {
	return nil, nil
}

func (s *BTreeService[T]) Get(key int) *models.Element[T] {
	return s.tree.GetElementByKey(key)
}

func (s *BTreeService[T]) Remove(node models.Element[T]) error {
	return nil
}

func (s BTreeService[T]) String() string {
	return s.tree.String()
}
