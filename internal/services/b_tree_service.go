package services

import "github.com/Lechenco/b_tree/internal/models"

type BTreeService[T any] struct {
	tree models.BTree[T]
}

func (s *BTreeService[T]) InitService(config models.BTreeConfig) {
	s.tree = models.BTree[T]{
		Config: config,
	}
}

func (s *BTreeService[T]) Add(element models.Element[T]) (*models.Element[T], error) {
	err := s.tree.FindNodeToAddElement(element)

	return &element, err
}

func (s *BTreeService[T]) Get(key int) (*models.Element[T], error) {
	return nil, nil
}

func (s *BTreeService[T]) Remove(node models.Element[T]) error {
	return nil
}

func (s BTreeService[T]) String() string {
	return s.tree.String()
}
