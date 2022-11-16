package gostream

type forEachSink[T any] struct {
	mapper func(T)
}

func (s *forEachSink[T]) begin(size int) {
}

func (s *forEachSink[T]) accept(u T) {
	s.mapper(u)
}

func (s *forEachSink[T]) end() {
}

func (s *forEachSink[T]) isCancellationWasRequested() bool {
	return false
}

func (s *forEachSink[T]) cancellationRequested() bool {
	return false
}

func (s *forEachSink[T]) setDownStreamSink(downstream sink[T]) {
}
