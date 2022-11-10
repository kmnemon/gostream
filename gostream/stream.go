package gostream

type Stream[T any, R any] interface {
	Map(func(T) R) Stream[T, R]
	// filter(func(T2) bool) Stream[T1, T2]
	// sorted() Stream[T1, T2]
	// limit(uint64) Stream[T1, T2]

	OpWrapSink(Sink[T]) Sink[T]
}

func StreamOf[T any, R any]() Stream[T, R] {
	p := Pipeline[T, R]{}
	p.New(nil, Head)
	return &p
}
