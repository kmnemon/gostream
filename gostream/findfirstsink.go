package gostream

type findFirstSink[T any] struct {
	result T
	cancel bool
}

func (s *findFirstSink[T]) begin(size int) {
}

func (s *findFirstSink[T]) accept(u T) {
	s.result = u
	s.cancel = true
}

func (s *findFirstSink[T]) end() {
}

func (s *findFirstSink[T]) isCancellationWasRequested() bool {
	return true
}

func (s *findFirstSink[T]) cancellationRequested() bool {
	return s.cancel
}

func (s *findFirstSink[T]) setDownStreamSink(downstream sink[T]) {
}
