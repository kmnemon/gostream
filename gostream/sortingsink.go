package gostream

import (
	"reflect"
	"sort"
)

type sortingSink[T any] struct {
	less       func(T, T) bool
	list       []T
	downstream sink[T]
}

func (s *sortingSink[T]) begin(size int) {
}

func (s *sortingSink[T]) accept(u T) {
	s.list = append(s.list, u)
}

func (s *sortingSink[T]) end() {
	if s.less == nil {
		sort.Slice(s.list, func(i, j int) bool {
			return isLessBasicTypes(s.list[i], s.list[j])
		})
	} else {
		s.sortedWith()
	}

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

func (s *sortingSink[T]) sortedWith() {
	s.list = quickSort[T](s.list, s.less)
}

// basicTypeSort
func isLessBasicTypes[T any](a T, b T) bool {
	va := reflect.ValueOf(a)
	vb := reflect.ValueOf(b)
	switch va.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return va.Int() < vb.Int()
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return va.Uint() < vb.Uint()
	case reflect.Float32, reflect.Float64:
		return va.Float() < vb.Float()
	case reflect.String:
		return va.String() < vb.String()
	default:
		panic("sort with wrong basic types")
	}
}

// quicksort
func quickSort[T any](arr []T, less func(T, T) bool) []T {
	return _quickSort(arr, 0, len(arr)-1, less)
}

func _quickSort[T any](arr []T, left, right int, less func(T, T) bool) []T {
	if left < right {
		partitionIndex := partition(arr, left, right, less)
		_quickSort(arr, left, partitionIndex-1, less)
		_quickSort(arr, partitionIndex+1, right, less)
	}
	return arr
}

func partition[T any](arr []T, left, right int, less func(T, T) bool) int {
	pivot := left
	index := pivot + 1

	for i := index; i <= right; i++ {
		if less(arr[i], arr[pivot]) {
			swap(arr, i, index)
			index += 1
		}
	}
	swap(arr, pivot, index-1)
	return index - 1
}

func swap[T any](arr []T, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
