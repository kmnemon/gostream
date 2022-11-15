package gostream

import "golang.org/x/exp/constraints"

type stream[T constraints.Ordered] interface {
	Map(func(T) T) stream[T]
	ForEach(func(T))
	// filter(func(T) bool) stream[T]
	Sorted() stream[T]
	// limit(uint64) Stream[T1, T2]

	// OpWrapSink(Sink[T]) Sink[T]
}

func StreamOf[T constraints.Ordered](sourceData []T) stream[T] {
	p := pipeline[T]{}
	p.init(nil, head, nil)
	p.sourceData = sourceData
	return &p
}
