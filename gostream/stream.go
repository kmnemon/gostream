package gostream

type stream[T any] interface {
	Map(func(T) T) stream[T]
	ForEach(func(T))
	// filter(func(T2) bool) Stream[T1, T2]
	// sorted() Stream[T1, T2]
	// limit(uint64) Stream[T1, T2]

	// OpWrapSink(Sink[T]) Sink[T]
}

func StreamOf[T any](slice []T) stream[T] {
	p := pipeline[T]{}
	p.new(nil, head, nil)
	return &p
}
