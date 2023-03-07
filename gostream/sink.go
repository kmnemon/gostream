package gostream

type sink[T any] interface {
	begin(int)
	accept(T)
	end()
	isCancellationWasRequested() bool
	cancellationRequested() bool
	canParallel() bool

	setDownStreamSink(sink[T])
}
