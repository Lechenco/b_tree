package models

type DataStore[T any] interface {
	Add(T) Node[T]
	Get(int) Node[T]
	Remove(Node[T])
}
