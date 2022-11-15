package gostream

import (
	"sort"

	"golang.org/x/exp/constraints"
)

type sortingSink[T constraints.Ordered] struct {
	downstream sink[T]
	list       []T
}

func (s *sortingSink[T]) begin(size int) {
}

func (s *sortingSink[T]) accept(u T) {
	s.list = append(s.list, u)
}

func (s *sortingSink[T]) end() {
	sort.Slice(s.list, func(i, j int) bool {
		return s.list[i] < s.list[j]
	})

	s.downstream.begin(len(s.list))

	for _, v := range s.list {
		s.downstream.accept(v)
	}

	s.downstream.end()
	s.list = nil
}

func (s *sortingSink[T]) cancellationRequested() bool {
	return false
}

func (s *sortingSink[T]) setDownStreamSink(downstream sink[T]) {
	s.downstream = downstream
}
