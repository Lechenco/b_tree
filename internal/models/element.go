package models

import "fmt"

type Element[T any] struct {
	Key  int
	Data T
}

func GreaterKeyComparator[T any](key int) func(*Element[T]) bool {
	return func(elem *Element[T]) bool { return elem.Key > key }
}

func EqualsKeyComparator[T any](key int) func(*Element[T]) bool {
	return func(elem *Element[T]) bool { return elem.Key == key }
}

func (e *Element[T]) String() string {
	return fmt.Sprintf("{Element: %d}", e.Key)
}
