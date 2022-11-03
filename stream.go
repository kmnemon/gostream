package gostream

type Stream[T any] interface {
	of(T) Stream[T]
	mapR(func(T2) T2) Stream[T1, T2]
	// filter(func(T2) bool) Stream[T1, T2]
	// sorted() Stream[T1, T2]
	// limit(uint64) Stream[T1, T2]
}

type MakeStream[T any] struct {
}

func (ms *MakeStream[T]) of(T) Stream[T] {
	return &Head[T]{}
}

// type HH[T any] struct {
// }

// func (h HH[T]) of(T) Stream[T] {
// 	return nil
// }

// type H1 struct{}

// func hh() {
// 	h := HH[[]any]{}

// }
