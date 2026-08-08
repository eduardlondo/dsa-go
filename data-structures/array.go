package datastructures

type DynamicArray[T any] struct {
	length   int
	capacity int
	data     []T
}

func NewDynamicArray[T any](initialCap int) *DynamicArray[T] {
	if initialCap <= 0 {
		initialCap = 2
	}
	return &DynamicArray[T]{
		data:     make([]T, initialCap),
		length:   0,
		capacity: initialCap,
	}
}
