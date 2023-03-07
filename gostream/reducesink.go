package gostream

type reduceSink[T any] struct {
	binaryOperator func(T, T) T
	downstream     sink[T]
	i              T
	isFirstValue   bool
}

func (s *reduceSink[T]) begin(size int) {
	s.downstream.begin(size)
}

func (s *reduceSink[T]) accept(u T) {
	if s.isFirstValue {
		s.i = u
		s.isFirstValue = false
	} else {
		s.i = s.binaryOperator(s.i, u)
	}
	s.downstream.accept(s.i)
}

func (s *reduceSink[T]) end() {
	s.downstream.end()
}

func (s *reduceSink[T]) isCancellationWasRequested() bool {
	return s.downstream.isCancellationWasRequested()
}

func (s *reduceSink[T]) cancellationRequested() bool {
	return s.downstream.cancellationRequested()
}

func (s *reduceSink[T]) setDownStreamSink(downstream sink[T]) {
	s.downstream = downstream
}

func (s *reduceSink[T]) canParallel() bool {
	return false
}
