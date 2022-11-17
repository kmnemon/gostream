package gostream

type toListSink[T any] struct {
	list []T
}

func (s *toListSink[T]) begin(size int) {
}

func (s *toListSink[T]) accept(u T) {
	s.list = append(s.list, u)
}

func (s *toListSink[T]) end() {
}

func (s *toListSink[T]) isCancellationWasRequested() bool {
	return false
}

func (s *toListSink[T]) cancellationRequested() bool {
	return false
}

func (s *toListSink[T]) setDownStreamSink(downstream sink[T]) {
}
