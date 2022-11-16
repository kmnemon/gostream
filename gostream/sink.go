package gostream

type sink[T any] interface {
	begin(int)
	accept(T)
	end()
	isCancellationWasRequested() bool
	cancellationRequested() bool

	setDownStreamSink(sink[T])
}
