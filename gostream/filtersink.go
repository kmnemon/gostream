package gostream

type filterSink[T any] struct {
	predicate  func(T) bool
	downstream sink[T]
}

func (s *filterSink[T]) begin(size int) {
	s.downstream.begin(-1)
}

func (s *filterSink[T]) accept(u T) {
	if s.predicate(u) {
		s.downstream.accept(u)
	}
}

func (s *filterSink[T]) end() {
	s.downstream.end()
}

func (s *filterSink[T]) cancellationRequested() bool {
	return false
}

func (s *filterSink[T]) setDownStreamSink(downstream sink[T]) {
	s.downstream = downstream
}
