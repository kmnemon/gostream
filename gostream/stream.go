package gostream

type stream[T any] interface {
	Map(func(T) T) stream[T]
	ForEach(func(T))
	// filter(func(T) bool) stream[T]
	Sorted() stream[T]
	SortedWith(func(T, T) bool) stream[T]
	// limit(uint64) Stream[T1, T2]

	// OpWrapSink(Sink[T]) Sink[T]
}

func StreamOf[T any](sourceData []T) stream[T] {
	p := pipeline[T]{}
	p.init(nil, head, nil)
	p.sourceData = sourceData
	return &p
}
