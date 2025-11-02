package models

type Node[T any] struct {
	Data  T
	Keys  []int
	Nodes []Node[T]
}
