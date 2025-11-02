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

func (s *BTreeService[T]) Add(data T) (*models.Node[T], error) {
	return nil, nil
}

func (s *BTreeService[T]) Get(key int) (*models.Node[T], error) {
	return nil, nil
}

func (s *BTreeService[T]) Remove(node models.Node[T]) error {
	return nil
}
