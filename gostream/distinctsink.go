package gostream

type distinctSink[T any] struct {
	equal      func(T, T) bool
	list       []T
	downstream sink[T]
}

func (s *distinctSink[T]) begin(size int) {
}

func (s *distinctSink[T]) accept(u T) {
	if contains(s.list, u) {

	}
	s.list = append(s.list, u)

}

func (s *distinctSink[T]) end() {

	s.downstream.begin(len(s.list))

	if !s.isCancellationWasRequested() {
		for _, v := range s.list {
			s.downstream.accept(v)
		}
	} else {
		for _, v := range s.list {
			if s.cancellationRequested() {
				break
			}
			s.downstream.accept(v)
		}

	}

	s.downstream.end()
	s.list = nil
}

func (s *distinctSink[T]) isCancellationWasRequested() bool {
	return s.downstream.isCancellationWasRequested()
}

func (s *distinctSink[T]) cancellationRequested() bool {
	return s.downstream.cancellationRequested()
}

func (s *distinctSink[T]) setDownStreamSink(downstream sink[T]) {
	s.downstream = downstream
}

func contains[T comparable](a []T, c T) bool {
	for _, v := range a {
		if v == c {
			return true
		}
	}

	return false
}
