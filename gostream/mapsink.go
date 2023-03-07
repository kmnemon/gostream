package gostream

type mapSink[T any] struct {
	mapper     func(T) T
	downstream sink[T]
}

func (s *mapSink[T]) begin(size int) {
	s.downstream.begin(size)
}

func (s *mapSink[T]) accept(u T) {
	r := s.mapper(u)
	s.downstream.accept(r)
}

func (s *mapSink[T]) end() {
	s.downstream.end()
}

func (s *mapSink[T]) isCancellationWasRequested() bool {
	return s.downstream.isCancellationWasRequested()
}

func (s *mapSink[T]) cancellationRequested() bool {
	return s.downstream.cancellationRequested()
}

func (s *mapSink[T]) setDownStreamSink(downstream sink[T]) {
	s.downstream = downstream
}

func (s *mapSink[T]) canParallel() bool {
	return s.downstream.canParallel()
}
