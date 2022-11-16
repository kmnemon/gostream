package gostream

type limitSink[T any] struct {
	maxSize    int
	downstream sink[T]
}

func (s *limitSink[T]) begin(size int) {
	s.downstream.begin(-1)
}

func (s *limitSink[T]) accept(u T) {
	if s.maxSize > 0 {
		s.downstream.accept(u)
		s.maxSize--
	}
}

func (s *limitSink[T]) end() {
	s.downstream.end()
}

func (s *limitSink[T]) isCancellationWasRequested() bool {
	return s.downstream.isCancellationWasRequested()
}

func (s *limitSink[T]) cancellationRequested() bool {
	return s.downstream.cancellationRequested()
}

func (s *limitSink[T]) setDownStreamSink(downstream sink[T]) {
	s.downstream = downstream
}
