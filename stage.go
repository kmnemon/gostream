package gostream

type Head[T any] struct {
}

func (h *Head[T]) of(T) Stream[T] {
	return nil
}

type StatelessOp struct {
}

type StatefulOp struct {
}
