package gostream

import (
	"runtime"
	"sync"
)

type parallelSink[T any] struct {
	canparallel bool
	list        []T
	downstream  sink[T]
}

func (s *parallelSink[T]) begin(size int) {
	s.canparallel = s.canParallel()
	s.downstream.begin(size)
}

func (s *parallelSink[T]) accept(u T) {
	if !s.canparallel {
		s.downstream.accept(u)
	} else {
		s.list = append(s.list, u)
	}
}

func (s *parallelSink[T]) end() {
	if s.canparallel {
		var wg sync.WaitGroup
		cores := runtime.NumCPU()
		for _, slice := range splitSlice(s.list, cores) {
			wg.Add(1)
			go func(slice []T) {
				defer wg.Done()
				for _, v := range slice {
					s.downstream.accept(v)
				}
			}(slice)
		}

		wg.Wait()
	}

	s.downstream.end()

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
	copySlices := make([][]T, 0, count)
	for ; count > 1; count-- {
		cpy := make([]T, len(orig[index:index+sublen]))
		copy(cpy, orig[index:index+sublen])
		copySlices = append(copySlices, cpy)
		index += sublen
	}

	cpy := make([]T, len(orig[index:]))
	copy(cpy, orig[index:])
	copySlices = append(copySlices, cpy)

	return copySlices
}
