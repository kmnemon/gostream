package gostream

type sink[T any] interface {
	begin(int)
	accept(T)
	end()
	cancellationRequested() bool

	setDownStreamSink(sink[T])
}

type mapSink[T any] struct {
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

type forEachSink[T any] struct {
	mapper func(T)
}

func (s *forEachSink[T]) begin(size int) {
}

func (s *forEachSink[T]) accept(u T) {
	s.mapper(u)
}

func (s *forEachSink[T]) end() {
}

func (s *forEachSink[T]) cancellationRequested() bool {
	return false
}

func (s *forEachSink[T]) setDownStreamSink(downstream sink[T]) {
}
