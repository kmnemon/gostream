package gostream

type stream[T any] interface {
	Map(func(T) T) stream[T]
	ForEach(func(T))
	Sorted() stream[T]
	SortedWith(func(T, T) bool) stream[T]
	Filter(func(T) bool) stream[T]
	Limit(int) stream[T]
	FindFirst() T
}

func StreamOf[T any](sourceData []T) stream[T] {
	p := pipeline[T]{}
	p.init(nil, head, nil)
	p.sourceData = sourceData
	return &p
}
