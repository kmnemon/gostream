package gostream

import (
	"reflect"
)

type distinctSink[T any] struct {
	equal      func(T, T) bool
	list       []T
	downstream sink[T]
}

func (s *distinctSink[T]) begin(size int) {
}

func (s *distinctSink[T]) accept(u T) {
	if s.equal == nil {
		s.equal = isEqualBasicTypes[T]
	}

	if !s.contains(s.list, u) {
		s.list = append(s.list, u)
	}
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

func (s *distinctSink[T]) contains(a []T, c T) bool {
	for _, v := range a {
		if s.equal(v, c) {
			return true
		}
	}

	return false
}

// basicTypeEqual
func isEqualBasicTypes[T any](a T, b T) bool {
	va := reflect.ValueOf(a)
	vb := reflect.ValueOf(b)
	switch va.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return va.Int() == vb.Int()
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return va.Uint() == vb.Uint()
	case reflect.Float32, reflect.Float64:
		return va.Float() == vb.Float()
	case reflect.String:
		return va.String() == vb.String()
	default:
		panic("equal with wrong basic types")
	}
}
