package gostream

type Sink[T any] interface {
	begin(uint64)
	accept(T)
	end()
	cancellationRequested() bool
}
