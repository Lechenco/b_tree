package models

type DataStore[T any] interface {
	Add(T) (*Node[T], error)
	Get(int) (*Node[T], error)
	Remove(Node[T]) error
}
