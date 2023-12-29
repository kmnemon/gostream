package gostream

type reduceSink[T any] struct {
	binaryOperator func(T, T) T
	downstream     sink[T]
	result         T
	isFirstValue   bool
}

func (s *reduceSink[T]) begin(size int) {
}

func (s *reduceSink[T]) accept(u T) {
	if s.isFirstValue {
		s.result = u
		s.isFirstValue = false
	} else {
		s.result = s.binaryOperator(s.result, u)
	}
}

func (s *reduceSink[T]) end() {
}

func (s *reduceSink[T]) isCancellationWasRequested() bool {
	return false
}

func (s *reduceSink[T]) cancellationRequested() bool {
	return false
}

func (s *reduceSink[T]) setDownStreamSink(downstream sink[T]) {
}

func (s *reduceSink[T]) canParallel() bool {
	return false
}
