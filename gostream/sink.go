package gostream

type Sink[T any] interface {
	begin(uint64)
	accept(T)
	end()
	cancellationRequested() bool
}

type MapSink[T any, R any] struct {
	downstream *Sink[R]
	mapper     func(T) R
}

func (s *MapSink[T, R]) accept(u T) {
	r := s.mapper(u)
	(*s.downstream).accept(r)
}

type FilterSink[T any] struct {
	downstream *Sink[T]
}
