package gostream

import (
	"sync"
)

type toListSink[T any] struct {
	list     []T
	parallel bool
	buf      chan T
	wg       sync.WaitGroup
}

func (s *toListSink[T]) begin(size int) {
	if s.parallel {
		s.buf = make(chan T, 100)
		s.wg.Add(1)
		go func() {
			defer s.wg.Done()
			for u := range s.buf {
				s.list = append(s.list, u)
			}
		}()
	}

}

func (s *toListSink[T]) accept(u T) {
	if !s.parallel {
		s.list = append(s.list, u)
	} else {
		s.buf <- u
	}
}

func (s *toListSink[T]) end() {
	if s.parallel {
		close(s.buf)
	}

	s.wg.Wait()
}

func (s *toListSink[T]) isCancellationWasRequested() bool {
	return false
}

func (s *toListSink[T]) cancellationRequested() bool {
	return false
}

func (s *toListSink[T]) setDownStreamSink(downstream sink[T]) {
}

func (s *toListSink[T]) canParallel() bool {
	s.parallel = true
	return s.parallel
}
