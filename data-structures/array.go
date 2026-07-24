package datastructures

type DynamicArray[T any] struct {
	length   int
	capacity int
	data     []T
}
