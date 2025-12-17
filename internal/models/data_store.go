package models

type DataStore[T any] interface {
	Add(Element[T]) (*Element[T], error)
	Get(int) (*Element[T], error)
	Update(Element[T]) *Element[T]
	Remove(Element[T]) error
}
