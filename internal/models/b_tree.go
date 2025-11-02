package models

type BTreeConfig struct {
	MaxKeysPerNode int
}

type BTree[T any] struct {
	Root   Node[T]
	Config BTreeConfig
}

func NewBTree[T any]() BTree[T] {
	return BTree[T]{
		Config: BTreeConfig{
			MaxKeysPerNode: 1,
		},
	}
}
