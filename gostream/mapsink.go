package gostream

import "golang.org/x/exp/constraints"

type mapSink[T constraints.Ordered] struct {
	downstream sink[T]
	mapper     func(T) T
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

func (s *mapSink[T]) cancellationRequested() bool {
	return false
}

func (s *mapSink[T]) setDownStreamSink(downstream sink[T]) {
	s.downstream = downstream
}
