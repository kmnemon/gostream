package gostream

type Stream[T any] interface {
	Map(func(T) T) Stream[T]
	// filter(func(T2) bool) Stream[T1, T2]
	// sorted() Stream[T1, T2]
	// limit(uint64) Stream[T1, T2]

	// OpWrapSink(Sink[T]) Sink[T]
}

func StreamOf[T any]() Stream[T] {
	p := Pipeline[T]{}
	p.New(nil, Head)
	return &p
}
