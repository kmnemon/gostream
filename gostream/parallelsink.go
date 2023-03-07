package gostream

import "runtime"

type parallelSink[T any] struct {
	canparallel bool
	list        []T
	downstream  sink[T]
}

func (s *parallelSink[T]) begin(size int) {
	s.downstream.begin(size)
	s.canparallel = s.canParallel()
}

func (s *parallelSink[T]) accept(u T) {
	if !s.canparallel {
		s.downstream.accept(u)
	} else {
		s.list = append(s.list, u)
	}
}

func (s *parallelSink[T]) end() {
	if !s.canparallel {
		s.downstream.end()
	} else {
		cores := runtime.NumCPU()
		for partslice := range splitSlice(s.list, cores) {

		}

	}
}

func (s *parallelSink[T]) isCancellationWasRequested() bool {
	return s.downstream.isCancellationWasRequested()
}

func (s *parallelSink[T]) cancellationRequested() bool {
	return s.downstream.cancellationRequested()
}

func (s *parallelSink[T]) setDownStreamSink(downstream sink[T]) {
	s.downstream = downstream
}

func (s *parallelSink[T]) canParallel() bool {
	return s.downstream.canParallel()
}

func splitSlice[T any](orig []T, count int) [][]T {
	sublen := len(orig) / count
	index := 0
	copySlices := make([][]T, 0, sublen)
	for ; count > 1; count-- {
		cpy := make([]T, 0, len(orig[index:index+sublen]))
		copy(cpy, orig[index:index+sublen])
		copySlices = append(copySlices, cpy)
	}

	//todo copy last one

	return nil
}
