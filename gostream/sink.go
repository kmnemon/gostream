package gostream

type sink[T any] interface {
	begin(int)
	accept(T)
	end()
	cancellationRequested() bool

	setDownStreamSink(sink[T])
}






