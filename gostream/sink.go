package gostream

type Sink[T any] interface {
	begin(uint64)
	accept(T)
	end()
	cancellationRequested() bool
}

type MapSink[T any, R any] struct {
	downstream Sink[R]
	mapper     func(T) R
}

func (s *MapSink[T, R]) begin(size uint64) {
}

func (s *MapSink[T, R]) accept(u T) {
	r := s.mapper(u)
	s.downstream.accept(r)
}

func (s *MapSink[T, R]) end() {
	s.downstream.end()
}

func (s *MapSink[T, R]) cancellationRequested() bool {
	return false
}

type FilterSink[T any] struct {
	downstream *Sink[T]
}
