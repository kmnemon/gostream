package gostream

type limitSink[T any] struct {
	maxSize    int
	downstream sink[T]
	cancel     bool
}

func (s *limitSink[T]) begin(size int) {
	s.downstream.begin(-1)
}

func (s *limitSink[T]) accept(u T) {
	if s.maxSize > 0 {
		s.downstream.accept(u)
		s.maxSize--
	} else {
		s.cancel = true
	}
}

func (s *limitSink[T]) end() {
	s.downstream.end()
}

func (s *limitSink[T]) isCancellationWasRequested() bool {
	return true
}

func (s *limitSink[T]) cancellationRequested() bool {
	return s.cancel
}

func (s *limitSink[T]) setDownStreamSink(downstream sink[T]) {
	s.downstream = downstream
}

func (s *limitSink[T]) canParallel() bool {
	return false
}
