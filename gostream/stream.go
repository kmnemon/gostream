package gostream

type stream[T any] interface {
	Parallel() stream[T]
	Map(func(T) T) stream[T]
	Reduce(func(T, T) T) T
	ReduceWithInitValue(T, func(T, T) T) T
	ForEach(func(T))
	Sorted() stream[T]
	SortedWith(func(T, T) bool) stream[T]
	Filter(func(T) bool) stream[T]
	Limit(int) stream[T]
	FindFirst() T
	ToList() []T
	Distinct() stream[T]
	DistinctWith(func(T, T) bool) stream[T]
}

type EntrySet[K any, V any] struct {
	K K
	V V
}

func StreamOf[T any](sourceData []T) stream[T] {
	p := pipeline[T]{}
	p.init(nil, head, nil)
	p.sourceData = sourceData
	return &p
}

func StreamOfMap[K comparable, V any](sourceData map[K]V) stream[EntrySet[K, V]] {
	entrysets := make([]EntrySet[K, V], 0)
	for k, v := range sourceData {
		entrysets = append(entrysets, EntrySet[K, V]{k, v})
	}

	return StreamOf(entrysets)
}
